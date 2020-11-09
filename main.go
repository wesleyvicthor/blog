package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var t *template.Template

type Page struct {
	Host  string
	Post  Post
	Posts []*Post
}

func init() {
	t, _ = template.ParseFiles("tpl/home.html", "tpl/post.html")
}

func me(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "About myself")
}

// once deployed move template parser outside
func home(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if len(path) == 0 {
		must(t.ExecuteTemplate(w, "home.html", Page{Host: "https://wmsan.dev"}))
		return
	}
	f, err := os.Open(path + ".md")
	if err != nil {
		http.NotFound(w, r)
		return
	}

	page := Page{
		Host: "https://wmsan.dev",
		Post: NewPost(f),
	}

	must(t.ExecuteTemplate(w, "post.html", page))
}

func main() {
	handler := http.NewServeMux()

	static := http.FileServer(http.Dir("./assets"))
	handler.Handle("/assets/", http.StripPrefix("/assets/", static))

	handler.HandleFunc("/", home)
	handler.HandleFunc("/me", me)

	s := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		fmt.Println("listening on " + s.Addr)
		//s.ListenAndServeTLS("/etc/letsencrypt/live/wmsan.dev/fullchain.pem", "/etc/letsencrypt/live/wmsan.dev/privkey.pem")
		must(s.ListenAndServe())
	}()

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	<-term
	must(s.Shutdown(context.Background()))
	fmt.Println("Server Down")
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

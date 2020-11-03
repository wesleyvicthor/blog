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

var posts = map[string][]byte{
	"how-this-works": []byte("something"),
}

func me(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "About myself")
}

// once deployed move the parser outside
func home(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if _, ok := posts[path]; !ok && len(path) > 0 {
		http.NotFound(w, r)
		return
	}

	post := Post{}

	vars := struct {
		Host  string
		Post  Post
		Posts []*Post
	}{
		Host: "https://wmsan.dev",
		Post: post,
	}

	var t *template.Template
	if len(path) > 0 {
		t, _ = template.ParseFiles("tpl/post.html")
		must(t.Execute(w, vars))
		return
	}

	t, _ = template.ParseFiles("tpl/home.html")
	must(t.Execute(w, vars))
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

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
	"time"
)

var t *template.Template

type Page struct {
	Year  int
	Host  string
	Post  Post
	Posts []*Post
}

func NewPage(post Post) *Page {
	return &Page{
		Year: time.Now().Year(),
		Host: "https://wmsan.dev",
		Post: post,
	}
}

func init() {
	t, _ = template.ParseFiles("tpl/home.html", "tpl/post.html")
}

func me(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "About myself")
}

func home(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if len(path) == 0 {
		must(t.ExecuteTemplate(w, "home.html", NewPage(Post{})))
		return
	}

	f, err := os.Open(path + ".html")
	if err != nil {
		http.NotFound(w, r)
		return
	}

	must(t.ExecuteTemplate(w, "post.html", NewPage(NewPost(f))))
}

func main() {
	handler := http.NewServeMux()

	static := http.FileServer(http.Dir("./assets"))
	handler.Handle("/assets/", http.StripPrefix("/assets/", static))

	handler.HandleFunc("/", home)
	handler.HandleFunc("/me", me)

	port := os.Getenv("B_ADDR")
	s := http.Server{
		Addr:    port,
		Handler: handler,
	}

	go func() {
		fmt.Println("listening on " + s.Addr)

		if os.Getenv("B_ENV") == "dev" {
			must(s.ListenAndServe())

			return
		}

		must(s.ListenAndServeTLS(os.Getenv("B_SSL_FULLCHAIN"), os.Getenv("B_SSL_PRIVKEY")))
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

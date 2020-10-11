package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tpl/home.html")

	t.Execute(w, nil)
}

func main() {
	users := Users()
	session := Session{}
	handler := http.NewServeMux()
	handler.HandleFunc("/", home)
	handler.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		var result string
		t, _ := template.ParseFiles("tpl/auth.html")

		if r.Method != http.MethodPost {
			t.Execute(w, result)
			return
		}

		username := r.FormValue("username")
		passw := r.FormValue("passw")

		user, ok := users[username]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			t.Execute(w, "Invalid username or password")
			return
		}

		err := user.CheckPassw(passw)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			t.Execute(w, "Invalid username or password")
			return
		}

		_, noCookie := r.Cookie("_ss")
		if noCookie != nil {
			http.SetCookie(w, &http.Cookie{
				Name:  "_ss",
				Value: session.GenerateUuid(),
			})
		}

		// generate session id
		// store session id local and on browser
		t.Execute(w, "Logged :D")
	})

	s := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		fmt.Println("listening on " + s.Addr)
		s.ListenAndServe()
	}()

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	<-term
	s.Shutdown(context.Background())
	fmt.Println("Server Down")
}

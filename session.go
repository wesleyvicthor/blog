package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

type Session struct {
}

func New() *Session {
	return &Session{}
}

func (s *Session) GenerateUuid() string {
	chars := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	rand.Shuffle(len(chars), func(i, j int) {
		chars[i], chars[j] = chars[j], chars[i]
	})

	return string(chars[3:9])
}

func (s *Session) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var result string
	users := make(map[string]*User, 3)
	session := New()
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
}

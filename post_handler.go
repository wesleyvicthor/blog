package main

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type HttpContext struct {
	w http.ResponseWriter
	r *http.Request
}
type HttpResource interface {
}

type Routes map[string]HttpResource

type BlogPost struct {
	HttpResource
}

func (p *BlogPost) Get(id int, ctx HttpContext) {
	fmt.Println(id, ctx)
}

func (p *BlogPost) Post(id, name, ref, date string) {

}

func (p *BlogPost) Put(id int) {

}

func (p *BlogPost) Delete(id int) {

}

func register() {
	resource := &BlogPost{}
	routes := Routes{
		"/posts/:id/:name/:ref/:date": resource,
		"/posts":                      resource,
	}

	// POST -> /posts
	// GET -> /posts
	// DELETE -> /posts !( http resources are the combination of url and the method defined parameters)
	// PUT -> /posts

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		for u, re := range routes {
			// use /: to ensure match, later use /<alpha> to check path after url variables
			inx := regexp.MustCompile(":").FindIndex([]byte(u))
			pref := u[:inx[1]-1]

			seg := strings.Split(strings.TrimPrefix(r.URL.Path, pref), "/")

			reg := regexp.MustCompile(":([a-zA-Z0-9]+)")
			m := reg.FindAll([]byte(u), len(u))

			if len(seg) != len(m) {
				continue
			}

			fmt.Fprint(w, strings.Title(strings.ToLower(r.Method)))
			method := strings.Title(strings.ToLower(r.Method))
			t := reflect.ValueOf(re).MethodByName(method)

			// receive method parameter names
			//t.FieldByNameFunc()

			for _, k := range m {
				// check if method parameter names match the route defined ones
				v := t.FieldByName(string(k[1:]))
				fmt.Println(v.String())
			}

			args := []reflect.Value{}
			ctx := HttpContext{w, r}
			args = append(
				args,
				reflect.ValueOf(12),
				reflect.ValueOf(ctx))

			t.Call(args)
		}
	})

	http.ListenAndServe(":8080", server)
}

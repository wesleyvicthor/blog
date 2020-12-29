package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
	"time"
)

type Post struct {
	ID       string
	Headline string
	About    string
	content  string
	datetime string
}

func NewPost(f *os.File) Post {
	var p Post
	p.ID = f.Name()[:len(f.Name())-3]
	s := bufio.NewScanner(f)
	s.Scan()
	p.Headline = s.Text()
	s.Scan()
	p.About = s.Text()
	s.Scan()
	p.datetime = s.Text()

	for s.Scan() {
		if s.Err() == io.EOF {
			break
		}
		p.content += "\n" + s.Text()
	}

	return p
}

func (p Post) Date() string {
	t, _ := time.Parse("2006-01-02", p.datetime)

	return fmt.Sprintf("%s %d, %d", t.Month().String()[:3], t.Day(), t.Year())
}

func (p Post) Body() []template.HTML {
	var b []template.HTML
	for _, s := range strings.Split(p.content, "\n\n") {
		b = append(b, template.HTML(s))
	}

	return b
}

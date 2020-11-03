package main

import (
	"fmt"
	"strings"
	"time"
)

type Post struct {
	id       string
	headline string
	about    string
	body     string
	date     time.Time
}

// think about this...
func NewPost() *Post {
	return &Post{}
}

func (p Post) ID() string {
	return p.id
}

func (Post) Headline() string {
	return "How this works ?"
}

func (Post) About() string {
	return "Its a dance of Golang, AWS, Let's Encrypt and S3"
}

func (Post) Date() string {
	t := time.Now()

	return fmt.Sprintf("%d %s, %d", t.Day(), t.Month().String()[:3], t.Year())
}

func (Post) Body() []string {
	c := `When I decided to start sharing something publicly I took months to choose on which would be
the
approach that would fit my goal... should I do GitHub Pages, Medium bring up a Wordpress or
some staticly
generation tool like jenkyll or hugo. As a core software engineer and a passion to craft I
chose the
not so straight forward path and a more laborious one.

This site is a tiny golang application whose you can find here, the content is served from
an aws ec2 micro instance. as well know it has only 1GB of
RAM and something more than this of storage. which I do not use. The posts are fetched from
S3, I also avoid requesting
s3 for every made request... It uses PureCSS and Microdata Schema

for structuring its information
It uses PureCSS and Microdata Schema for structuring its informationIt uses PureCSS and
Microdata Schema for structuring its informationIt uses PureCSS and Microdata Schema for
structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information

It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its informationIt uses PureCSS and
Microdata Schema for structuring its informationIt uses PureCSS and Microdata Schema for
structuring its informationIt uses PureCSS and Microdata Schema for structuring its
information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
It uses PureCSS and Microdata Schema for structuring its information
`
	return strings.Split(c, "\n\n")
}

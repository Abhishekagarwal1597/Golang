package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Story map[string]Chapter

type handler struct {
	stories  Story
	template *template.Template
}

func newHanlder(s Story, t *template.Template) http.Handler {
	return &handler{stories: s, template: t}
}

// handler interface has implemented the ServerHTTP function so its satisfying http.handler interface
// so it will be called by default
func (sh *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var path string
	if r.URL.Path == "/" {
		path = "intro"
	} else {
		path = strings.TrimLeft(r.URL.Path, "/")
	}
	storyArc := sh.stories[path]
	sh.template.Execute(w, storyArc)
}


var tpl *template.Template

func main() {
	fd := flag.String("h", "gopher.json", "sammple gopher file")
	flag.Parse()
	file, err := os.Open(*fd)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	var data []byte
	data, _ = io.ReadAll(file)
	fmt.Println("data:", string(data))

	var story map[string]Chapter
	json.Unmarshal([]byte(string(data)), &story)

	fmt.Println("story:", story)
	t, err := template.ParseFiles("story.html")
	if err != nil {
		panic(err)
	}

	//we provide a handler function here which should return http.handler
	http.ListenAndServe(":3000", newHanlder(story, t))
	// fmt.Println("jsonData:", jsonData)
}

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

const (
	PATH = "web_app/"
)

var templates = template.Must(template.ParseFiles(PATH + "edit.html", PATH + "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")


func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, _ := getTitle(w, r)
		fn(w, r, title)
	}
}

func (p *Page) save() error {
	fileName := PATH + p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := PATH + title + ".txt"
	body, error := os.ReadFile(fileName)

	if error != nil {
		return nil, error
	}

	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Opaque)

	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title.")
	}

	return m[2], nil
}

func handler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Fprintf(w, "Hi there, I love %s", title)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/view/"+title, http.StatusCreated)
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}

	err := p1.save()
	if err != nil {
		fmt.Println(err)
		return
	}
	p2, _ := loadPage("TestPage")

	fmt.Println(string(p2.Body))

	http.HandleFunc("/", makeHandler(handler))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

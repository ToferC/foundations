package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/toferc/foundations/models"
)

// WebUser represents a generic user struct
type WebUser struct {
	IsAuthor    bool
	SessionUser string
	IsLoggedIn  string
	IsAdmin     string
	Users       []*models.User
}

// WebView is a framework to send objects & data to a Web view
type WebView struct {
	User        *models.User
	Episode     *models.Episode
	IsAuthor    bool
	SessionUser string
	IsLoggedIn  string
	IsAdmin     string

	Episodes []*models.Episode

	Counter    []int
	MidCounter []int
	BigCounter []int

	Flashes []interface{}
}

// SplitLines transfomrs results text string into slice
func SplitLines(s string) []string {
	sli := strings.Split(s, "/n")
	return sli
}

func sliceString(s string, i int) string {

	l := len(s)

	if l > i {
		return s[:i] + "..."
	}
	return s[:l]
}

func subtract(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func isIn(s []int, t int) bool {
	for _, n := range s {
		if n == t {
			return true
		}
	}
	return false
}

func isInString(s []string, t string) bool {
	for _, n := range s {
		if n == t {
			return true
		}
	}
	return false
}

// Render combines templates, funcs and renders all Web pages in the app
func Render(w http.ResponseWriter, filename string, data interface{}) {

	tmpl := make(map[string]*template.Template)

	// Set up FuncMap
	funcMap := template.FuncMap{
		"subtract":    subtract,
		"add":         add,
		"multiply":    multiply,
		"isIn":        isIn,
		"sliceString": sliceString,
		"isInString":  isInString,
	}

	baseTemplate := "templates/layout.html"
	userFrame := "templates/userframe.html"

	tmpl[filename] = template.Must(template.New("").Funcs(funcMap).ParseFiles(filename, baseTemplate, userFrame))

	if err := tmpl[filename].ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

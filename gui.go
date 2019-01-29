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
	UserFrame   bool
}

// WebView is a framework to send objects & data to a Web view
type WebView struct {
	User        *models.User
	Episode     *models.Episode
	Experience  *models.Experience
	IsAuthor    bool
	SessionUser string
	IsLoggedIn  string
	IsAdmin     string
	UserFrame   bool

	Episodes          []*models.Episode
	Experiences       []*models.Experience
	LearningResources []*models.LearningResource
	Markdown          template.HTML

	Architecture []models.Stream

	CategoryMap map[string]int

	Users []*models.User

	Counter     []int
	MidCounter  []int
	BigCounter  []int
	StringArray []string

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

func divide(a, b int) int {
	return a / b
}

func percent(a, b int) float32 {
	return (float32(a) / float32(b)) * 100.0
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
		"divide":      divide,
		"percent":     percent,
		"isIn":        isIn,
		"sliceString": sliceString,
		"isInString":  isInString,
	}

	baseTemplate := "templates/layout.html"
	userFrame := "templates/userframe.html"
	footer := "templates/footer.html"

	tmpl[filename] = template.Must(template.New("").Funcs(funcMap).ParseFiles(filename, baseTemplate, userFrame, footer))

	if err := tmpl[filename].ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

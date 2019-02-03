package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/toferc/foundations/models"
)

// WebView is a framework to send objects & data to a Web view
type WebView struct {
	User             *models.User
	Users            []*models.User
	Episode          *models.Episode
	Experience       *models.Experience
	Stream           *models.Stream
	LearningResource *models.LearningResource
	IsAuthor         bool
	SessionUser      string
	IsLoggedIn       string
	IsAdmin          string
	UserFrame        bool

	Episodes          []*models.Episode
	Experiences       []*models.Experience
	LearningResources []*models.LearningResource
	Markdown          template.HTML

	Architecture []models.Stream

	CategoryMap map[string]int

	Counter     []int
	MidCounter  []int
	BigCounter  []int
	StringArray []string
	NumMap      map[int]string
	StringMap   map[string]string

	Flashes []interface{}
}

var fontMap = map[string]string{
	"read":        `<i class="fas fa-book"></i>`,
	"watch":       `<i class="fab fa-youtube"></i>`,
	"listen":      `<i class="fas fa-podcast"></i>`,
	"participate": `<i class="fas fa-users"></i>`,
	"practice":    `<i class="fas fa-pen"></i>`,
	"study":       `<i class="fas fa-university"></i>`,
	"do":          `<i class="fas fa-tools"></i>`,
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

func sliceFormat(sl []string) string {

	text := ""

	for _, s := range sl {
		text += fmt.Sprintf("%s | ", s)
	}

	return strings.TrimSuffix(text, " | ")
}

func noEscape(s string) template.HTML {
	return template.HTML(s)
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
		"sliceFormat": sliceFormat,
		"noEscape":    noEscape,
	}

	baseTemplate := "templates/layout.html"
	userFrame := "templates/userframe.html"
	footer := "templates/footer.html"

	tmpl[filename] = template.Must(template.New("").Funcs(funcMap).ParseFiles(filename, baseTemplate, userFrame, footer))

	if err := tmpl[filename].ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

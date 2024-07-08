package main

import (
	"bootdev-clan/internal/ranks"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed templates/*
var templates embed.FS

//go:embed static/*
var static embed.FS

var indexTpl *template.Template
var ranksTpl *template.Template

var site = Configuration{}

type Configuration struct {
	Title string
	Ranks *ranks.Ranks
}

func (c *Configuration) resolveConfiguration() {
	titleRaw, ok := os.LookupEnv("TITLE")
	if ok {
		c.Title = strings.TrimSpace(titleRaw)
	}
	usernamesRaw, ok := os.LookupEnv("USERNAMES")
	if ok {
		usernames := strings.Split(usernamesRaw, ",")
		for i := range usernames {
			usernames[i] = strings.TrimSpace(usernames[i])
		}
		c.Ranks.Usernames = usernames
	}
}

// Template function
func add(a int, b int) int {
	return a + b
}

// Handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTpl.Execute(w, site)
}

func ranksHandler(w http.ResponseWriter, r *http.Request) {
	site.Ranks.UpdateRanks()
	ranksTpl.Execute(w, site)
}

func main() {
	var err error

	// Setup templates
	indexTpl, err = template.ParseFS(templates, "templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	funcMap := template.FuncMap{"add": add}
	ranksTpl, err = template.New("ranks.gohtml").Funcs(funcMap).ParseFS(templates, "templates/ranks.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	// Setup ranks
	site.Ranks = &ranks.Ranks{}
	site.resolveConfiguration()
	site.Ranks.UpdateRanks()

	// Setup web server and routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ranks", ranksHandler)
	http.HandleFunc("/static/", http.FileServerFS(static).ServeHTTP)
	log.Fatal(http.ListenAndServe(":8000", nil))
	log.Print("server started")
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gosimple/slug"
	blackfriday "gopkg.in/russross/blackfriday.v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"github.com/thewhitetulip/Tasks/sessions"
	"github.com/toferc/foundations/database"
	"github.com/toferc/foundations/models"
)

// SplashPageHandler renders the basic character roster page
func SplashPageHandler(w http.ResponseWriter, req *http.Request) {

	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		Render(w, "templates/login.html", nil)
		return
		// in case of error
	}

	// Prep for user authentication
	sessionMap := getUserSessionValues(session)

	username := sessionMap["username"]
	loggedIn := sessionMap["loggedin"]
	isAdmin := sessionMap["isAdmin"]

	episodes, err := database.ListEpisodes(db)
	if err != nil {
		panic(err)
	}

	for _, ep := range episodes {
		if ep.Image == nil {
			ep.Image = new(models.Image)
			ep.Image.Path = DefaultEpisodeImage
		}
	}

	wv := WebView{
		SessionUser:  username,
		IsLoggedIn:   loggedIn,
		IsAdmin:      isAdmin,
		Episodes:     episodes,
		UserFrame:    true,
		Architecture: baseArchitecture,
	}
	Render(w, "templates/splash_page.html", wv)
}

// EpisodeHandler renders a character in a Web page
func EpisodeHandler(w http.ResponseWriter, req *http.Request) {

	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		// in case of error
	}

	// Prep for user authentication
	sessionMap := getUserSessionValues(session)

	username := sessionMap["username"]
	loggedIn := sessionMap["loggedin"]
	isAdmin := sessionMap["isAdmin"]

	vars := mux.Vars(req)
	slug := vars["slug"]

	ep, err := database.SlugLoadEpisode(db, slug)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to load Episode")
	}

	IsAuthor := false

	if username == ep.Author.UserName {
		IsAuthor = true
	}

	if ep.Image == nil {
		ep.Image = new(models.Image)
		ep.Image.Path = DefaultEpisodeImage
	}

	input := []byte(ep.Body)

	output := template.HTML(blackfriday.Run(input))

	wv := WebView{
		Episode:      ep,
		IsAuthor:     IsAuthor,
		IsLoggedIn:   loggedIn,
		SessionUser:  username,
		IsAdmin:      isAdmin,
		Markdown:     output,
		UserFrame:    true,
		Architecture: baseArchitecture,
	}

	// Render page
	Render(w, "templates/view_episode.html", wv)

}

// AddEpisodeHandler creates a user-generated episode
func AddEpisodeHandler(w http.ResponseWriter, req *http.Request) {

	// Get session values or redirect to Login
	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		http.Redirect(w, req, "/login/", 302)
		return
		// in case of error
	}

	// Prep for user authentication
	sessionMap := getUserSessionValues(session)

	username := sessionMap["username"]
	loggedIn := sessionMap["loggedin"]
	isAdmin := sessionMap["isAdmin"]

	if username == "" {
		// Add user message
		http.Redirect(w, req, "/", 302)
	}

	ep := models.Episode{
		Title: "New",
	}

	wv := WebView{
		Episode:      &ep,
		IsAuthor:     true,
		SessionUser:  username,
		IsLoggedIn:   loggedIn,
		IsAdmin:      isAdmin,
		Counter:      numToArray(7),
		BigCounter:   numToArray(15),
		Architecture: baseArchitecture,
	}

	if req.Method == "GET" {

		// Render page
		Render(w, "templates/add_episode.html", wv)

	}

	if req.Method == "POST" { // POST

		err := req.ParseMultipartForm(MaxMemory)
		if err != nil {
			panic(err)
		}

		// Map default Episode to Character.Episodes
		ep := models.Episode{}

		author, err := database.LoadUser(db, username)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, req, "/", 302)
		}

		ep.Author = author

		// Pull form values into Episode via gorilla/schema
		err = decoder.Decode(&ep, req.PostForm)
		if err != nil {
			panic(err)
		}

		// Upload image to s3
		file, h, err := req.FormFile("ImagePath")
		switch err {
		case nil:
			// Process image
			defer file.Close()
			// example path media/Major/TestImage/Jason_White.jpg
			path := fmt.Sprintf("/media/%s/%s/%s",
				ep.Author.UserName,
				ToSnakeCase(ep.Title),
				h.Filename,
			)

			_, err = uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String(os.Getenv("BUCKET")),
				Key:    aws.String(path),
				Body:   file,
			})
			if err != nil {
				log.Panic(err)
				fmt.Println("Error uploading file ", err)
			}
			fmt.Printf("successfully uploaded %q to %q\n",
				h.Filename, os.Getenv("BUCKET"))

			ep.Image = new(models.Image)
			ep.Image.Path = path

			fmt.Println(path)

		case http.ErrMissingFile:
			log.Println("no file")

		default:
			log.Panic(err)
			fmt.Println("Error getting file ", err)
		}

		// Add other Episode fields

		if len(ep.Videos) > 0 {
			for _, v := range ep.Videos {
				if v.Path != "" {
					v.Path = ConvertURLToEmbededURL(v.Path)
					// Call API to get Title and description
				}
			}
		}

		ep.PublishedOn = time.Now()

		fmt.Println(ep.Tags)

		for _, stream := range baseArchitecture {
			if ep.Experience.Stream.Name == stream.Name {
				ep.Experience.Stream = &models.Stream{
					Name:        stream.Name,
					Practices:   map[string]*models.Practice{},
					Description: stream.Description,
					Image:       stream.Image,
					Slug:        stream.Slug,
				}
				fmt.Println("Added stream " + stream.Name)
			}
		}

		ep.Slug = slug.Make(fmt.Sprintf("%s-%s", ep.Author.UserName, ep.Title))

		err = database.SaveEpisode(db, &ep)
		if err != nil {
			log.Panic(err)
		} else {
			fmt.Println("Saved Episode")
		}

		url := fmt.Sprintf("/add_episode_practices/%s", ep.Slug)

		http.Redirect(w, req, url, http.StatusFound)
	}
}

// AddEpisodePracticesHandler creates a user-generated episode
func AddEpisodePracticesHandler(w http.ResponseWriter, req *http.Request) {

	// Get session values or redirect to Login
	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		http.Redirect(w, req, "/login/", 302)
		return
		// in case of error
	}

	// Prep for user authentication
	sessionMap := getUserSessionValues(session)

	username := sessionMap["username"]
	loggedIn := sessionMap["loggedin"]
	isAdmin := sessionMap["isAdmin"]

	if username == "" {
		// Add user message
		http.Redirect(w, req, "/", 302)
	}

	vars := mux.Vars(req)
	slugString := vars["slug"]

	ep, err := database.SlugLoadEpisode(db, slugString)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, req, "/", 302)
	}

	// Track existing practices
	stringArray := []string{}

	for _, practice := range ep.Experience.Practices {
		stringArray = append(stringArray, practice.Name)
	}

	// Validate that User == Author
	IsAuthor := false

	if username == ep.Author.UserName || isAdmin == "true" {
		IsAuthor = true
	} else {
		http.Redirect(w, req, "/", 302)
	}

	wv := WebView{
		Episode:      ep,
		IsAuthor:     IsAuthor,
		SessionUser:  username,
		IsLoggedIn:   loggedIn,
		IsAdmin:      isAdmin,
		Counter:      numToArray(7),
		BigCounter:   numToArray(15),
		StringArray:  stringArray,
		Architecture: baseArchitecture,
	}

	if req.Method == "GET" {

		// Render page
		Render(w, "templates/add_episode_practices.html", wv)

	}

	if req.Method == "POST" { // POST

		err := req.ParseForm()
		if err != nil {
			panic(err)
		}

		ep.Experience.Practices = []*models.Practice{}

		// Add practices to user.Stream
		for _, stream := range baseArchitecture {
			if stream.Name == ep.Experience.Stream.Name {
				for pk, pv := range stream.Practices {
					if req.FormValue(pk) != "" {
						ep.Experience.Practices = append(ep.Experience.Practices, pv)
					}
				}
			}
		}

		ep.Slug = slug.Make(fmt.Sprintf("%s-%s", ep.Author.UserName, ep.Title))

		fmt.Println(ep)

		// Update Episode
		err = database.UpdateEpisode(db, ep)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Saved")
		}

		url := fmt.Sprintf("/view_episode/%s", ep.Slug)

		http.Redirect(w, req, url, http.StatusFound)
	}
}

// ModifyEpisodeHandler renders an editable Episode in a Web page
func ModifyEpisodeHandler(w http.ResponseWriter, req *http.Request) {

	// Get session values or redirect to Login
	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		http.Redirect(w, req, "/login/", 302)
		return
		// in case of error
	}

	// Prep for user authentication
	sessionMap := getUserSessionValues(session)

	username := sessionMap["username"]
	loggedIn := sessionMap["loggedin"]
	isAdmin := sessionMap["isAdmin"]

	vars := mux.Vars(req)
	slug := vars["slug"]

	ep, err := database.SlugLoadEpisode(db, slug)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, req, "/", 302)
	}

	if len(ep.Tags) < 3 {
		for i := 0; i < 3; i++ {
			ep.Tags = append(ep.Tags, "")
		}
	}

	// Validate that User == Author
	IsAuthor := false

	if username == ep.Author.UserName || isAdmin == "true" {
		IsAuthor = true
	} else {
		http.Redirect(w, req, "/", 302)
	}

	wv := WebView{
		Episode:      ep,
		Stream:       ep.Experience.Stream,
		IsAuthor:     IsAuthor,
		SessionUser:  username,
		IsLoggedIn:   loggedIn,
		IsAdmin:      isAdmin,
		Counter:      numToArray(9),
		BigCounter:   numToArray(15),
		Architecture: baseArchitecture,
	}

	if req.Method == "GET" {

		// Render page
		Render(w, "templates/modify_episode.html", wv)

	}

	if req.Method == "POST" { // POST

		err := req.ParseMultipartForm(MaxMemory)
		if err != nil {
			panic(err)
		}

		// Pull form values into Episode via gorilla/schema
		err = decoder.Decode(ep, req.PostForm)
		if err != nil {
			panic(err)
		}

		// Upload image to s3
		file, h, err := req.FormFile("ImagePath")
		switch err {
		case nil:
			// Prepess image
			defer file.Close()
			// example path media/Major/TestImage/Jason_White.jpg
			path := fmt.Sprintf("/media/%s/%s/%s",
				ep.Author.UserName,
				ToSnakeCase(ep.Title),
				h.Filename,
			)

			_, err = uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String(os.Getenv("BUCKET")),
				Key:    aws.String(path),
				Body:   file,
			})
			if err != nil {
				log.Panic(err)
				fmt.Println("Error uploading file ", err)
			}
			fmt.Printf("successfully uploaded %q to %q\n",
				h.Filename, os.Getenv("BUCKET"))

			ep.Image = new(models.Image)
			ep.Image.Path = path

			fmt.Println(path)

		case http.ErrMissingFile:
			log.Println("no file")

		default:
			log.Panic(err)
			fmt.Println("Error getting file ", err)
		}

		// Do things

		for _, v := range ep.Videos {
			v.Path = ConvertURLToEmbededURL(v.Path)
		}

		for _, stream := range baseArchitecture {
			if ep.Experience.Stream.Name == stream.Name {
				ep.Experience.Stream = &models.Stream{
					Name:        stream.Name,
					Practices:   map[string]*models.Practice{},
					Description: stream.Description,
					Image:       stream.Image,
					Slug:        stream.Slug,
				}
				fmt.Println("Added stream " + stream.Name)
			}
		}

		// Insert Episode into App archive
		err = database.UpdateEpisode(db, ep)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Saved")
		}

		fmt.Println(ep)

		url := fmt.Sprintf("/add_episode_practices/%s", ep.Slug)

		http.Redirect(w, req, url, http.StatusSeeOther)
	}
}

// DeleteEpisodeHandler renders a character in a Web page
func DeleteEpisodeHandler(w http.ResponseWriter, req *http.Request) {

	// Get session values or redirect to Login
	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		http.Redirect(w, req, "/login/", 302)
		return
		// in case of error
	}

	// Prep for user authentication
	sessionMap := getUserSessionValues(session)

	username := sessionMap["username"]
	loggedIn := sessionMap["loggedin"]
	isAdmin := sessionMap["isAdmin"]

	vars := mux.Vars(req)
	pk := vars["id"]

	id, err := strconv.Atoi(pk)
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	ep, err := database.PKLoadEpisode(db, int64(id))
	if err != nil {
		fmt.Println(err)
	}

	if ep.Image == nil {
		ep.Image = new(models.Image)
		ep.Image.Path = DefaultEpisodeImage
	}

	input := []byte(ep.Body)

	output := template.HTML(blackfriday.Run(input))

	// Validate that User == Author
	IsAuthor := false

	if username == ep.Author.UserName || isAdmin == "true" {
		IsAuthor = true
	} else {
		http.Redirect(w, req, "/", 302)
	}

	wv := WebView{
		Episode:      ep,
		IsAuthor:     IsAuthor,
		SessionUser:  username,
		IsLoggedIn:   loggedIn,
		IsAdmin:      isAdmin,
		Markdown:     output,
		Architecture: baseArchitecture,
	}

	if req.Method == "GET" {

		// Render page
		Render(w, "templates/delete_episode.html", wv)

	}

	if req.Method == "POST" {

		err := database.DeleteEpisode(db, ep.ID)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Deleted Episode")
		}

		url := fmt.Sprint("/")

		http.Redirect(w, req, url, http.StatusSeeOther)
	}
}

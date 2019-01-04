package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"github.com/thewhitetulip/Tasks/sessions"
	"github.com/toferc/foundations/database"
	"github.com/toferc/foundations/models"
)

// EpisodeIndexHandler renders the basic character roster page
func EpisodeIndexHandler(w http.ResponseWriter, req *http.Request) {

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

	wc := WebView{
		SessionUser: username,
		IsLoggedIn:  loggedIn,
		IsAdmin:     isAdmin,
		Episodes:    episodes,
	}
	Render(w, "templates/episode_index.html", wc)
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
	pk := vars["id"]

	if len(pk) == 0 {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	id, err := strconv.Atoi(pk)
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	ep, err := database.PKLoadEpisode(db, int64(id))
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to load Episode")
	}

	fmt.Println(ep)

	IsAuthor := false

	if username == ep.Author.UserName {
		IsAuthor = true
	}

	if ep.Image == nil {
		ep.Image = new(models.Image)
		ep.Image.Path = DefaultEpisodeImage
	}

	episodes, err := database.ListEpisodes(db)
	if err != nil {
		panic(err)
	}

	wc := WebView{
		Episode:     ep,
		IsAuthor:    IsAuthor,
		IsLoggedIn:  loggedIn,
		SessionUser: username,
		IsAdmin:     isAdmin,
		Episodes:    episodes,
	}

	// Render page
	Render(w, "templates/view_episode.html", wc)

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

	episodes, err := database.ListEpisodes(db)
	if err != nil {
		panic(err)
	}

	wc := WebView{
		Episode:     &ep,
		IsAuthor:    true,
		SessionUser: username,
		IsLoggedIn:  loggedIn,
		IsAdmin:     isAdmin,
		Counter:     numToArray(7),
		BigCounter:  numToArray(15),
		Episodes:    episodes,
	}

	if req.Method == "GET" {

		// Render page
		Render(w, "templates/add_episode.html", wc)

	}

	if req.Method == "POST" { // POST

		err := req.ParseMultipartForm(MaxMemory)
		if err != nil {
			panic(err)
		}

		user, err := database.LoadUser(db, username)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, req, "/", 302)
		}

		// Map default Episode to Character.Episodes
		ep := models.Episode{
			Title: req.FormValue("Title"),
			Body:  req.FormValue("Body"),
		}

		// Upload image to s3
		file, h, err := req.FormFile("image")
		switch err {
		case nil:
			// Process image
			defer file.Close()
			// example path media/Major/TestImage/Jason_White.jpg
			path := fmt.Sprintf("/media/%s/%s/%s",
				ep.Author.UserName,
				runequest.ToSnakeCase(ep.Episode.Name),
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

		author, err := database.LoadUser(db, username)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, req, "/", 302)
		}

		ep.Author = author

		err = database.SaveEpisode(db, &ep)
		if err != nil {
			log.Panic(err)
		} else {
			fmt.Println("Saved Episode")
		}

		url := fmt.Sprintf("/view_episode/%d", ep.ID)

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
	pk := vars["id"]

	id, err := strconv.Atoi(pk)
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	ep, err := database.PKLoadEpisode(db, int64(id))
	if err != nil {
		fmt.Println(err)
	}

	if ep.Author == nil {
		ep.Author = &models.User{
			UserName: "",
		}
	}

	// Validate that User == Author
	IsAuthor := false

	if username == ep.Author.UserName || isAdmin == "true" {
		IsAuthor = true
	} else {
		http.Redirect(w, req, "/", 302)
	}

	wc := WebView{
		Episode:     ep,
		IsAuthor:    IsAuthor,
		SessionUser: username,
		IsLoggedIn:  loggedIn,
		IsAdmin:     isAdmin,
		Counter:     numToArray(9),
		BigCounter:  numToArray(15),
	}

	if req.Method == "GET" {

		// Render page
		Render(w, "templates/modify_episode.html", wc)

	}

	if req.Method == "POST" { // POST

		err := req.ParseMultipartForm(MaxMemory)
		if err != nil {
			panic(err)
		}

		// Update Episode here
		ep.Name = req.FormValue("Name")
		ep.Description = req.FormValue("Description")

		// Upload image to s3
		file, h, err := req.FormFile("image")
		switch err {
		case nil:
			// Prepess image
			defer file.Close()
			// example path media/Major/TestImage/Jason_White.jpg
			path := fmt.Sprintf("/media/%s/%s/%s",
				ep.Author.UserName,
				runequest.ToSnakeCase(ep.Episode.Name),
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

		// Insert Episode into App archive
		err = database.UpdateEpisode(db, ep)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Saved")
		}

		fmt.Println(ep)

		url := fmt.Sprintf("/view_episode/%d", ep.ID)

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
		ep.Image.Path = defaultEpisodeImage
	}

	// Validate that User == Author
	IsAuthor := false

	if username == ep.Author.UserName || isAdmin == "true" {
		IsAuthor = true
	} else {
		http.Redirect(w, req, "/", 302)
	}

	wc := WebView{
		Episode:     ep,
		IsAuthor:    IsAuthor,
		SessionUser: username,
		IsLoggedIn:  loggedIn,
		IsAdmin:     isAdmin,
	}

	if req.Method == "GET" {

		// Render page
		Render(w, "templates/delete_episode.html", wc)

	}

	if req.Method == "POST" {

		err := database.DeleteEpisode(db, ep.ID)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Deleted Episode")
		}

		url := fmt.Sprint("/episode_index/")

		http.Redirect(w, req, url, http.StatusSeeOther)
	}
}

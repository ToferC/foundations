package main

/*

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"gopkg.in/russross/blackfriday.v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"github.com/thewhitetulip/Tasks/sessions"
	"github.com/toferc/foundations/database"
	"github.com/toferc/foundations/models"
)


// ListUserExperiencesHandler renders the basic character roster page
func ListUserExperiencesHandler(w http.ResponseWriter, req *http.Request) {

	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		Render(w, "templates/login.html", nil)
		return
		// in case of error
	}

	// Prex for user authentication
	sessionMap := getUserSessionValues(session)

	username := sessionMap["username"]
	loggedIn := sessionMap["loggedin"]
	isAdmin := sessionMap["isAdmin"]

	experiences, err := database.ListExperiences(db)
	if err != nil {
		panic(err)
	}

	for _, ex := range experiences {
		if ex.Image == nil {
			ex.Image = new(models.Image)
			ex.Image.Path = DefaultExperienceImage
		}
	}

	wv := WebView{
		SessionUser: username,
		IsLoggedIn:  loggedIn,
		IsAdmin:     isAdmin,
		Experiences: experiences,
	}
	Render(w, "templates/experiences.html", wv)
}

// ViewExperienceHandler renders a character in a Web page
func ViewExperienceHandler(w http.ResponseWriter, req *http.Request) {

	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		// in case of error
	}

	// Prex for user authentication
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

	ex, err := database.PKLoadExperience(db, int64(id))
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to load Experience")
	}

	IsAuthor := false

	if username == ex.Author.UserName {
		IsAuthor = true
	}

	if ex.Image == nil {
		ex.Image = new(models.Image)
		ex.Image.Path = DefaultExperienceImage
	}

	input := []byte(ex.Body)

	output := template.HTML(blackfriday.Run(input))

	wv := WebView{
		Experience:  ex,
		IsAuthor:    IsAuthor,
		IsLoggedIn:  loggedIn,
		SessionUser: username,
		IsAdmin:     isAdmin,
		Markdown:    output,
	}

	// Render page
	Render(w, "templates/view_experience.html", wv)

}

// AddExperienceHandler creates a user-generated experience
func AddExperienceHandler(w http.ResponseWriter, req *http.Request) {

	// Get session values or redirect to Login
	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		http.Redirect(w, req, "/login/", 302)
		return
		// in case of error
	}

	// Prex for user authentication
	sessionMap := getUserSessionValues(session)

	username := sessionMap["username"]
	loggedIn := sessionMap["loggedin"]
	isAdmin := sessionMap["isAdmin"]

	if username == "" {
		// Add user message
		http.Redirect(w, req, "/", 302)
	}

	ex := models.Experience{
		Title: "New",
	}

	wv := WebView{
		Experience:  &ex,
		IsAuthor:    true,
		SessionUser: username,
		IsLoggedIn:  loggedIn,
		IsAdmin:     isAdmin,
		Counter:     numToArray(7),
		BigCounter:  numToArray(15),
	}

	if req.Method == "GET" {

		// Render page
		Render(w, "templates/add_experience.html", wv)

	}

	if req.Method == "POST" { // POST

		err := req.ParseMultipartForm(MaxMemory)
		if err != nil {
			panic(err)
		}

		// Map default Experience to Character.Experiences
		ex := models.Experience{}

		author, err := database.LoadUser(db, username)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, req, "/", 302)
		}

		ex.Author = author

		// Pull form values into Experience via gorilla/schema
		err = decoder.Decode(&ex, req.PostForm)
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
				ex.Author.UserName,
				ToSnakeCase(ex.Title),
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

			ex.Image = new(models.Image)
			ex.Image.Path = path

			fmt.Println(path)

		case http.ErrMissingFile:
			log.Println("no file")

		default:
			log.Panic(err)
			fmt.Println("Error getting file ", err)
		}

		// Add other Experience fields

		for _, v := range ex.Videos {
			v.Path = ConvertURLToEmbededURL(v.Path)
		}

		ex.PublishedOn = time.Now()

		fmt.Println(ex.Tags)

		err = database.SaveExperience(db, &ex)
		if err != nil {
			log.Panic(err)
		} else {
			fmt.Println("Saved Experience")
		}

		url := fmt.Sprintf("/view_experience/%d", ex.ID)

		http.Redirect(w, req, url, http.StatusFound)
	}
}

// ModifyExperienceHandler renders an editable Experience in a Web page
func ModifyExperienceHandler(w http.ResponseWriter, req *http.Request) {

	// Get session values or redirect to Login
	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		http.Redirect(w, req, "/login/", 302)
		return
		// in case of error
	}

	// Prex for user authentication
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

	ex, err := database.PKLoadExperience(db, int64(id))
	if err != nil {
		fmt.Println(err)
	}

	if ex.Author == nil {
		ex.Author = &models.User{
			UserName: "",
		}
	}

	// Validate that User == Author
	IsAuthor := false

	if username == ex.Author.UserName || isAdmin == "true" {
		IsAuthor = true
	} else {
		http.Redirect(w, req, "/", 302)
	}

	wv := WebView{
		Experience:  ex,
		IsAuthor:    IsAuthor,
		SessionUser: username,
		IsLoggedIn:  loggedIn,
		IsAdmin:     isAdmin,
		Counter:     numToArray(9),
		BigCounter:  numToArray(15),
	}

	if req.Method == "GET" {

		// Render page
		Render(w, "templates/modify_experience.html", wv)

	}

	if req.Method == "POST" { // POST

		err := req.ParseMultipartForm(MaxMemory)
		if err != nil {
			panic(err)
		}

		// Pull form values into Experience via gorilla/schema
		err = decoder.Decode(ex, req.PostForm)
		if err != nil {
			panic(err)
		}

		// Upload image to s3
		file, h, err := req.FormFile("ImagePath")
		switch err {
		case nil:
			// Prexess image
			defer file.Close()
			// example path media/Major/TestImage/Jason_White.jpg
			path := fmt.Sprintf("/media/%s/%s/%s",
				ex.Author.UserName,
				ToSnakeCase(ex.Title),
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

			ex.Image = new(models.Image)
			ex.Image.Path = path

			fmt.Println(path)

		case http.ErrMissingFile:
			log.Println("no file")

		default:
			log.Panic(err)
			fmt.Println("Error getting file ", err)
		}

		// Do things

		for _, v := range ex.Videos {
			v.Path = ConvertURLToEmbededURL(v.Path)
		}

		// Insert Experience into App archive
		err = database.UpdateExperience(db, ex)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Saved")
		}

		fmt.Println(ex)

		url := fmt.Sprintf("/view_experience/%d", ex.ID)

		http.Redirect(w, req, url, http.StatusSeeOther)
	}
}

// DeleteExperienceHandler renders a character in a Web page
func DeleteExperienceHandler(w http.ResponseWriter, req *http.Request) {

	// Get session values or redirect to Login
	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		http.Redirect(w, req, "/login/", 302)
		return
		// in case of error
	}

	// Prex for user authentication
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

	ex, err := database.PKLoadExperience(db, int64(id))
	if err != nil {
		fmt.Println(err)
	}

	if ex.Image == nil {
		ex.Image = new(models.Image)
		ex.Image.Path = DefaultExperienceImage
	}

	// Validate that User == Author
	IsAuthor := false

	if username == ex.Author.UserName || isAdmin == "true" {
		IsAuthor = true
	} else {
		http.Redirect(w, req, "/", 302)
	}

	wv := WebView{
		Experience:  ex,
		IsAuthor:    IsAuthor,
		SessionUser: username,
		IsLoggedIn:  loggedIn,
		IsAdmin:     isAdmin,
	}

	if req.Method == "GET" {

		// Render page
		Render(w, "templates/delete_experience.html", wv)

	}

	if req.Method == "POST" {

		err := database.DeleteExperience(db, ex.ID)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Deleted Experience")
		}

		url := fmt.Sprint("/experience_index/")

		http.Redirect(w, req, url, http.StatusSeeOther)
	}
}

*/

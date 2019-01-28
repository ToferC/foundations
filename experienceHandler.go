package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

	wv := WebView{
		SessionUser: username,
		IsLoggedIn:  loggedIn,
		IsAdmin:     isAdmin,
		Experiences: experiences,
		UserFrame:   true,
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

	wv := WebView{
		Experience:  ex,
		IsLoggedIn:  loggedIn,
		SessionUser: username,
		IsAdmin:     isAdmin,
		UserFrame:   true,
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

	vars := mux.Vars(req)
	verb := vars["verb"]

	ex := &models.Experience{}

	wv := WebView{
		Experience:   ex,
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
		Render(w, "templates/add_experience.html", wv)

	}

	if req.Method == "POST" { // POST

		err := req.ParseForm()
		if err != nil {
			panic(err)
		}

		// Map default Experience to Character.Experiences

		user, err := database.LoadUser(db, username)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, req, "/", 302)
		}

		// Error here
		// Pull form values into Experience via gorilla/schema
		err = decoder.Decode(ex, req.PostForm)
		if err != nil {
			panic(err)
		}

		// Add other Experience fields

		scrapeArray, err := getWebPageDetails(ex.LearningResource.Path, "title")
		if err != nil {
			fmt.Println("Not a valid Web page or page lacking title")
		}
		ex.LearningResource.Title = scrapeArray[0]
		ex.LearningResource.AddedOn = time.Now()
		// ex.LearningResource.Author = username

		ex.UserName = username
		ex.Verb = verb
		ex.OccurredAt = time.Now()

		// See if LearningResource exists and create if needed
		lrExists := database.LearningResourceExists(db, ex.LearningResource.Path)
		fmt.Println(lrExists)

		if !lrExists {
			fmt.Println("Learning Resource not found in DB")
			database.SaveLearningResource(db, ex.LearningResource)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Saved Learning Resource")
			}
		} else {
			ex.LearningResource, _ = database.LoadLearningResource(db, ex.LearningResource.Path)
		}

		// Determine Points for experience - default calculation

		ex.Points = (ex.Time + ex.Value + ex.Difficulty) * 100
		ex.LearningResourceID = ex.LearningResource.ID

		fmt.Println(ex)

		// Save Experience in Database
		err = database.SaveExperience(db, ex)
		if err != nil {
			log.Fatal(err)
		}

		streams := user.Interests.Streams

		for _, s := range streams {
			if s.Name == ex.Stream.Name {
				s.LearningTargets[user.LearnerProfile.CurrentYear][1] += ex.Points
			}
		}

		err = database.UpdateUser(db, user)
		if err != nil {
			log.Panic(err)
		} else {
			fmt.Println("Saved Experience to user LearnerProfile")
		}

		url := "/learner_profile/"

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

	IsAuthor := false

	// Validate that User == Author

	if username == ex.UserName || isAdmin == "true" {
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

		// Do things

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

	// Validate that User == Author
	IsAuthor := false

	if username == ex.UserName || isAdmin == "true" {
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

		url := fmt.Sprint("/learner_profile/")

		http.Redirect(w, req, url, http.StatusSeeOther)
	}
}

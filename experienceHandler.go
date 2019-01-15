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
	fmt.Println(verb)

	ex := &models.Experience{}

	wv := WebView{
		Experience:  ex,
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

		ex.OccurredAt = time.Now()
		ex.Noun.AddedOn = time.Now()
		ex.Noun.Author = username
		ex.UserName = username
		ex.Verb = verb

		fmt.Println(ex)

		// Save Experience in Database
		err = database.SaveExperience(db, ex)
		if err != nil {
			log.Panic(err)
		}

		if user.LearnerProfile == nil {
			user.LearnerProfile = &models.LearnerProfile{}
		}

		user.LearnerProfile.Experiences = []*models.Experience{}

		user.LearnerProfile.Experiences = append(user.LearnerProfile.Experiences, ex)

		err = database.UpdateUser(db, user)
		if err != nil {
			log.Panic(err)
		} else {
			fmt.Println("Saved Experience to user LearnerProfile")
		}

		lrExists := database.LearningResourceExists(db, ex.Noun.Path)
		fmt.Println(lrExists)

		if !lrExists {

			database.SaveLearningResource(db, ex.Noun)
			if err != nil {
				log.Panic(err)
			} else {
				fmt.Println("Saved Learning Resource")
			}
		}

		url := fmt.Sprintf("/learner_profile/%d", user.ID)

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
	/*

		if username == ex.user.UserName || isAdmin == "true" {
			IsAuthor = true
		} else {
			http.Redirect(w, req, "/", 302)
		}
	*/

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
	/*
		if username == ex.user.UserName || isAdmin == "true" {
			IsAuthor = true
		} else {
			http.Redirect(w, req, "/", 302)
		}
	*/

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

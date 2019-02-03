package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thewhitetulip/Tasks/sessions"
	"github.com/toferc/foundations/database"
	"github.com/toferc/foundations/models"
)

// UserIndexHandler handles the basic roster rendering for the app
func UserIndexHandler(w http.ResponseWriter, req *http.Request) {

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

	fmt.Println(session)

	if isAdmin != "true" {
		http.Redirect(w, req, "/", 302)
		return
	}

	users, err := database.ListUsers(db)
	if err != nil {
		panic(err)
	}

	wu := WebUser{
		SessionUser: username,
		IsLoggedIn:  loggedIn,
		IsAdmin:     isAdmin,
		Users:       users,
		UserFrame:   true,
	}

	Render(w, "templates/user_index.html", wu)
}

// UserViewHandler handles the basic roster rendering for the app
func UserViewHandler(w http.ResponseWriter, req *http.Request) {

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
	idString := vars["id"]

	pk, err := strconv.Atoi(idString)
	if err != nil {
		pk = 0
		log.Println(err)
	}

	fmt.Println(session)

	if isAdmin != "true" {
		http.Redirect(w, req, "/", 302)
		return
	}

	user, err := database.PKLoadUser(db, int64(pk))
	if err != nil {
		log.Fatal(err)
		fmt.Println("Unable to load User")
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	// Validate that User == Admin
	IsAuthor := false

	if isAdmin != "true" {
		http.Redirect(w, req, "/", 302)
	}

	exps, err := database.ListUserExperiences(db, username)
	if err != nil {
		log.Fatal(err)
	}

	lrs := []*models.LearningResource{}
	lrStrings := []string{}

	for _, ex := range exps {
		if !isInString(lrStrings, ex.LearningResource.Path) {
			lrs = append(lrs, ex.LearningResource)
			lrStrings = append(lrStrings, ex.LearningResource.Path)
		}
	}

	// Determine overall progression

	var targetTotal, currentTotal int

	for _, s := range user.Streams {
		targetTotal += s.LearningTargets[user.LearnerProfile.CurrentYear][0]
		currentTotal += s.LearningTargets[user.LearnerProfile.CurrentYear][1]
	}

	categories := map[string]int{
		"currentTotal": currentTotal,
		"targetTotal":  targetTotal,
	}

	for _, ex := range exps {
		categories[ex.Stream.Name] += ex.Points
		categories["max"] += ex.Points
	}

	wv := WebView{
		User:              user,
		IsAuthor:          IsAuthor,
		IsLoggedIn:        loggedIn,
		SessionUser:       username,
		IsAdmin:           isAdmin,
		LearningResources: lrs,
		Experiences:       exps,
		CategoryMap:       categories,
		UserFrame:         true,
		NumMap:            skillMap,
		StringMap:         fontMap,
		Architecture:      baseArchitecture,
	}

	Render(w, "templates/user_view.html", wv)
}

// MakeAdminHandler handles the basic roster rendering for the app
func MakeAdminHandler(w http.ResponseWriter, req *http.Request) {

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

	isAdmin := sessionMap["isAdmin"]

	vars := mux.Vars(req)
	idString := vars["id"]

	pk, err := strconv.Atoi(idString)
	if err != nil {
		pk = 0
		log.Println(err)
	}

	fmt.Println(session)

	if isAdmin != "true" {
		http.Redirect(w, req, "/", 302)
		return
	}

	user, err := database.PKLoadUser(db, int64(pk))
	if err != nil {
		log.Fatal(err)
		fmt.Println("Unable to load User")
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	user.IsAdmin = true

	err = database.UpdateUser(db, user)
	if err != nil {
		log.Println(err)
	}

	url := "/user_index/"

	http.Redirect(w, req, url, http.StatusFound)
}

// DeleteUserHandler handles the basic roster rendering for the app
func DeleteUserHandler(w http.ResponseWriter, req *http.Request) {

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
	idString := vars["id"]

	pk, err := strconv.Atoi(idString)
	if err != nil {
		pk = 0
		log.Println(err)
	}

	fmt.Println(session)

	if isAdmin != "true" {
		http.Redirect(w, req, "/", 302)
		return
	}

	user, err := database.PKLoadUser(db, int64(pk))
	if err != nil {
		log.Fatal(err)
		fmt.Println("Unable to load User")
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	user.IsAdmin = true

	err = database.UpdateUser(db, user)
	if err != nil {
		log.Println(err)
	}

	wv := WebView{
		User:         user,
		IsLoggedIn:   loggedIn,
		SessionUser:  username,
		IsAdmin:      isAdmin,
		UserFrame:    false,
		Architecture: baseArchitecture,
	}

	if req.Method == "GET" {
		Render(w, "templates/delete_user.html", wv)
	}

	if req.Method == "POST" {

		err := database.DeleteUser(db, user.ID)
		if err != nil {
			log.Println(err)
		}

		url := "/user_index/"

		http.Redirect(w, req, url, http.StatusFound)
	}

}

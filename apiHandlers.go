package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/toferc/foundations/models"

	"github.com/thewhitetulip/Tasks/sessions"
	"github.com/toferc/foundations/database"
)

// GetExperiences handles the basic roster rendering for the app
func GetExperiences(w http.ResponseWriter, req *http.Request) {

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

	fmt.Println(loggedIn, isAdmin)

	fmt.Println(session)

	if username == "" {
		http.Redirect(w, req, "/", 302)
		return
	}

	exps, err := database.ListExperiences(db)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(exps)
}

// GetExperience handles the basic roster rendering for the app
func GetExperience(w http.ResponseWriter, req *http.Request) {

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

	fmt.Println(loggedIn, isAdmin)

	fmt.Println(session)

	if username == "" {
		http.Redirect(w, req, "/", 302)
		return
	}

	vars := mux.Vars(req)
	idString := vars["id"]

	pk, err := strconv.Atoi(idString)
	if err != nil {
		pk = 0
		log.Println(err)
	}

	exp, err := database.PKLoadExperience(db, int64(pk))
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(exp)
}

// CreateExperience handles the basic roster rendering for the app
func CreateExperience(w http.ResponseWriter, req *http.Request) {

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

	fmt.Println(loggedIn, isAdmin)

	fmt.Println(session)

	if username == "" {
		http.Redirect(w, req, "/", 302)
		return
	}

	vars := mux.Vars(req)
	idString := vars["id"]

	pk, err := strconv.Atoi(idString)
	if err != nil {
		pk = 0
		log.Println(err)
	}

	experience := &models.Experience{}

	err = json.NewDecoder(req.Body).Decode(experience)
	if err != nil {
		log.Println(err)
		http.Redirect(w, req, "/", 302)
	}

	experience.ID = int64(pk)

	_, err = database.SaveExperience(db, experience)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(experience)
}

// UpdateExperience handles the basic roster rendering for the app
func UpdateExperience(w http.ResponseWriter, req *http.Request) {

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

	fmt.Println(loggedIn, isAdmin)

	fmt.Println(session)

	if username == "" {
		http.Redirect(w, req, "/", 302)
		return
	}

	vars := mux.Vars(req)
	idString := vars["id"]

	pk, err := strconv.Atoi(idString)
	if err != nil {
		pk = 0
		log.Println(err)
	}

	experience, err := database.PKLoadExperience(db, int64(pk))

	err = json.NewDecoder(req.Body).Decode(experience)
	if err != nil {
		log.Println(err)
		http.Redirect(w, req, "/", 302)
	}

	_, err = database.UpdateExperience(db, experience)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(experience)
}

// DeleteExperience handles the basic roster rendering for the app
func DeleteExperience(w http.ResponseWriter, req *http.Request) {

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

	fmt.Println(loggedIn, isAdmin)

	fmt.Println(session)

	if username == "" {
		http.Redirect(w, req, "/", 302)
		return
	}

	vars := mux.Vars(req)
	idString := vars["id"]

	pk, err := strconv.Atoi(idString)
	if err != nil {
		pk = 0
		log.Println(err)
	}

	err = database.DeleteExperience(db, int64(pk))
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode("Deleted")
}

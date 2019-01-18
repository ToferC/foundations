package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thewhitetulip/Tasks/sessions"
	"github.com/toferc/foundations/database"
)

// ListLearnerProfileHandler renders the basic character roster page
func ListLearnerProfileHandler(w http.ResponseWriter, req *http.Request) {

	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		Render(w, "templates/login.html", nil)
		return
		// in case of error
	}

	// Prlp for user authentication
	sessionMap := getUserSessionValues(session)

	username := sessionMap["username"]
	loggedIn := sessionMap["loggedin"]
	isAdmin := sessionMap["isAdmin"]

	users, err := database.ListUsers(db)
	if err != nil {
		panic(err)
	}

	wv := WebView{
		SessionUser: username,
		IsLoggedIn:  loggedIn,
		IsAdmin:     isAdmin,
		Users:       users,
	}
	Render(w, "templates/learners.html", wv)
}

// ViewLearnerProfileHandler renders a Learner's profile in a Web page
func ViewLearnerProfileHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Rendering Learner Profile")

	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		// in case of error
	}

	// Prlp for user authentication
	sessionMap := getUserSessionValues(session)

	username := sessionMap["username"]
	loggedIn := sessionMap["loggedin"]
	isAdmin := sessionMap["isAdmin"]

	user, err := database.LoadUser(db, username)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to load User")
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	// Validate that User == Author
	IsAuthor := false

	if username == user.UserName || isAdmin == "true" {
		IsAuthor = true
	} else {
		http.Redirect(w, req, "/", 302)
	}

	exps, err := database.ListUserExperiences(db, username)
	if err != nil {
		fmt.Println(err)
	}

	lrs, err := database.ListUserLearningResources(db, username)
	if err != nil {
		fmt.Println(err)
	}

	// Determine learning category values

	categories := map[string]int{"max": 100}

	for _, ex := range exps {
		categories[ex.Stream.Name] += ex.Points
		categories["max"] += ex.Points
	}

	add := float32(categories["max"]) * 1.3

	categories["max"] = int(add)

	wv := WebView{
		User:              user,
		IsAuthor:          IsAuthor,
		IsLoggedIn:        loggedIn,
		SessionUser:       username,
		IsAdmin:           isAdmin,
		LearningResources: lrs,
		Experiences:       exps,
		CategoryMap:       categories,
	}

	// Render page
	Render(w, "templates/learner_profile.html", wv)

}

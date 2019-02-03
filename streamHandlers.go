package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/toferc/foundations/models"

	"github.com/gorilla/mux"
	"github.com/thewhitetulip/Tasks/sessions"
	"github.com/toferc/foundations/database"
)

// ViewStreamHandler renders a character in a Web page
func ViewStreamHandler(w http.ResponseWriter, req *http.Request) {

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
	streamString := vars["stream"]

	stream := &models.Stream{}

	for _, s := range baseArchitecture {
		if s.Slug == streamString {
			stream = &s
			break
		}
	}

	// Load stream experiences
	exs, err := database.LoadStreamExperiences(db, stream.Name)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to load Experiences")
	}

	wv := WebView{
		Stream:       stream,
		Experiences:  exs,
		IsLoggedIn:   loggedIn,
		SessionUser:  username,
		IsAdmin:      isAdmin,
		StringMap:    fontMap,
		UserFrame:    true,
		Architecture: baseArchitecture,
	}

	if req.Method == "GET" {
		// Render page
		Render(w, "templates/view_stream.html", wv)
	}
}

// AddStreamHandler renders a character in a Web page
func AddStreamHandler(w http.ResponseWriter, req *http.Request) {

	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		// in case of error
	}

	// Prex for user authentication
	sessionMap := getUserSessionValues(session)

	username := sessionMap["username"]

	vars := mux.Vars(req)
	streamString := vars["stream"]

	user, err := database.LoadUser(db, username)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Unable to load User")
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	for _, stream := range baseArchitecture {
		if stream.Slug == streamString {
			user.Streams[stream.Name] = &models.Stream{
				Name:        stream.Name,
				Practices:   map[string]*models.Practice{},
				Description: stream.Description,
				Image:       stream.Image,
				Selected:    true,
				LearningTargets: map[string][]int{
					user.LearnerProfile.CurrentYear: []int{1000, 0},
				},
				Expertise: 1,
			}

			fmt.Println("Added stream " + stream.Name)
			break
		}
	}

	fmt.Println(user.Streams)

	err = database.UpdateUser(db, user)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("Saved LearnerProfile to user LearnerProfile")
	}

	url := "/add_rating_target"

	http.Redirect(w, req, url, http.StatusFound)
}

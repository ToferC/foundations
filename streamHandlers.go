package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	streamInt, err := strconv.Atoi(streamString)
	if err != nil {
		streamInt = 0
		log.Println(err)
	}

	stream := baseArchitecture[streamInt]

	fmt.Println(stream.Name)

	exs, err := database.LoadStreamExperiences(db, stream.Name)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to load Experiences")
	}

	wv := WebView{
		Stream:      &stream,
		Experiences: exs,
		IsLoggedIn:  loggedIn,
		SessionUser: username,
		IsAdmin:     isAdmin,
		StringMap:   fontMap,
		UserFrame:   true,
	}

	if req.Method == "GET" {
		// Render page
		Render(w, "templates/view_stream.html", wv)
	}

}

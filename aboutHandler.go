package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thewhitetulip/Tasks/sessions"
	"github.com/toferc/foundations/models"
)

// AboutHandler renders a character in a Web page
func AboutHandler(w http.ResponseWriter, req *http.Request) {

	// Get session values or redirect to Login
	session, err := sessions.Store.Get(req, "session")

	if err != nil {
		log.Println("error identifying session")
		http.Redirect(w, req, "/login/", 302)
		return
		// in case of error
	}

	// Prep for user authentication
	um := &models.User{}
	username := ""

	// Get session User
	u := session.Values["username"]

	// Type assertation
	if user, ok := u.(string); !ok {
		um.UserName = ""
	} else {
		fmt.Println(user)
		username = user
	}

	fmt.Println(um)

	wv := WebView{
		SessionUser:  username,
		UserFrame:    true,
		Architecture: baseArchitecture,
	}

	// Render page

	Render(w, "templates/about.html", wv)
}

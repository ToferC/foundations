package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/thewhitetulip/Tasks/sessions"
	"github.com/toferc/foundations/database"
	"github.com/toferc/foundations/models"
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
		UserFrame:   true,
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
		log.Fatal(err)
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

	for _, s := range user.Interests.Streams {
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
	}

	// Render page
	Render(w, "templates/learner_profile.html", wv)

}

// AddLearnerProfileHandler renders a Learner's profile in a Web page
func AddLearnerProfileHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Add Learner Profile")

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
		log.Fatal(err)
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

	// Determine learning category values

	if user.Interests == nil {
		user.Interests = &models.InterestMap{}
	}

	categories := map[string]int{"max": 100}

	wv := WebView{
		User:         user,
		IsAuthor:     IsAuthor,
		IsLoggedIn:   loggedIn,
		SessionUser:  username,
		IsAdmin:      isAdmin,
		CategoryMap:  categories,
		Architecture: baseArchitecture,
	}

	if req.Method == "GET" {
		// Render page
		Render(w, "templates/add_learner_profile.html", wv)
	}

	if req.Method == "POST" { // POST

		// Do something with form data

		err := req.ParseForm()
		if err != nil {
			panic(err)
		}

		fmt.Println(req.Form)

		currentTime := time.Now().Year()
		currentYear := strconv.Itoa(currentTime)

		user.LearnerProfile.CurrentYear = currentYear

		for _, stream := range baseArchitecture {
			if req.FormValue(stream.Name) != "" {
				user.Interests.Streams = append(user.Interests.Streams, &models.Stream{
					Name: stream.Name,
					LearningTargets: map[string][]int{
						user.LearnerProfile.CurrentYear: []int{1000, 0},
					},
					Expertise: 1,
				})
			}
		}

		fmt.Println(user.Interests.Streams)

		/*
			if user.LearnerProfile == nil {
				user.LearnerProfile = &models.LearnerProfile{}
			}

			if user.LearnerProfile.LearningTargets == nil {
				user.LearnerProfile.LearningTargets = map[string][]int{
					"2019": []int{10000, 0},
				}
				user.LearnerProfile.CurrentYear = "2019"
			}
		*/

		err = database.UpdateUser(db, user)
		if err != nil {
			log.Panic(err)
		} else {
			fmt.Println("Saved LearnerProfile to user LearnerProfile")
		}

		url := "/add_rating_target"

		http.Redirect(w, req, url, http.StatusFound)
	}

}

// AddRatingTargetHandler renders a Learner's profile in a Web page
func AddRatingTargetHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Add Ratings & Targets")

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
		log.Fatal(err)
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

	// Determine learning category values

	if user.Interests == nil {
		user.Interests = &models.InterestMap{}
	}

	wv := WebView{
		User:         user,
		IsAuthor:     IsAuthor,
		IsLoggedIn:   loggedIn,
		SessionUser:  username,
		IsAdmin:      isAdmin,
		Architecture: baseArchitecture,
	}

	if req.Method == "GET" {
		// Render page
		Render(w, "templates/add_rating_target.html", wv)
	}

	if req.Method == "POST" { // POST

		// Do something with form data

		err := req.ParseForm()
		if err != nil {
			panic(err)
		}

		fmt.Println(req.Form)

		for _, stream := range baseArchitecture {
			for _, iStream := range user.Interests.Streams {
				if iStream.Name == stream.Name {
					// Pull data from form
					skill := req.FormValue(fmt.Sprintf("%s-Skill", iStream.Name))
					sN, err := strconv.Atoi(skill)
					if err != nil {
						log.Panic(err)
						sN = 1
					}

					iStream.Expertise = sN

					target := req.FormValue(fmt.Sprintf("%s-Target", iStream.Name))
					tN, err := strconv.Atoi(target)
					if err != nil {
						log.Panic(err)
						tN = 1000
					}

					if iStream.LearningTargets == nil {
						iStream.LearningTargets = map[string][]int{
							user.LearnerProfile.CurrentYear: []int{tN, 0},
						}
					} else {
						iStream.LearningTargets[user.LearnerProfile.CurrentYear][0] = tN
					}

					for _, p := range stream.Practices {
						if req.FormValue(p.Name) != "" {
							iStream.Practices = append(iStream.Practices, p)
						}
					}

				} // End inner loop
			}
		}

		fmt.Println(user.Interests.Streams)

		err = database.UpdateUser(db, user)
		if err != nil {
			log.Panic(err)
		} else {
			fmt.Println("Saved LearnerProfile to user LearnerProfile")
		}

		url := "/learner_profile/"

		http.Redirect(w, req, url, http.StatusFound)
	}

}

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
		SessionUser:  username,
		IsLoggedIn:   loggedIn,
		IsAdmin:      isAdmin,
		Users:        users,
		UserFrame:    true,
		Architecture: baseArchitecture,
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

	cYear := user.LearnerProfile.CurrentYear

	// Loop over experiences and set up data
	for _, ex := range exps {
		// Add Learning Resources to Slice
		if !isInString(lrStrings, ex.LearningResource.Path) {
			lrs = append(lrs, ex.LearningResource)
			lrStrings = append(lrStrings, ex.LearningResource.Path)
		}
	}

	// Determine overall progression

	var targetTotal, currentTotal int

	for _, s := range user.Streams {
		targetTotal += s.LearningTargets[cYear][0]
		currentTotal += s.LearningTargets[cYear][1]
	}

	categories := map[string]int{
		"currentTotal": currentTotal,
		"targetTotal":  targetTotal,
	}

	for _, ex := range exps {
		categories[ex.Stream.Name] += ex.Points
		categories["max"] += ex.Points
	}

	slugMap := map[string]string{}

	for _, stream := range baseArchitecture {
		slugMap[stream.Name] = stream.Slug
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
		StringMap2:        slugMap,
		Architecture:      baseArchitecture,
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

	stringArray := []string{}

	for _, s := range baseArchitecture {
		for k := range user.Streams {
			if k == s.Name {
				stringArray = append(stringArray, k)
			}
		}
	}

	if user.Streams == nil {
		user.Streams = map[string]*models.Stream{}
	}

	wv := WebView{
		User:         user,
		IsAuthor:     IsAuthor,
		IsLoggedIn:   loggedIn,
		SessionUser:  username,
		IsAdmin:      isAdmin,
		StringArray:  stringArray,
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
			_, ok := user.Streams[stream.Name]
			if req.FormValue(stream.Name) != "" {
				if !ok {
					user.Streams[stream.Name] = &models.Stream{
						Name:        stream.Name,
						Practices:   map[string]*models.Practice{},
						Description: stream.Description,
						Image:       stream.Image,
						Selected:    true,
					}
					fmt.Println("Added stream " + stream.Name)
				}
			} else {
				if ok {
					delete(user.Streams, stream.Name)
					fmt.Println("Deleted Stream " + stream.Name)
				}
			}
		}

		// Add Learning Targets and base interests if needed
		for _, s := range user.Streams {
			if s.LearningTargets == nil {
				s.LearningTargets = map[string][]int{
					user.LearnerProfile.CurrentYear: []int{1000, 0},
				}
				s.Expertise = 1
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

	stringArray := []string{}

	for _, v := range user.Streams {
		for _, practice := range v.Practices {
			stringArray = append(stringArray, practice.Name)
		}
	}

	fmt.Println(stringArray)

	wv := WebView{
		User:         user,
		IsAuthor:     IsAuthor,
		IsLoggedIn:   loggedIn,
		SessionUser:  username,
		IsAdmin:      isAdmin,
		StringArray:  stringArray,
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

		fmt.Println("Form: ")
		fmt.Println(req.Form)

		for k, v := range user.Streams {

			// Reset practices
			v.Practices = map[string]*models.Practice{}

			// Pull data from form

			// Get Skill Rating
			skill := req.FormValue(fmt.Sprintf("%s-Skill", k))
			sN, err := strconv.Atoi(skill)
			if err != nil {
				log.Panic(err)
				sN = 1
			}

			// Set Skill rating
			v.Expertise = sN

			// Get Learning Target
			target := req.FormValue(fmt.Sprintf("%s-Target", k))
			tN, err := strconv.Atoi(target)
			if err != nil {
				log.Panic(err)
				tN = 1000
			}

			// Set learning target for currentYear
			v.LearningTargets[user.LearnerProfile.CurrentYear][0] = tN

			// Add practices to user.Stream
			for _, stream := range baseArchitecture {
				if stream.Name == k {
					for pk, pv := range stream.Practices {
						if req.FormValue(pk) != "" {
							fmt.Println("Found " + pk)
							_, ok := v.Practices[pk]
							if !ok {
								user.Streams[k].Practices[pk] = pv
							}
						}
					}
				}
			}

		} // End upper loop

		fmt.Println(user.Streams)

		err = database.UpdateUser(db, user)
		if err != nil {
			log.Panic(err)
		} else {
			fmt.Println("Saved LearnerProfile to user LearnerProfile")
		}

		var url string

		if user.Onboarded {
			url = "/learner_profile/"
		} else {
			url = "/add_first_experience"
		}

		http.Redirect(w, req, url, http.StatusFound)
	}
}

// AddFirstExperienceHandler renders a Learner's profile in a Web page
func AddFirstExperienceHandler(w http.ResponseWriter, req *http.Request) {

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

	wv := WebView{
		User:         user,
		IsAuthor:     IsAuthor,
		IsLoggedIn:   loggedIn,
		SessionUser:  username,
		IsAdmin:      isAdmin,
		UserFrame:    false,
		Architecture: baseArchitecture,
	}

	// Render page
	Render(w, "templates/add_first_experience.html", wv)
}

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gorilla/sessions"
	"google.golang.org/api/youtube/v3"
)

func getUserSessionValues(s *sessions.Session) map[string]string {

	sessionMap := map[string]string{
		"username": "",
		"loggedin": "false",
		"isAdmin":  "false",
	}

	// Prep for user authentication

	u := s.Values["username"]
	l := s.Values["loggedin"]
	a := s.Values["isAdmin"]

	// Type assertation
	if user, ok := u.(string); !ok {
	} else {
		fmt.Println(user)
		sessionMap["username"] = user
	}

	// Type assertation
	if loggin, ok := l.(string); !ok {
	} else {
		fmt.Println(loggin)
		sessionMap["loggedin"] = loggin
	}

	// Type assertation
	if admin, ok := a.(string); !ok {
	} else {
		fmt.Println(admin)
		sessionMap["isAdmin"] = admin
	}
	return sessionMap
}

// numToArray takes and int and returns an array of [1:int]
func numToArray(m int) []int {

	a := []int{}

	for i := 1; i < m+1; i++ {
		a = append(a, i)
	}

	return a
}

// TrimSliceBrackets trims the brackets from a slice and return ints as a string
func TrimSliceBrackets(s []int) string {
	rs := fmt.Sprintf("%d", s)
	rs = strings.Trim(rs, "[]")
	return rs
}

// UserQuery creates and question and returns the User's input as a string
func UserQuery(q string) string {
	question := bufio.NewReader(os.Stdin)
	fmt.Print(q)
	r, _ := question.ReadString('\n')

	input := strings.Trim(r, " \n")

	return input
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase transforms a string to snake case
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// ConvertURLToEmbededURL receives this https://youtu.be/G3PvTWRIhZA and convert to embedded Youtube URL
func ConvertURLToEmbededURL(s string) string {

	var c []string

	if strings.Contains(s, "embed") {
		return s
	}

	c = strings.Split(s, "youtu.be/")
	url := fmt.Sprintf("https://youtube.com/embed/%s", c[1])
	return url
}

func getWebPageDetails(url string, targets ...string) ([]string, error) {

	webPageInfo := []string{}

	response, err := http.Get(url)
	if err != nil {
		log.Panic(err)
		return []string{}, err
	}
	defer response.Body.Close()

	// Get response body as string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	for _, t := range targets {
		targetIndexes, err := findSubString(t, pageContent)
		if err != nil {
			fmt.Println(err)
			return []string{}, err
		}

		targetStartIndex := targetIndexes[0]
		targetEndIndex := targetIndexes[1]

		resultString := []byte(pageContent[targetStartIndex:targetEndIndex])

		webPageInfo = append(webPageInfo, string(resultString))
	}

	return webPageInfo, nil
}

func findSubString(s, pageContent string) ([]int, error) {
	// Find substrings

	startString := fmt.Sprintf("<%s>", s)
	endString := fmt.Sprintf("</%s>", s)

	startIndex := strings.Index(pageContent, startString)
	if startIndex == -1 {
		return []int{}, fmt.Errorf("No %s index found", s)
	}

	// Advance to end of index to grab target
	startIndex += len(startString)

	endIndex := strings.Index(pageContent, endString)
	if endIndex == -1 {
		return []int{}, fmt.Errorf("No %s index found", s)
	}
	return []int{startIndex, endIndex}, nil
}

func mapVideosListResults(response *youtube.VideoListResponse) map[string]string {

	m := map[string]string{}

	for _, item := range response.Items {
		fmt.Println(item.Id, ": ", item.Snippet.Title)

		switch item.Id {
		case "Title":
			m["Title"] = item.Id
		case "Description":
			m["Description"] = item.Id
		}
	}
	return m
}

func videosListByID(service *youtube.Service, part string, id string) {
	call := service.Videos.List(part)
	if id != "" {
		call = call.Id(id)
	}
	response, err := call.Do()
	if err != nil {
		fmt.Println(err)
	}
	mapVideosListResults(response)
}

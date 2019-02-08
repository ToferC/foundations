package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-pg/pg"
	"github.com/toferc/foundations/database"
	"github.com/toferc/foundations/models"
)

var db *pg.DB

func init() {
	os.Setenv("DBUser", "chris")
	os.Setenv("DBPass", "12345")
	os.Setenv("DBName", "foundations")
}

func main() {

	if os.Getenv("ENVIRONMENT") == "production" {
		url, ok := os.LookupEnv("DATABASE_URL")

		if !ok {
			log.Fatalln("$DATABASE_URL is required")
		}

		options, err := pg.ParseURL(url)

		if err != nil {
			log.Fatalf("Connection error: %s", err.Error())
		}

		db = pg.Connect(options)
	} else {
		db = pg.Connect(&pg.Options{
			User:     os.Getenv("DBUser"),
			Password: os.Getenv("DBPass"),
			Database: os.Getenv("DBName"),
		})
	}

	defer db.Close()

	err := database.InitDB(db)
	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	for {
		var username, email, password, password2 string

		fmt.Println("Create SuperUser for Foundations")

		username = UserQuery("Enter user name (or hit Enter to quit): ")

		if username == "" {
			break
		}

		email = UserQuery("Enter user email: ")
		password = UserQuery("Enter password: ")
		password2 = UserQuery("Re-enter password: ")

		hashedPassword, err := database.HashPassword(password)
		if err != nil {
			fmt.Println(err)
		}

		if password != password2 {
			fmt.Println("Passwords do not match")
			break
		}
		user := models.User{
			UserName: username,
			Email:    email,
			Password: hashedPassword,
			IsAdmin:  true,
			Role:     "admin",
		}

		database.SaveUser(db, &user)
		fmt.Println(user)
		fmt.Printf("Superuser %s created", user.UserName)
		break
	}

}

// UserQuery creates and question and returns the User's input as a string
func UserQuery(q string) string {
	question := bufio.NewReader(os.Stdin)
	fmt.Print(q)
	r, _ := question.ReadString('\n')

	input := strings.Trim(r, " \n")

	return input
}

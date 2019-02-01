package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"github.com/dghubble/gologin"
	"github.com/dghubble/gologin/google"
	"golang.org/x/oauth2"
	googleOAuth2 "golang.org/x/oauth2/google"

	"github.com/joho/godotenv"

	"github.com/go-pg/pg"
	"github.com/toferc/foundations/database"
)

var (
	db       *pg.DB
	svc      *s3.S3
	uploader *s3manager.Uploader
	decoder  = schema.NewDecoder()
)

// MaxMemory is the max upload size for images
const MaxMemory = 2 * 1024 * 1024

// DefaultEpisodeImage is a base image used as a default
const DefaultEpisodeImage = "/media/digital.jpg"

func init() {
	fmt.Println("Initializing")
	os.Setenv("DBUser", "chris")
	os.Setenv("DBPass", "12345")
	os.Setenv("DBName", "foundations")
}

func main() {

	var callback string

	if os.Getenv("ENVIRONMENT") == "production" {
		// Production system
		url, ok := os.LookupEnv("DATABASE_URL")

		if !ok {
			log.Fatalln("$DATABASE_URL is required")
		}

		options, err := pg.ParseURL(url)

		if err != nil {
			log.Fatalf("Connection error: %s", err.Error())
		}

		db = pg.Connect(options)

		// Set Google Oauth Callback
		callback = "https://foundationsapp.herokuapp.com/google/callback"

	} else {
		// Not production
		db = pg.Connect(&pg.Options{
			User:     os.Getenv("DBUser"),
			Password: os.Getenv("DBPass"),
			Database: os.Getenv("DBName"),
		})
		os.Setenv("CookieSecret", "kimchee-typhoon")
		os.Setenv("BUCKET", "foundationsapp")
		os.Setenv("AWS_REGION", "us-east-1")

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		configGoogleOAUTH()

		callback = "http://localhost:8080/google/callback"
	}

	defer db.Close()

	err := database.InitDB(db)
	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	// Create AWS session using local default config
	// or Env Variables if on Heroku
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error creating session", err)
		os.Exit(1)
	}

	fmt.Println("Session created ", sess)
	svc := s3.New(sess)

	fmt.Println(svc)

	uploader = s3manager.NewUploader(sess)

	// Config Google Oauth
	config := &Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	}

	oauth2Config := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Endpoint:     googleOAuth2.Endpoint,
		RedirectURL:  callback,
		Scopes:       []string{"profile", "email"},
	}
	stateConfig := gologin.DebugOnlyCookieConfig

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// Set Schema ignoreunknownkeys to true
	decoder.IgnoreUnknownKeys(true)

	r := mux.NewRouter()

	fmt.Println("Starting Webserver at port " + port)
	r.HandleFunc("/", SplashPageHandler)
	r.HandleFunc("/about/", AboutHandler)

	r.HandleFunc("/signup/", SignUpFunc)
	r.HandleFunc("/login/", LoginFunc)
	r.HandleFunc("/logout/", LogoutFunc)

	r.HandleFunc("/learner_profile/", ViewLearnerProfileHandler)

	r.Handle("/google/login", google.StateHandler(stateConfig, google.LoginHandler(oauth2Config, nil)))
	r.Handle("/google/callback", google.StateHandler(stateConfig, google.CallbackHandler(oauth2Config, googleLoginFunc(), nil)))

	r.HandleFunc("/users/", UserIndexHandler)
	r.HandleFunc("/add_learner_profile", AddLearnerProfileHandler)
	r.HandleFunc("/add_rating_target", AddRatingTargetHandler)

	r.HandleFunc("/view_episode/{id}", EpisodeHandler)
	r.HandleFunc("/new/", AddEpisodeHandler)
	r.HandleFunc("/modify/{id}", ModifyEpisodeHandler)
	r.HandleFunc("/delete/{id}", DeleteEpisodeHandler)

	r.HandleFunc("/add_first_experience", AddFirstExperienceHandler)
	r.HandleFunc("/add_experience/{verb}/", AddExperienceHandler)
	r.HandleFunc("/add_experience_practices/{id}", AddExperiencePracticesHandler)
	r.HandleFunc("/delete_experience/{id}", DeleteExperienceHandler)

	r.HandleFunc("/view_stream/{stream}", ViewStreamHandler)

	r.HandleFunc("/user_index/", UserIndexHandler)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

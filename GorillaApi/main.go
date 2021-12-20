package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB stores the database session information. It is instanciated once.
type DB struct {
	session    *mgo.Session
	collection *mgo.Collection
}

type Language struct {
	// the bson tag defines the communication behaviour with mongo
	// the json tag, the behaviour with the client
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"` // auto generated, hashed id (if omitted)
	Name       string        `json:"name" bson:"name"`
	Year       string        `json:"year" bson:"year"`
	Creators   []string      `json:"creators" bson:"creators"`
	Frameworks []string      `json:"frameworks" bson:"frameworks"`
}

func main() {
	session, err := mgo.Dial("127.0.0.1")           // creates session
	c := session.DB("languages_api").C("languages") // assign database name and collection

	db := &DB{session: session, collection: c}
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Create a new router
	r := mux.NewRouter()

	// Get method will display the existing languages and their respective ids
	r.HandleFunc("/api/languages", db.ShowInfo).Methods("GET")

	// Get method will display a single item from collection
	r.HandleFunc("/api/languages/{id:[a-zA-Z0-9]*}", db.GetLanguage).Methods("GET")

	// Post method will create a new item inside collection
	r.HandleFunc("/api/languages", db.CreateLanguage).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// As a good practice I will enforce timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

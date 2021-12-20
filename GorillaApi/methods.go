package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// This method will expose all languages and their ids from the database
func (db *DB) ShowInfo(w http.ResponseWriter, r *http.Request) {
	var all []Language
	err := db.collection.Find(nil).All(&all) // cast all findings on the 'all' slice
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		// formating response =>     Language : ID
		for _, l := range all {
			w.Write([]byte(l.Name))
			w.Write([]byte(" : "))
			w.Write([]byte(l.ID.Hex()))
			w.Write([]byte("\n"))
		}
	}
}

// Gets complete language information from database
func (db *DB) GetLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // catch the id passed as an extra argument
	w.WriteHeader(http.StatusOK)
	var lang Language

	// find the item that has the corresponding id
	err := db.collection.Find(bson.M{"_id": bson.ObjectIdHex(vars["id"])}).One(&lang)
	if err != nil {
		w.Write([]byte(err.Error())) // id non existent
	} else { // write the json format of the item to the user
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(lang)
		w.Write(response)
		w.Write([]byte("\n"))
	}
}

// Adds new language record to the database
func (db *DB) CreateLanguage(w http.ResponseWriter, r *http.Request) {
	var lang Language

	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &lang)
	// Create a Hash ID to insert
	lang.ID = bson.NewObjectId()

	err := db.collection.Insert(lang) // method to insert new item
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(lang)
		w.Write(response)
	}
}

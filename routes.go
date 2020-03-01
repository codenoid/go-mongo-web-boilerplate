package main

import (
	"gmwb/controllers"
	"net/http"

	"github.com/globalsign/mgo"
)

// Init route initializing
func Init() {
	url := "mongodb://127.0.0.1:27017"

	// connect to mongodb
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	database := session.DB("gmwb_test")

	conn := controllers.Connector{
		Mongo: database,
	}

	// store routes information
	Routes(conn)
}

// Routes define your routes here
func Routes(conn controllers.Connector) {
	// normal page
	http.HandleFunc("/", conn.StaticIndex)

	// setup route
	http.HandleFunc("/setup", conn.Setup)

	// auth page
	http.HandleFunc("/login", conn.Login)
	http.HandleFunc("/logout", conn.Logout)
}

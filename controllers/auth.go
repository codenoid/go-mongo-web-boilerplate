package controllers

import (
	"fmt"
	"gmwb/helpers"
	"gmwb/structs"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

// Login show hello world
func (conn Connector) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		form := `<form action="/login" method="post">
					<input name="username" type="text">
					<input name="password" type="password">
					
					<button type="submit">login</button>
				</form>`

		w.Header().Add("Content-Type", "text/html")
		fmt.Fprint(w, form)
		return
	}

	if r.Method == "POST" {
		r.ParseForm()

		username := r.FormValue("username")
		password := r.FormValue("password")

		var user structs.User

		// find by username
		err := conn.Mongo.C("users").Find(bson.M{"username": username}).One(&user)

		if err != nil {
			http.Redirect(w, r, "/login?msg=error", 302)
			return
		}

		// make sure username is equal with what you got from database
		if user.Username == username {
			// password check
			hashed := helpers.GetMD5Hash(password)
			if user.Password == hashed {
				// todo: use cookie & session
				http.Redirect(w, r, "/?msg=success login", 302)
				return
			}
		}

		http.Redirect(w, r, "/login?msg=failed", 302)
		return
	}
}

// Logout show hello world
func (conn Connector) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

// Setup route
func (conn Connector) Setup(w http.ResponseWriter, r *http.Request) {
	hashed := helpers.GetMD5Hash("admin")
	data := map[string]string{
		"_id":      "static_id",
		"username": "admin",
		"password": hashed,
	}
	conn.Mongo.C("users").Insert(data)

	http.Redirect(w, r, "/login", 302)
}

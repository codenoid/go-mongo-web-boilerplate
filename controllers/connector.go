package controllers

import "github.com/globalsign/mgo"

// Connector used to connecting people; no; thats nokia
// used to share database* connection across package
type Connector struct {
	Mongo *mgo.Database
}

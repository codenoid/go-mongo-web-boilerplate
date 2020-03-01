package structs

// User used to store user-related information
type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

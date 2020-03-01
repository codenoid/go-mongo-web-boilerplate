package main

import "net/http"

// this is main.go file; but you need to make sure
// if this is a main.go file with package main
// and func main() {}, if you already sure this is main.go
// file, then you currently reading main.go file

func main() {
	// init database & routes
	Init()

	http.ListenAndServe(":6969", nil)
}

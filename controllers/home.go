package controllers

import (
	"fmt"
	"net/http"
)

// StaticIndex show hello world
func (c Connector) StaticIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

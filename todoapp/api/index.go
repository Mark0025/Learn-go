package handler

import (
	"fmt"
	"net/http"
)

// Handler - our serverless function handler
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! Harry was here! %s", time.Now())
}

func main() {
	fmt.Println("Run server!")
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

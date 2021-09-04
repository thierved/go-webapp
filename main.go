package main

import (
	"fmt"
	"net/http"
)


func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my page!<h1>")
}

func main() {
	http.HandleFunc("/", handleHome)
	http.ListenAndServe(":3000", nil)
}
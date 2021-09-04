package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my new page!</h1>")

}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<p>More about us?</p>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome)
	r.HandleFunc("/about", handleAbout)
	http.ListenAndServe(":3000", r)
}

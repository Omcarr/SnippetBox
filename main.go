package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	para, err2 := strconv.Atoi(r.PathValue("para"))

	if err != nil || id < 0 || err2 != nil || para < 0 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet for view id %d & para %d", id, para)
	w.Write([]byte(msg))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	// Register the two new handler functions and corresponding route patterns with
	// the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}/{para}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

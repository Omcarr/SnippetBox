package main

import (
	"errors"
	"fmt"

	// "html/template"
	"net/http"
	"strconv"

	"github.com/Omcarr/SnippetBox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	//we can modify the header as much as we please before writing
	w.Header().Add("Server", "Go")
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// w.Write([]byte("Hello from Snippetbox"))
	// files := []string{
	// 	"/home/omkar/Documents/LearningGO/ui/html/pages/home.tmpl",
	// 	"/home/omkar/Documents/LearningGO/ui/html/base.tmpl",
	// 	"/home/omkar/Documents/LearningGO/ui/html/partials/nav.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// Write the snippet data as a plain-text HTTP response body.
	fmt.Fprintf(w, "%+v", snippet)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	//dummy data
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte("Save a new snippet..."))
}

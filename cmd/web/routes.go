package main

import (
	"net/http"
	"os"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir(os.Getenv("STATIC")))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileserver))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	//chain of middlerwears: recoverpanic->logging ip etc->common header manipulation->app routes accesed by the user
	//1. recover panic-> to handle no response case in case a panics arrives in a goroutine
	//after panic the routine will be abandend but whiler decoupling the stack the defered fucntion will be executed
	//because of this structuring of go routines we can recover fron panics and send 500 error to the user even though the app had paniced and couldnt respond with anything
	//2. we can log each requests ip, method and url using the middlewear
	//3. header manipulation. add coomon safety realted headers to each request before execution. security >>>

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux) //alice package lets us make this chain of middlerwears, after all middlerwear layering the servermux is started
	//this approach lets us improve security as well handle repeated work needs for each requets
}

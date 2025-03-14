package main

import (
	// "flag"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// addr := flag.String("addr", ":4000", "HTTP network address")
	// flag.Parse()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000" // Default to 4000 if not set
	}

	//logger creation
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	//injecting logger dependency to the entire app
	app := &application{
		logger: logger,
	}

	// logger.Info("port", port)
	logger.Info("starting server on", slog.String("port", port))
	err = http.ListenAndServe(port, app.routes())

	logger.Error(err.Error())
	os.Exit(1)

}

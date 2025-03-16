package main

import (
	// "flag"
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Omcarr/SnippetBox/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
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


	//get a db connection pool amnd defer close it
	dbCredentials := os.Getenv("dbCredentials")
	// dbCredentials := flag.String("dsn", "omkar:pass@password/snippetbox?parseTime=true", "snippetbox")

	db, err := openDB(dbCredentials)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("database pool established succesfully")
	defer db.Close()


	//injecting logger dependency to the entire app
	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{DB: db},
	}

	// logger.Info("port", port)
	logger.Info("starting server on", slog.String("port", port))
	err = http.ListenAndServe(port, app.routes())

	logger.Error(err.Error())
	os.Exit(1)

}

func openDB(dbCredentials string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbCredentials)

	//failed to connect to db
	if err != nil {
		return nil, err
	}

	//db didnt respond
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	//db connection pool established succesfully
	return db, nil

}

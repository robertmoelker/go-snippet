package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/robertmoelker/lets-go/internal/models"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	tasks    *models.TaskModel
}

func main() {
	port := flag.String("port", ":4000", "HTTP server port")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		errorLog.Fatal("Error loading .env file")
	}

	// Initialize the database
	db, err := initDb()
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// Setup dependencies for the application
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		tasks:    &models.TaskModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *port)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func initDb() (*sql.DB, error) {
	databaseUrl := os.Getenv("TURSO_DATABASE_URL")
	databaseToken := os.Getenv("TURSO_AUTH_TOKEN")

	url := databaseUrl + "?authToken=" + databaseToken

	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

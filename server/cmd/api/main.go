package main

import (
	"context"
	"database/sql"
	"distnet/internal/models"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config   config
	errorLog *log.Logger
	infoLog  *log.Logger
	models   models.Models
}

func main() {
	var cfg config

	flag.StringVar(&cfg.db.dsn, "db-dsn", "user:password@/distnetdb?parseTime=true", "MySQL DSN")
	// flag.StringVar(&cfg.db.dsn, "db-dsn", "user:password@tcp(db:3306)/distnetdb?parseTime=true", "MySQL DSN")
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	infoLog.Printf("Database connection established")
	app := &application{
		config:   cfg,
		errorLog: errorLog,
		infoLog:  infoLog,
		models:   models.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	infoLog.Printf("Starting server on %d", cfg.port)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

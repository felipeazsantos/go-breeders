package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/felipeazsantos/breeders/configuration"
	_ "github.com/go-sql-driver/mysql"
)

const port = ":4001"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application
	catService  *RemoteService
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:password@tcp(127.0.0.1:3310)/breeders?parseTime=true&tls=false&collation=utf8mb4_unicode_ci&timeout=5s", "MariaDB connection string")
	flag.Parse()

	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}

	// jsonBackend := &JSONBackend{}
	// jsonAdapter := &RemoteService{Remote: jsonBackend}

	xmlBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{Remote: xmlBackend}

	app.catService = xmlAdapter
	app.App = configuration.New(db)

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting web application on port", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"os"
	"testing"

	"github.com/felipeazsantos/breeders/models"
)

var testApp application

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags

	db, err := initMySQLDB("mariadb:password@tcp(127.0.0.1:3310)/breeders?parseTime=true&tls=false&collation=utf8mb4_unicode_ci&timeout=5s")
	if err != nil {
		log.Panic(err)
	}
	testApp = application{
		DB:     db,
		Models: models.New(db),
	}

	os.Exit(m.Run())
}

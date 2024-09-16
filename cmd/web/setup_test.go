package main

import (
	"os"
	"testing"

	"github.com/felipeazsantos/breeders/models"
)

var testApp application

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	testApp = application{
		Models: models.New(nil),
	}

	os.Exit(m.Run())
}

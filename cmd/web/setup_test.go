package main

import (
	"os"
	"testing"

	"github.com/felipeazsantos/breeders/configuration"
)

var testApp application

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	testApp = application{
		App: configuration.New(nil),
	}

	os.Exit(m.Run())
}

package main

import (
	"os"
	"testing"

	"github.com/felipeazsantos/breeders/adapters"
	"github.com/felipeazsantos/breeders/configuration"
)

var testApp application

func TestMain(m *testing.M) {
	testBackend := &adapters.TestBackend{}
	testAdapter := &adapters.RemoteService{Remote: testBackend}

	// call flag.Parse() here if TestMain uses flags
	testApp = application{
		App: configuration.New(nil, testAdapter),
	}

	os.Exit(m.Run())
}

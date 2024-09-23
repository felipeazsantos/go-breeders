package main

import (
	"os"
	"testing"

	"github.com/felipeazsantos/breeders/configuration"
	"github.com/felipeazsantos/breeders/models"
)

var testApp application

func TestMain(m *testing.M) {
	testBackend := &TestBackend{}
	testAdapter := &RemoteService{Remote: testBackend}

	// call flag.Parse() here if TestMain uses flags
	testApp = application{
		App:        configuration.New(nil),
		catService: testAdapter,
	}

	os.Exit(m.Run())
}

type TestBackend struct{}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		{
			ID:      1,
			Breed:   "Tomcat",
			Details: "Some details",
		},
	}
	return breeds, nil
}

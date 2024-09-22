package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/felipeazsantos/breeders/models"
)

// CatBreedsInterface is simply our target interface, which defines all the methods that
// any type which implements this interface must have
type CatBreedsInterface interface {
	GetAllCatBreeds() ([]*models.CatBreed, error)
}

// RemoteService is the Adaptor type. It embeds a DataInterface interface
// (which is critical to the pattern)
type RemoteService struct {
	Remote CatBreedsInterface
}

// GetAllCatBreeds is the function on RemoteService which lest us
// call any adaptor which implements the DataInterface type
func (rs *RemoteService) GetAllCatBreeds() ([]*models.CatBreed, error) {
	return rs.Remote.GetAllCatBreeds()
}

// JSONBackend is the JSON adaptee, which needs to satisfy the CatBreedsInterface by
// have a GetAllCatBreeds method
type JSONBackend struct{}

func (jb *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var breeds []*models.CatBreed
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds, nil
}

type XMLBackend struct{}

func (xb *XMLBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/xml")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type catBreeds struct {
		XMLName struct{}           `xml:"cat-breeds"`
		Breeds  []*models.CatBreed `xml:"cat-breed"`
	}

	var breeds catBreeds
	err = xml.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds.Breeds, nil
}

package pets

import (
	"errors"
	"fmt"
	"log"

	"github.com/felipeazsantos/breeders/configuration"
	"github.com/felipeazsantos/breeders/models"
)

// AnimalInterface is the interface for the types we will return from our abstract factory
// We'll create two types which implement this interface: DogFromFactory and CatFromFactory.
// In order to implement this interface, both types must implement the Show() method
type AnimalInterface interface {
	Show() string
}

// DogFromFactory is a type which implements the AnimalInterface, and embeds a dog.
type DogFromFactory struct {
	Pet *models.Dog
}

// Show is a method on the DogFromFactory type, which makes the type implement the AnimalInterface
func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", dff.Pet.Breed.Breed)
}

// CatFromFactory is a type which implements the AnimalInterface, and embeds a cat.
type CatFromFactory struct {
	Pet *models.Cat
}

// Show is a method on the CatFromFactory type, which makes the type implement the AnimalInterface
func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", cff.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
	newPetWithBreed(breed string) AnimalInterface
}

type DogAbstractFactory struct{}

func (df *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

func (df *DogAbstractFactory) newPetWithBreed(b string) AnimalInterface {
	app := configuration.GetInstance()
	breed, _ := app.Models.DogBreed.GetBreedByName(b)
	return &DogFromFactory{
		Pet: &models.Dog{Breed: *breed},
	}
}

type CatAbstractFactory struct{}

func (cf *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func (cf *CatAbstractFactory) newPetWithBreed(b string) AnimalInterface {
	app := configuration.GetInstance()
	breed, err := app.CatService.Remote.GetCatBreedByName(b)
	if err != nil {
		log.Println("Error getting cat breed by name", err)
		return nil
	}
	return &CatFromFactory{
		Pet: &models.Cat{Breed: *breed},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
	default:
		return nil, fmt.Errorf("invalid species: %s", species)
	}
}

func NewPetWithBreedFromAbstractFactory(species, breed string) (AnimalInterface, error) {
	switch species {
	case "dog":
		// return a dog with breed embedded
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPetWithBreed(breed)
		return dog, nil
	case "cat":
		// return a cat with breed embedded
		var catFactory CatAbstractFactory
		cat := catFactory.newPetWithBreed(breed)
		return cat, nil
	default:
		return nil, errors.New("invalid species supplied")
	}
}

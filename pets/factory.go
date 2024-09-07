package pets

import "github.com/felipeazsantos/breeders/models"

func NewPet(species string) *models.Pet {
	return &models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "no description entered yet",
		LifeSpan:    0,
	}
}

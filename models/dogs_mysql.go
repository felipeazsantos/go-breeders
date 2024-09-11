package models

import (
	"context"
	"time"
)

func (d *DogBreed) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, breed, weight_low_lbs, weight_high_lbs, 
				cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight, 
				lifespan, coalesce(details, '') as details, 
				coalesce(alternate_names, '') as alternate_names, 
				coalesce(geographic_origin, '') as geographic_origin 
			  FROM dog_breeds order by breed`

	var breeds []*DogBreed
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbreed DogBreed
		err := rows.Scan(&dbreed.ID, &dbreed.Breed, &dbreed.WeightLowLbs, &dbreed.WeightHighLbs,
			&dbreed.AverageWeight, &dbreed.Lifespan, &dbreed.Details, &dbreed.AlternateNames,
			&dbreed.GeographicOrigin)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, &dbreed)
	}

	return breeds, nil
}

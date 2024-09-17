package pets

import "errors"

type PetInterface interface {
	SetSpecies(s string) *Pet
	SetBreed(b string) *Pet
	SetMinWeight(m int) *Pet
	SetMaxWeight(m int) *Pet
	SetAverageWeight(a int) *Pet
	SetWeight(w int) *Pet
	SetDescription(d string) *Pet
	SetLifeSpan(l int) *Pet
	SetGeographicOrigin(g string) *Pet
	SetColor(c string) *Pet
	SetAge(a int) *Pet
	SetAgeEstimated(ae int) *Pet
}

func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}

func (p *Pet) SetBreed(b string) *Pet {
	p.Breeed = b
	return p
}

func (p *Pet) SetMinWeight(m int) *Pet {
	p.MinWeight = m
	return p
}

func (p *Pet) SetMaxWeight(m int) *Pet {
	p.MaxWeight = m
	return p
}

func (p *Pet) SetAverageWeight(a int) *Pet {
	p.AverageWeight = a
	return p
}

func (p *Pet) SetWeight(w int) *Pet {
	p.Weight = w
	return p
}

func (p *Pet) SetDescription(d string) *Pet {
	p.Description = d
	return p
}

func (p *Pet) SetLifeSpan(l int) *Pet {
	p.LifeSpan = l
	return p
}

func (p *Pet) SetGeographicOrigin(g string) *Pet {
	p.GeographicOrigin = g
	return p
}

func (p *Pet) SetColor(c string) *Pet {
	p.Color = c
	return p
}

func (p *Pet) SetAge(a int) *Pet {
	p.Age = a
	return p
}

func (p *Pet) SetAgeEstimated(ae int) *Pet {
	p.AgeEstimated = ae
	return p
}

func (p *Pet) Build() (*Pet, error) {
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("min weight cannot be greater than max weight")
	}

	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2

	return p, nil
}

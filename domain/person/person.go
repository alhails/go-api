package person

import (
	"errors"
	"time"

	"github.com/alhails/go-api/domain"
	"github.com/pborman/uuid"
)

// Created represents the creation of a Person entity
type Created struct {
	FirstName string
	LastName  string
}

// Updated represents the update of a Person's name
type Updated struct {
	FirstName   string
	LastName    string
	DateOfBirth *time.Time
}

// Person is the domain object for a Person entity
type Person struct {
	domain.Aggregate
	id          uuid.UUID
	firstName   string
	lastName    string
	dateOfBirth *time.Time
}

// NewPerson constructs a new empty Person entity
func NewPerson() *Person {
	p := &Person{}
	p.ApplyChange = p.apply
	return p
}

func (p *Person) apply(event interface{}) {
	switch e := event.(type) {
	case Created:
		p.id = uuid.NewUUID()
		p.firstName = e.FirstName
		p.lastName = e.LastName
	case Updated:
		p.firstName = e.FirstName
		p.lastName = e.LastName
		p.dateOfBirth = e.DateOfBirth
	}
}

func (p *Person) validateName(firstName string, lastName string) error {
	if len(firstName) == 0 {
		return errors.New("First name must be supplied")
	}
	if len(lastName) == 0 {
		return errors.New("Last name must be supplied")
	}

	return nil
}

// Create initialises a new Person entity
func (p *Person) Create(firstName string, lastName string) error {
	err := p.validateName(firstName, lastName)
	if err != nil {
		return err
	}

	e := Created{firstName, lastName}
	p.Apply(e, true)
	return nil
}

// Update updates a Person entity
func (p *Person) Update(firstName string, lastName string, dateOfBirth *time.Time) error {
	err := p.validateName(firstName, lastName)
	if err != nil {
		return err
	}
	e := Updated{firstName, lastName, dateOfBirth}
	p.Apply(e, true)
	return nil
}

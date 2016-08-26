package main

import (
	"errors"
	"time"

	"github.com/pborman/uuid"
)

// PersonCreated represents the creation of a Person entity
type PersonCreated struct {
	firstName string
	lastName  string
}

// PersonUpdated represents the update of a Person's name
type PersonUpdated struct {
	firstName   string
	lastName    string
	dateOfBirth *time.Time
}

// Person is the domain object for a Person entity
type Person struct {
	Aggregate
	id          uuid.UUID
	firstName   string
	lastName    string
	dateOfBirth *time.Time
}

// NewPerson constructs a new empty Person entity
func NewPerson() *Person {
	p := &Person{}
	p.applyChange = p.apply
	return p
}

func (p *Person) apply(event interface{}) {
	switch e := event.(type) {
	case PersonCreated:
		p.id = uuid.NewUUID()
		p.firstName = e.firstName
		p.lastName = e.lastName
	case PersonUpdated:
		p.firstName = e.firstName
		p.lastName = e.lastName
		p.dateOfBirth = e.dateOfBirth
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

	e := PersonCreated{firstName, lastName}
	p.Apply(e, true)
	return nil
}

func (p *Person) Update(firstName string, lastName string, dateOfBirth *time.Time) error {
	err := p.validateName(firstName, lastName)
	if err != nil {
		return err
	}
	e := PersonUpdated{firstName, lastName, dateOfBirth}
	p.Apply(e, true)
	return nil
}

func (p *Person) GetAge() (int, error) {
	if p.dateOfBirth == nil {
		return 0, errors.New("Date of birth is not set")
	}
	hours := int(time.Since(*p.dateOfBirth).Hours())
	return (24 * 365) / hours, nil
}

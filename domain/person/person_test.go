package person

import (
	"testing"
	"time"

	"github.com/alhails/go-api/domain"
)

func TestWhenPersonCreatedThenSuccesful(t *testing.T) {
	p := NewPerson()
	p.Create("Bob", "Smith")
	if p.firstName != "Bob" || p.lastName != "Smith" {
		t.Error("First and last name were not populated correctly")
	}
	changes := p.GetChanges()
	if len(changes) != 1 {
		t.Fail()
	}
	if _, ok := changes[0].(Created); !ok {
		t.Error("Created event was not logged")
	}
}

func TestWhenPersonCreatedWithNoFirstNameThenError(t *testing.T) {
	p := NewPerson()
	err := p.Create("", "Smith")
	if err == nil || err.Error() != "First name must be supplied" {
		t.Fail()
	}
}

func TestWhenPersonCreatedWithNoLastNameThenError(t *testing.T) {
	p := NewPerson()
	err := p.Create("Bob", "")
	if err == nil || err.Error() != "Last name must be supplied" {
		t.Fail()
	}
}

func TestWhenLoadFromSingleCreatedEventThenPopulatedSuccesfully(t *testing.T) {
	p := NewPerson()
	p.Load([]interface{}{Created{"Sarah", "Jones"}})
	if p.firstName != "Sarah" || p.lastName != "Jones" {
		t.Error("First and last name were not populated correctly")
	}
}

func TestWhenLoadFromMultipleEventsThenPopulatedSuccesfully(t *testing.T) {
	p := NewPerson()
	dob := time.Date(1975, time.November, 10, 23, 0, 0, 0, time.UTC)
	p.Load([]interface{}{Created{"Sarah", "Jones"}, Updated{"Sarah", "Smith", &dob}})
	if p.firstName != "Sarah" || p.lastName != "Smith" {
		t.Error("First and last name were not populated correctly")
	}
}

func TestWhen(t *testing.T) {
	p := NewPerson()
	p.Create("Bob", "Smith")
	dob := time.Date(1975, time.November, 10, 23, 0, 0, 0, time.UTC)
	p.Update("Bob", "Smith", &dob)
	r := domain.Repository{}
	r.Save(p)
	r.Persist()
}

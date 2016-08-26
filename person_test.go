package main

import (
	"testing"
)

func TestWhenPersonCreatedThenSuccesful(t *testing.T) {
	p := NewPerson()
	p.Create("Bob", "Smith")
	if p.firstName != "Bob" && p.lastName != "Smith" {
		t.Fail()
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

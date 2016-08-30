package domain

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const update = "U"
const insert = "I"

type repositoryItem struct {
	changeType string
	item       interface{}
}

// Repository manages the persistence of Person entities
type Repository struct {
	items []repositoryItem
}

// Save marks the supplied entity for insert
func (r *Repository) Save(p interface{}) {
	r.items = append(r.items, repositoryItem{insert, p})
}

// Update marks the supplied entity for update
func (r *Repository) Update(p interface{}) {
	r.items = append(r.items, repositoryItem{update, p})
}

// Persist persists the contained entities to the data store
func (r *Repository) Persist() {
	for _, item := range r.items {
		agg := item.item.(isAggregate)
		for _, change := range agg.GetChanges() {
			json, _ := json.Marshal(change)
			fmt.Printf("%v %v\n", reflect.TypeOf(change), string(json))
		}
	}
}

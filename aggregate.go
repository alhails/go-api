package main

// Aggregate is the base type for aggregate root types
type Aggregate struct {
	version     int
	changes     []interface{}
	applyChange func(event interface{})
}

// Apply applies an event to an aggregate root
func (agg *Aggregate) Apply(event interface{}, isNew bool) error {
	if isNew {
		agg.changes = append(agg.changes, event)
	}
	agg.applyChange(event)
	return nil
}

// Load rehydates an aggregate from history
func (agg *Aggregate) Load(events []interface{}) {
	for _, event := range events {
		agg.Apply(event, false)
		agg.version++
	}
}

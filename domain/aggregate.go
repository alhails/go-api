package domain

type isAggregate interface {
	GetVersion() int
	GetChanges() []interface{}
}

// Aggregate provides basic functionality for aggregate root types
type Aggregate struct {
	version     int
	changes     []interface{}
	ApplyChange func(event interface{})
}

// Apply applies an event to an aggregate root
func (agg *Aggregate) Apply(event interface{}, isNew bool) error {
	if isNew {
		agg.changes = append(agg.changes, event)
	}
	agg.ApplyChange(event)
	agg.version++
	return nil
}

// Load rehydates an aggregate from history
func (agg *Aggregate) Load(events []interface{}) {
	for _, event := range events {
		agg.Apply(event, false)
		agg.version++
	}
}

// GetVersion returns the current version of the aggregate
func (agg *Aggregate) GetVersion() int {
	return agg.version
}

// GetChanges returns the current list of uncommitted changes to the aggregate
func (agg *Aggregate) GetChanges() []interface{} {
	cpy := make([]interface{}, len(agg.changes), cap(agg.changes))
	copy(cpy, agg.changes)
	return cpy
}

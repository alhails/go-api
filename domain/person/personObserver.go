package person

import (
	"errors"
	"time"
)

func (p *Person) GetAge() (int, error) {
	if p.dateOfBirth == nil {
		return 0, errors.New("Date of birth is not set")
	}
	hours := int(time.Since(*p.dateOfBirth).Hours())
	return (24 * 365) / hours, nil
}

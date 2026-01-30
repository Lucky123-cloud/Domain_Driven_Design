package domain

import "time"

type Patient struct {
	ID          string
	Identifiers map[string]string
	FullName    string
	Gender      string
	BirthDate   *time.Time
	Active      bool
}

package fhir

type Patient struct {
	ResourceType ResourceType `json:"resourceType"`
	ID           string       `json:"id,omitempty"`

	Identifiers []Identifier `json:"identifier,omitempty"`
	Names       []HumanName  `json:"name,omitempty"`

	Gender    Gender `json:"gender,omitempty"`
	BirthDate Date   `json:"birthDate,omitempty"`
	Active    bool   `json:"active"`
}

type Identifier struct {
	System string `json:"system,omitempty"`
	Value  string `json:"value,omitempty"`
}

type HumanName struct {
	Family string   `json:"family,omitempty"`
	Given  []string `json:"given,omitempty"`
}

type ResourceType string

const (
	ResourceTypePatient ResourceType = "Patient"
)

type Gender string

const (
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
	GenderOther   Gender = "other"
	GenderUnknown Gender = "unknown"
)

type Date string

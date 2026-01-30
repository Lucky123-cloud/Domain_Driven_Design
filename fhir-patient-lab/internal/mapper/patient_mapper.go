package mapper

import (
	"strings"
	"time"

	"github.com/lucky123-cloud/fhir-patient-lab/internal/domain"
	"github.com/lucky123-cloud/fhir-patient-lab/internal/fhir"
)

func ToDomainPatient(p *fhir.Patient) *domain.Patient {
	identifiers := make(map[string]string)
	for _, id := range p.Identifier {
		identifiers[id.System] = id.Value
	}

	fullName := ""
	if len(p.Name) > 0 {
		fullName = strings.Join(p.Name[0].Given, " ") + " " + p.Name[0].Family
	}

	var birthDate *time.Time
	if p.BirthDate != "" {
		if t, err := time.Parse("2006-01-02", p.BirthDate); err == nil {
			birthDate = &t
		}
	}

	return &domain.Patient{
		ID:          p.ID,
		Identifiers: identifiers,
		FullName:    strings.TrimSpace(fullName),
		Gender:      p.Gender,
		BirthDate:   birthDate,
		Active:      p.Active,
	}
}

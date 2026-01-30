package main

import (
	"fmt"
	"log"

	"github.com/lucky123-cloud/fhir-patient-lab/internal/fhir"
	"github.com/lucky123-cloud/fhir-patient-lab/internal/mapper"
)

func main() {
	client := fhir.NewClient("https://r4.smarthealthit.org")

	// Try different IDs if one fails
	fhirPatient, err := client.GetPatientByID("smart-1288992")
	if err != nil {
		log.Fatal(err)
	}

	domainPatient := mapper.ToDomainPatient(fhirPatient)

	fmt.Println("Patient loaded successfully")
	fmt.Printf("ID: %s\n", domainPatient.ID)
	fmt.Printf("Name: %s\n", domainPatient.FullName)
	fmt.Printf("Gender: %s\n", domainPatient.Gender)
	fmt.Printf("Identifiers: %+v\n", domainPatient.Identifiers)
}

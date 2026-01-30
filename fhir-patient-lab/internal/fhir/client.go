package fhir

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) GetPatientByID(id string) (*Patient, error) {
	url := fmt.Sprintf("%s/Patient/%s", c.BaseURL, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("FHIR returned status %d", resp.StatusCode)
	}

	var patient Patient
	if err := json.NewDecoder(resp.Body).Decode(&patient); err != nil {
		return nil, err
	}

	return &patient, nil
}

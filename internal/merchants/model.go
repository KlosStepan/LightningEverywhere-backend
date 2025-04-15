package models

import "lightningeverywhere_backend/internal/base"

type Merchant struct {
	Type     string `json:"type"` // Always "Feature"
	Geometry struct {
		Type        string     `json:"type"`        // Always "Point"
		Coordinates [2]float64 `json:"coordinates"` // [lon, lat]
	} `json:"geometry"`
	Properties struct {
		base.BaseEntity          // Embedded: ID, Owner, Visible, Name, Description
		Images          []string `json:"images"`
		Address         struct {
			Address    string `json:"address"`
			City       string `json:"city"`
			PostalCode string `json:"postalCode"`
		} `json:"address"`
		Tags    []string `json:"tags"`
		Socials []Social `json:"socials"`
	} `json:"properties"`
}

type Social struct {
	Network string `json:"network"` // "web" | "facebook" | "instagram" | "twitter"
	Label   string `json:"label"`
	Link    string `json:"link"`
}

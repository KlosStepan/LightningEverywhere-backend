package merchant

import "lightningeverywhere_backend/internal/base"

type Merchant struct {
	Type       string         `json:"type"` // Always "Feature"
	Geometry   Geometry       `json:"geometry"`
	Properties MerchantFields `json:"properties"`
}

type Geometry struct {
	Type        string     `json:"type"`        // Always "Point"
	Coordinates [2]float64 `json:"coordinates"` // [lon, lat]
}

type MerchantFields struct {
	base.BaseEntity `bson:",inline"` // ID, Owner, Visible, Name, Description
	Images          []string         `json:"images"`
	Address         MerchantAddress  `json:"address"`
	Tags            []string         `json:"tags"`
	Socials         []Social         `json:"socials"`
}

type MerchantAddress struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
}

type Social struct {
	Network string `json:"network"` // "web" | "facebook" | "instagram" | "twitter"
	Label   string `json:"label"`
	Link    string `json:"link"`
}

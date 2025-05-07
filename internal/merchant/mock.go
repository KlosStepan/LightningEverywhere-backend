package merchant

import (
	"lightningeverywhere_backend/internal/base"
	"lightningeverywhere_backend/internal/models"
)

func MockMerchants() []models.Merchant {
	return []models.Merchant{
		{
			Type: "Feature",
			Geometry: models.Geometry{
				Type:        "Point",
				Coordinates: [2]float64{14.4483471, 50.1033561},
			},
			Properties: models.Properties{
				BaseEntity: base.BaseEntity{
					ID:          "63977929-fc0e-4695-9a04-3156e9d24c54",
					Owner:       ptr("EM6jd7CDU4PdHgF7LJTTvyMPNrJ3"),
					Visible:     true,
					Name:        "Paralelní Polis",
					Description: "lorem ipsum 2",
				},
				Images: []string{
					"https://firebasestorage.googleapis.com/v0/b/lightning-everywhere.firebasestorage.app/o/merchants-photos%2FgnO7T0wxtYB6jGRsztOS-polis1.jpeg?alt=media&token=347c84ab-d120-4ef0-b9d4-2d52e05ed400",
					"https://firebasestorage.googleapis.com/v0/b/lightning-everywhere.firebasestorage.app/o/merchants-photos%2FgnO7T0wxtYB6jGRsztOS-polis2.jpeg?alt=media&token=8042e4cc-16bc-4e59-af9a-29d8d409d7c5",
					"https://firebasestorage.googleapis.com/v0/b/lightning-everywhere.firebasestorage.app/o/merchants-photos%2FgnO7T0wxtYB6jGRsztOS-polis3.jpeg?alt=media&token=895fb62e-86de-4ca1-8427-ef45ae8cc6dd",
					"https://firebasestorage.googleapis.com/v0/b/lightning-everywhere.firebasestorage.app/o/merchants-photos%2FgnO7T0wxtYB6jGRsztOS-polis4.jpeg?alt=media&token=9692ebc4-0414-4be3-b9a8-e1ce2c527bd5",
				},
				Address: models.Address{
					Address:    "Delnicka 43",
					City:       "Praha 7",
					PostalCode: "170 00",
				},
				Tags: []string{"Shops", "Services"},
				Socials: []models.Social{
					{Network: "web", Label: "Web", Link: "https://www.paralelnipolis.com"},
					{Network: "facebook", Label: "FB", Link: "https://www.facebook.com/paralelnipolis"},
					{Network: "instagram", Label: "IG", Link: "https://www.instagram.com/paralelnipolis"},
					{Network: "twitter", Label: "X", Link: "https://www.twitter.com/paralelnipolis"},
				},
			},
		},
		{
			Type: "Feature",
			Geometry: models.Geometry{
				Type:        "Point",
				Coordinates: [2]float64{14.4440644, 50.0719584},
			},
			Properties: models.Properties{
				BaseEntity: base.BaseEntity{
					ID:          "10417593-72da-4ba3-ace0-93f6ba676b9e",
					Owner:       ptr("EM6jd7CDU4PdHgF7LJTTvyMPNrJ3"),
					Visible:     true,
					Name:        "Blue Vegan Pig Shop",
					Description: "lorem ipsum",
				},
				Images: []string{
					"https://firebasestorage.googleapis.com/v0/b/lightning-everywhere.firebasestorage.app/o/merchants-photos%2FkxFoOlquyi7iP7xtkqEJ-image-1-5.png?alt=media&token=1bbc8bda-abd4-4524-9385-419712e6c409",
					"https://firebasestorage.googleapis.com/v0/b/lightning-everywhere.firebasestorage.app/o/merchants-photos%2FkxFoOlquyi7iP7xtkqEJ-bluepig2.jpg?alt=media&token=9b1232b9-4f36-4d94-9551-a9777795b534",
					"https://firebasestorage.googleapis.com/v0/b/lightning-everywhere.firebasestorage.app/o/merchants-photos%2FkxFoOlquyi7iP7xtkqEJ-bluepig3.jpg?alt=media&token=0fefeee5-bd13-455e-a5f8-0cb45fc6ccda",
					"https://firebasestorage.googleapis.com/v0/b/lightning-everywhere.firebasestorage.app/o/merchants-photos%2FkxFoOlquyi7iP7xtkqEJ-bluepig4.jpg?alt=media&token=c3986452-b63f-44e5-bac2-d458681ccbe5",
				},
				Address: models.Address{
					Address:    "Štefánikova 6",
					City:       "Praha 5",
					PostalCode: "150 00",
				},
				Tags: []string{"Food & Drinks", "Shops"},
				Socials: []models.Social{
					{Network: "web", Label: "Web", Link: "https://www.blueveganpigshop.com"},
					{Network: "facebook", Label: "FB", Link: "https://www.facebook.com/blueveganpigshop"},
					{Network: "instagram", Label: "IG", Link: "https://www.instagram.com/blueveganpigshop"},
					{Network: "twitter", Label: "X", Link: "https://www.twitter.com/blueveganpigshop"},
				},
			},
		},
	}
}

func ptr[T any](v T) *T {
	return &v
}

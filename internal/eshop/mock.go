package eshop

import "lightningeverywhere_backend/internal/base"

func MockEshops() []Eshop {
	return []Eshop{
		{
			BaseEntity: base.BaseEntity{
				ID:          "54ae38db-2a48-42a9-a569-e9e48ea90d46",
				Owner:       ptr("EM6jd7CDU4PdHgF7LJTTvyMPNrJ3"),
				Visible:     true,
				Name:        "Alza.cz",
				Description: "Největší prodejce elektroniky v ČR",
			},
			Logo:    "https://firebasestorage.googleapis.com/v0/b/lightning-everywhere.firebasestorage.app/o/eshop-logos%2FLfcSSXzFXubrti5lnb4P-alza.png?alt=media&token=e5f83d6d-f313-4ef0-8f6d-ae7f5606d1e3",
			Country: "CZ",
			URL:     "https://www.alza.cz",
		},
		{
			BaseEntity: base.BaseEntity{
				ID:          "e20d248d-69cc-482d-9ce1-0a15cbe3bc2a",
				Owner:       ptr("EM6jd7CDU4PdHgF7LJTTvyMPNrJ3"),
				Visible:     true,
				Name:        "Kafe za bitcoin",
				Description: "Kafe, které si za peníze nekoupíš",
			},
			Logo:    "https://firebasestorage.googleapis.com/v0/b/lightning-everywhere.firebasestorage.app/o/eshop-logos%2FcnPF09Tf04V4CC3UVSbx-Screenshot_20250321_011958.png?alt=media&token=621f1835-4676-4446-805b-8eca4264c6c9",
			Country: "CZ",
			URL:     "https://www.kafezabitcoin.cz/",
		},
		{
			BaseEntity: base.BaseEntity{
				ID:          "b25efd17-3a50-4513-9dad-af47bbcfb687",
				Owner:       ptr("EM6jd7CDU4PdHgF7LJTTvyMPNrJ3"),
				Visible:     true,
				Name:        "Bitcoinovej Kanál ",
				Description: "Kicomuv merch Kicomuv merch Kicomuv merch",
			},
			Logo:    "https://firebasestorage.googleapis.com/v0/b/lightning-everywhere.firebasestorage.app/o/eshop-logos%2FthNQXwffucpx8oS3OakB-Screenshot_20250321_012223.png?alt=media&token=c2cebca5-f985-447a-9759-0936b8778e41",
			Country: "CZ",
			URL:     "https://bitcoinovejkanal.cz/eshop/",
		},
	}
}

func ptr[T any](v T) *T {
	return &v
}

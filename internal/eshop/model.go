package eshop

import "lightningeverywhere_backend/internal/base"

type Eshop struct {
	base.BaseEntity
	Logo    string `bson:"logo" json:"logo"`
	Country string `bson:"country" json:"country"`
	URL     string `bson:"url" json:"url"`
}

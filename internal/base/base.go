package base

type BaseEntity struct {
	ID          string  `bson:"id" json:"id"`
	Owner       *string `bson:"owner,omitempty" json:"owner"`
	Visible     bool    `bson:"visible" json:"visible"`
	Name        string  `bson:"name" json:"name"`
	Description string  `bson:"description" json:"description"`
}

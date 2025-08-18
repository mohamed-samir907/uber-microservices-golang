package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFare struct {
	ID                primitive.ObjectID
	UserId            string
	PackageSlug       string // e.g., "van", "luxury", "economy", "sedan"
	TotalPriceInCents float64
}

type RideFareRepository interface {
}

type RideFareService interface {
}

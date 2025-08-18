package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFare struct {
	ID                primitive.ObjectID `json:"id"`
	UserId            string             `json:"user_id"`
	PackageSlug       string             `json:"package_slug"` // e.g., "van", "luxury", "economy", "sedan"
	TotalPriceInCents float64            `json:"total_price_in_cents"`
}

type RideFareRepository interface {
}

type RideFareService interface {
}

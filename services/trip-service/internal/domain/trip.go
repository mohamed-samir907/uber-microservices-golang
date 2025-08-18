package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Trip struct {
	ID       primitive.ObjectID `json:"id"`
	UserId   string             `json:"user_id"`
	Status   string             `json:"status"`
	RideFare *RideFare          `json:"ride_fare"`
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *Trip) (*Trip, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFare) (*Trip, error)
}

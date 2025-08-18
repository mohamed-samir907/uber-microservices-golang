package service

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tripService struct {
	repo domain.TripRepository
}

func NewTripService(repo domain.TripRepository) domain.TripService {
	return &tripService{
		repo: repo,
	}
}

func (s *tripService) CreateTrip(ctx context.Context, fare *domain.RideFare) (*domain.Trip, error) {
	trip := domain.Trip{
		ID:       primitive.NewObjectID(),
		UserId:   fare.UserId,
		Status:   "pending",
		RideFare: fare,
	}

	return s.repo.CreateTrip(ctx, &trip)
}

package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
)

type inmemRepository struct {
	trips map[string]*domain.Trip
	fares map[string]*domain.RideFare
}

func NewInMemoryRepository() domain.TripRepository {
	return &inmemRepository{
		trips: make(map[string]*domain.Trip),
		fares: make(map[string]*domain.RideFare),
	}
}

func (r *inmemRepository) CreateTrip(ctx context.Context, trip *domain.Trip) (*domain.Trip, error) {
	r.trips[trip.ID.String()] = trip
	return trip, nil
}

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"

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

func (s *tripService) GetRoute(ctx context.Context, pickup, destination *types.Coordinate) (*types.OsrmApiResponse, error) {
	// Using OSRM to get the route
	url := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=full&geometries=geojson",
		pickup.Longitude, pickup.Latitude,
		destination.Longitude, destination.Latitude,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var route types.OsrmApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&route); err != nil {
		return nil, err
	}

	return &route, nil
}

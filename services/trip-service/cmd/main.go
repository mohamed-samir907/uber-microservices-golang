package main

import (
	"context"
	"log/slog"
	"os"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var (
// 	httpAddr = env.GetString("HTTP_ADDR", ":8083")
// )

func main() {
	inmemRepo := repository.NewInMemoryRepository()
	svc := service.NewTripService(inmemRepo)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	fare := domain.RideFare{
		ID:     primitive.NewObjectID(),
		UserId: "user123",
	}

	trip, err := svc.CreateTrip(ctx, &fare)
	if err != nil {
		slog.Error("Failed to create trip", "error", err)
		os.Exit(1)
	}

	slog.Info("Trip created successfully", "tripID", trip.ID)

	// keep the service running
	for {
		time.Sleep(time.Second * 10)
	}
}

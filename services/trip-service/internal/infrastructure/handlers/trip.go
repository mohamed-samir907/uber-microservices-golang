package handlers

import (
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tripHandler struct {
	svc domain.TripService
}

func NewTripHandler(svc domain.TripService) *tripHandler {
	return &tripHandler{
		svc: svc,
	}
}

func (h *tripHandler) CreateTrip(w http.ResponseWriter, r *http.Request) {
	fare := domain.RideFare{
		ID:     primitive.NewObjectID(),
		UserId: "user123",
	}

	trip, err := h.svc.CreateTrip(r.Context(), &fare)
	if err != nil {
		utils.ResponseJson(w, http.StatusBadRequest, nil, utils.APIError(err.Error(), http.StatusBadRequest))
		return
	}

	utils.ResponseJson(w, http.StatusCreated, trip, nil)
}

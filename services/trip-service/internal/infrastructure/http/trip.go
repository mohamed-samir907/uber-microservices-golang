package http

import (
	"encoding/json"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
	"ride-sharing/shared/utils"
)

type tripHandler struct {
	svc domain.TripService
}

func NewTripHandler(svc domain.TripService) *tripHandler {
	return &tripHandler{
		svc: svc,
	}
}

type previewTripRequest struct {
	UserId      string           `json:"userId"`
	Pikup       types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (h *tripHandler) TripPreview(w http.ResponseWriter, r *http.Request) {
	var req previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseJson(w, http.StatusBadRequest, nil, utils.APIError("Invalid request body", http.StatusBadRequest))
		return
	}

	resp, err := h.svc.GetRoute(r.Context(), &req.Pikup, &req.Destination)
	if err != nil {
		utils.ResponseJson(w, http.StatusBadRequest, nil, utils.APIError(err.Error(), http.StatusBadRequest))
		return
	}

	utils.ResponseJson(w, http.StatusCreated, resp, nil)
}

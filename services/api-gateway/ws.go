package main

import (
	"log/slog"
	"net/http"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleRidersWebSocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.APIError(err.Error(), http.StatusInternalServerError)
		return
	}
	defer c.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		utils.APIError("userID is required", http.StatusBadRequest)
		return
	}

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			utils.APIError(err.Error(), http.StatusInternalServerError)
			break
		}

		slog.Info("Received message from rider", "userID", userID, "message", string(msg))
	}
}

func handleDriversWebSocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.APIError(err.Error(), http.StatusInternalServerError)
		return
	}
	defer c.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		utils.APIError("userID is required", http.StatusBadRequest)
		return
	}

	packageSlug := r.URL.Query().Get("packageSlug")
	if packageSlug == "" {
		utils.APIError("packageSlug is required", http.StatusBadRequest)
		return
	}

	type Driver struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		ProfilePicture string `json:"profilePicture"`
		CarPlate       string `json:"carPlate"`
		PackageSlug    string `json:"packageSlug"`
	}

	msg := contracts.WSMessage{
		Type: "driver.cmd.register",
		Data: Driver{
			Id:             userID,
			Name:           "Mohamed Samir",
			ProfilePicture: utils.GetRandomAvatar(1),
			CarPlate:       "ABC123",
			PackageSlug:    packageSlug,
		},
	}

	if err := c.WriteJSON(msg); err != nil {
		utils.APIError(err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			utils.APIError(err.Error(), http.StatusInternalServerError)
			break
		}

		slog.Info("Received message from driver", "userID", userID, "packageSlug", packageSlug, "message", string(msg))
	}
}

package main

import "ride-sharing/shared/types"

type previewTripRequest struct {
	UserId      string           `json:"userId"`
	Pikup       types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

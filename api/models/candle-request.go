package models

import "time"

type CandleRequest struct {
	Resolution CandleResolution `json:"resolution" validate:"candle_resolution"`
	From       *time.Time       `json:"from" validate:"required"`
	To         *time.Time       `json:"to" validate:"required,gtefield=From"`
}

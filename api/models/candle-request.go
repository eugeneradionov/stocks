package models

type CandleRequest struct {
	Resolution CandleResolution `json:"resolution" validate:"candle_resolution"`
	From       *JSONTime        `json:"from" validate:"required,gte"`
	To         *JSONTime        `json:"to" validate:"required,gte,ltefield=From"`
}

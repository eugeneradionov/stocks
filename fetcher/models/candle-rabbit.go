package models

type CandleRPCRequest struct {
	Symbol     string           `json:"symbol"`
	Resolution CandleResolution `json:"resolution"`

	// From - UNIX time
	From int64 `json:"from"`

	// To - UNIX time
	To int64 `json:"to"`
}

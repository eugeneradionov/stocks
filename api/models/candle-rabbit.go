package models

type CandleRPCRequest struct {
	Symbol     string           `json:"symbol"`
	Resolution CandleResolution `json:"resolution"`

	// From - UNIX time
	From int64 `json:"from"`

	// To - UNIX time
	To int64 `json:"to"`
}

type CandleRabbit struct {
	// Open - open prices
	Open []float32 `json:"o"`

	// Close - close prices
	Close []float32 `json:"c"`

	// High - high prices
	High []float32 `json:"h"`

	// Low - low prices
	Low []float32 `json:"l"`

	// Volume - volume data
	Volume []int `json:"v"`

	// List of timestamps in UNIX format
	Timestamps []int64 `json:"t"`

	// Status of the response.
	// This field can either be `ok` or `no_data`.
	Status string `json:"s"`
}

func (cr CandleRabbit) ToCandleData() CandleData {
	return CandleData{
		Open:       cr.Open,
		Close:      cr.Close,
		High:       cr.High,
		Low:        cr.Low,
		Volume:     cr.Volume,
		Timestamps: cr.Timestamps,
	}
}

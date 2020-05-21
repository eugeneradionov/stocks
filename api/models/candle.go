package models

import "time"

type CandleResolution string

const (
	Minute         CandleResolution = "1"
	FiveMinutes    CandleResolution = "5"
	FifteenMinutes CandleResolution = "15"
	Day            CandleResolution = "D"
	Week           CandleResolution = "W"
	Month          CandleResolution = "M"
)

var CandleResolutionMap = map[CandleResolution]struct{}{
	Minute:         {},
	FiveMinutes:    {},
	FifteenMinutes: {},
	Day:            {},
	Week:           {},
	Month:          {},
}

type Candle struct {
	Symbol     string           `json:"symbol"`
	Resolution CandleResolution `json:"resolution"`
	From       time.Time        `json:"from"`
	To         time.Time        `json:"to"`

	Data CandleData `json:"data"`
}

type CandleData struct {
	// Open - open prices
	Open []float32 `json:"open"`

	// Close - close prices
	Close []float32 `json:"close"`

	// High - high prices
	High []float32 `json:"high"`

	// Low - low prices
	Low []float32 `json:"low"`

	// Volume - volume data
	Volume []int `json:"volume"`

	// List of timestamps in UNIX format
	Timestamps []int64 `json:"timestamps"`
}

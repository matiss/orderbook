package orderbook

import (
	"time"
)

// DepthEvent define depth event
type DepthEvent struct {
	Symbol        string
	FirstUpdateID int64  `json:"firstUpdateID"`
	FinalUpdateID int64  `json:"finalUpdateId"`
	Bids          []*Bid `json:"bids"`
	Asks          []*Ask `json:"asks"`
	Timestamp     time.Time
}

// Bid define bid info with price and quantity
type Bid struct {
	Price    int64
	Quantity int64
	Delete   bool
}

// Ask define ask info with price and quantity
type Ask struct {
	Price    int64
	Quantity int64
	Delete   bool
}

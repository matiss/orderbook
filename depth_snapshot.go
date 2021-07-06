package orderbook

// DepthSnapshot struct
type DepthSnapshot struct {
	LastUpdateID int64  `json:"lastUpdateId"`
	Asks         []*Ask `json:"asks"`
	Bids         []*Bid `json:"bids"`
}

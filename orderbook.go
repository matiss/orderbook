package orderbook

import (
	"fmt"
	"time"
)

// OrderBook struct
type OrderBook struct {
	Symbol         string
	LastUpdateID   int64
	UpdatedAt      time.Time
	PruneThreshold int

	Asks *OrderBookList
	Bids *OrderBookList

	loaded bool
}

// NewOrderBook creates new struct instance of *OrderBook
func NewOrderBook(symbol string, asks, bids *OrderBookList, pruneThreshold int) *OrderBook {
	return &OrderBook{
		Symbol:         symbol,
		Asks:           asks,
		Bids:           bids,
		PruneThreshold: pruneThreshold,
		// mu:             new(sync.Mutex),
	}
}

// ProcessSnapshot processes depth snapshot
func (ob *OrderBook) ProcessSnapshot(snapshot *DepthSnapshot, eventBuffer []*DepthEvent) {
	ob.UpdatedAt = time.Now()
	ob.LastUpdateID = snapshot.LastUpdateID

	// Asks
	for _, ask := range snapshot.Asks {
		ob.Asks.UpdateOrAddAsc(ask.Price, ask.Quantity)
	}

	// Bids
	for _, bid := range snapshot.Bids {
		ob.Bids.UpdateOrAddDesc(bid.Price, bid.Quantity)
	}

	// Process buffered events
	for _, event := range eventBuffer {
		if event.FinalUpdateID <= ob.LastUpdateID {
			// Ignore
			continue
		}

		// Asks
		for _, ask := range event.Asks {
			ob.Asks.UpdateOrAddAsc(ask.Price, ask.Quantity)
		}

		// Bids
		for _, bid := range event.Bids {
			ob.Bids.UpdateOrAddDesc(bid.Price, bid.Quantity)
		}

		ob.LastUpdateID = event.FinalUpdateID
	}

	// Mark as loaded
	ob.loaded = true
}

// ProcessEvent processes depth update event
func (ob *OrderBook) ProcessEvent(event *DepthEvent) error {
	if !ob.loaded {
		return fmt.Errorf("No orderbook to update for symbol: %s", ob.Symbol)
	}

	// Validate and process event
	if (event.FirstUpdateID <= ob.LastUpdateID+1) && (event.FinalUpdateID >= ob.LastUpdateID+1) {
		ob.UpdatedAt = time.Now()
		ob.LastUpdateID = event.FinalUpdateID

		// Process Asks
		for _, askUpdate := range event.Asks {
			if askUpdate.Delete {
				ob.Asks.Remove(askUpdate.Price)
			} else {
				ob.Asks.UpdateOrAddAsc(askUpdate.Price, askUpdate.Quantity)
			}
		}

		// Process Bids
		for _, bidUpdate := range event.Bids {
			if bidUpdate.Delete {
				ob.Bids.Remove(bidUpdate.Price)
			} else {
				ob.Bids.UpdateOrAddDesc(bidUpdate.Price, bidUpdate.Quantity)
			}
		}

		// Prune lists
		if ob.Asks.len > ob.PruneThreshold {
			ob.Asks.Prune(ob.PruneThreshold)
		}

		if ob.Bids.len > ob.PruneThreshold {
			ob.Bids.Prune(ob.PruneThreshold)
		}
	} else {
		return fmt.Errorf("Invalid event(%s): %d <= %d >= %d", event.Symbol, event.FirstUpdateID, ob.LastUpdateID+1, event.FinalUpdateID)
	}

	return nil
}

// Clear cache
func (ob *OrderBook) Clear() {
	ob.LastUpdateID = 0
	ob.Asks = &OrderBookList{
		len:  0,
		head: nil,
	}
	ob.Bids = &OrderBookList{
		len:  0,
		head: nil,
	}
	ob.UpdatedAt = time.Now()
	ob.loaded = false
}

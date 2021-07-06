package orderbook

import (
	"errors"

	"github.com/shopspring/decimal"
)

// OrderBookAskConversion bid ask conversion
func (ob *OrderBook) OrderBookAskConversion(amount decimal.Decimal) (decimal.Decimal, error) {
	if amount.IsZero() {
		return decimal.Decimal{}, nil
	}

	amountFrom := amount
	amountTo := decimal.NewFromInt(0)

	var price decimal.Decimal
	var quantity decimal.Decimal
	var exchangeableAmount decimal.Decimal

	// Iterate over asks
	ask, _ := ob.Asks.Front()
	for ask != nil {
		price = decimal.New(ask.Price, PriceDecimalExp)
		quantity = decimal.New(ask.Size, SizeDecimalExp)

		exchangeableAmount = quantity.Mul(price)
		if exchangeableAmount.LessThan(amountFrom) {
			// Fill
			amountFrom = amountFrom.Sub(exchangeableAmount)
			amountTo = amountTo.Add(quantity)
		} else {
			// Last fill
			return amountTo.Add(amountFrom.Div(price)), nil
		}

		// Get next
		ask = ask.next
	}

	// Too shallow orderbook to fill order
	return decimal.Decimal{}, errors.New("Too shallow Ask depth to fill order")
}

// OrderBookBidConversion bid conversion
func (ob *OrderBook) OrderBookBidConversion(amount decimal.Decimal) (decimal.Decimal, error) {
	if amount.IsZero() {
		return decimal.Decimal{}, nil
	}

	amountFrom := amount
	amountTo := decimal.NewFromInt(0)

	var price decimal.Decimal
	var quantity decimal.Decimal
	var exchangeableAmount decimal.Decimal

	// Iterate over bids
	bid, _ := ob.Bids.Front()
	for bid != nil {
		price = decimal.New(bid.Price, PriceDecimalExp)
		quantity = decimal.New(bid.Size, SizeDecimalExp)

		exchangeableAmount = quantity.Mul(price)
		if quantity.LessThan(amountFrom) {
			// Fill
			amountFrom = amountFrom.Sub(quantity)
			amountTo = amountTo.Add(exchangeableAmount)
		} else {
			// Last fill
			return amountTo.Add(amountFrom.Mul(price)), nil
		}

		// Get next
		bid = bid.next
	}

	// Too shallow orderbook to fill order
	return decimal.Decimal{}, errors.New("Too shallow Bid depth to fill order")
}

// OrderBooAskReverseConversion ask reverse conversion
func (ob *OrderBook) OrderBooAskReverseConversion(amount decimal.Decimal) (decimal.Decimal, error) {
	if amount.IsZero() {
		return decimal.Decimal{}, nil
	}

	amountFrom := amount
	amountTo := decimal.NewFromInt(0)

	var price decimal.Decimal
	var quantity decimal.Decimal
	var exchangeableAmount decimal.Decimal

	// Iterate over asks
	ask, _ := ob.Asks.Front()
	for ask != nil {
		price = decimal.New(ask.Price, PriceDecimalExp)
		quantity = decimal.New(ask.Size, SizeDecimalExp)

		exchangeableAmount = quantity.Mul(price)
		if quantity.LessThan(amountFrom) {
			// Fill
			amountFrom = amountFrom.Sub(quantity)
			amountTo = amountTo.Add(exchangeableAmount)
		} else {
			// Last fill
			return amountTo.Add(amountFrom.Mul(price)), nil
		}

		// Get next
		ask = ask.next
	}

	return decimal.Decimal{}, errors.New("Too shallow Ask depth to fill order")
}

// OrderBookBidReverseConversion bid reverse conversion
func (ob *OrderBook) OrderBookBidReverseConversion(amount decimal.Decimal) (decimal.Decimal, error) {
	if amount.IsZero() {
		return decimal.Decimal{}, nil
	}

	amountFrom := amount
	amountTo := decimal.NewFromInt(0)

	var price decimal.Decimal
	var quantity decimal.Decimal
	var exchangeableAmount decimal.Decimal

	// Iterate over bids
	bid, _ := ob.Bids.Front()
	for bid != nil {
		price = decimal.New(bid.Price, PriceDecimalExp)
		quantity = decimal.New(bid.Size, SizeDecimalExp)

		exchangeableAmount = quantity.Mul(price)
		if exchangeableAmount.LessThan(amountFrom) {
			// Fill
			amountFrom = amountFrom.Sub(exchangeableAmount)
			amountTo = amountTo.Add(quantity)
		} else {
			// Last fill
			return amountTo.Add(amountFrom.Div(price)), nil
		}

		// Get next
		bid = bid.next
	}

	return decimal.Decimal{}, errors.New("Too shallow Bid depth to fill order")
}

// GetBuyOrderBookDepthRequirement get orderbook depth fill requirement for buy order
func (ob *OrderBook) GetBuyOrderBookDepthRequirement(maxPrice decimal.Decimal, amount decimal.Decimal) int {
	i := 0

	var price decimal.Decimal
	var quantity decimal.Decimal
	var exchanged decimal.Decimal

	// Iterate over bids
	order, _ := ob.Asks.Front()
	for order != nil {
		price = decimal.New(order.Price, PriceDecimalExp)
		quantity = decimal.New(order.Size, SizeDecimalExp)

		// Check max price for first entry
		if i == 0 && price.GreaterThan(maxPrice) {
			return 0
		}

		exchanged = exchanged.Add(quantity)
		if exchanged.GreaterThanOrEqual(quantity) {
			return (i + 1)
		}

		// Get next
		order = order.next

		// Increment count
		i++
	}

	return i
}

// GetSellOrderBookDepthRequirement get orderbook depth fill requirement for sell order
func (ob *OrderBook) GetSellOrderBookDepthRequirement(minPrice decimal.Decimal, amount decimal.Decimal) int {
	i := 0

	var price decimal.Decimal
	var quantity decimal.Decimal
	var exchanged decimal.Decimal

	// Iterate over bids
	order, _ := ob.Bids.Front()
	for order != nil {
		price = decimal.New(order.Price, PriceDecimalExp)
		quantity = decimal.New(order.Size, SizeDecimalExp)

		// Check min price for first entry
		if i == 0 && price.LessThan(minPrice) {
			return 0
		}

		exchanged = exchanged.Add(quantity)
		if exchanged.GreaterThanOrEqual(quantity) {
			return (i + 1)
		}

		// Get next
		order = order.next

		// Increment count
		i++
	}

	return i
}

var two = decimal.NewFromInt(2)

// GetMarketPrice return current market price
func (ob *OrderBook) GetMarketPrice() (decimal.Decimal, error) {
	var price decimal.Decimal

	if ob.Asks == nil || ob.Bids == nil {
		return price, errors.New("Missing Bids/Asks for order book")
	}

	// Get first ask
	ask, err := ob.Asks.Front()
	if err != nil {
		return price, err
	}

	// Get first bid
	bid, err := ob.Bids.Front()
	if err != nil {
		return price, err
	}

	askPrice := decimal.New(ask.Price, PriceDecimalExp)
	bidPrice := decimal.New(bid.Price, PriceDecimalExp)

	// Calculate market price
	price = askPrice.Add(bidPrice).Div(two)

	return price, nil
}

// GetFirstAskPrice returns first ask price
func (ob *OrderBook) GetFirstAskPrice() (decimal.Decimal, error) {
	var price decimal.Decimal

	// Get first ask
	ask, err := ob.Asks.Front()
	if err != nil {
		return price, err
	}

	price = decimal.New(ask.Price, PriceDecimalExp)

	return price, nil
}

// GetFirstBidPrice returns first bid price
func (ob *OrderBook) GetFirstBidPrice() (decimal.Decimal, error) {
	var price decimal.Decimal

	// Get first bid
	bid, err := ob.Bids.Front()
	if err != nil {
		return price, err
	}

	price = decimal.New(bid.Price, PriceDecimalExp)

	return price, nil
}

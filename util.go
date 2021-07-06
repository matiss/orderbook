package orderbook

import (
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

const (
	// PriceDecimalExp satoshi
	PriceDecimalExp = -8
	// SizeDecimalExp 10^6
	SizeDecimalExp = -6
	// PercentageDecimalExp 10^4
	PercentageDecimalExp = -4
)

// Cached var for ROI
var oneHundred = decimal.NewFromInt(100)

// DecimalStringToSatoshi parses decimal string to int64 (satoshi)
func DecimalStringToSatoshi(value string) int64 {
	// Split string
	parts := strings.Split(value, ".")

	// Invalid string
	if len(parts) != 2 {
		return 0
	}

	var significandB strings.Builder

	significandB.WriteString(parts[0])
	significandB.WriteString(parts[1])

	// Validate part 1
	if len(parts[1]) < 8 {
		add := (8 - len(parts[1]))

		for i := add; i > 0; i-- {
			significandB.WriteString("0")
		}
	}

	significand, err := strconv.ParseInt(significandB.String(), 10, 64)
	if err != nil {
		return 0
	}

	return significand
}

// SatoshiToDecimalString converts int64 (satoshi) value to decimal string
func SatoshiToDecimalString(value int64) string {
	valueStr := strconv.FormatInt(value, 10)
	length := len(valueStr)

	if length < 8 {
		// Prepend zeros
		for i := (8 - length); i > 0; i-- {
			valueStr = "0" + valueStr
		}

		valueStr = ("0." + valueStr)
	} else if length == 8 {
		valueStr = ("0." + valueStr)
	} else {
		valueStr = valueStr[:length-8] + "." + valueStr[length-8:]
	}

	return valueStr
}

// DecimalStringToSize parses decimal string to int64 (Size)
func DecimalStringToSize(value string) int64 {
	satoshi := DecimalStringToSatoshi(value)

	if satoshi == 0 {
		return 0
	}

	return (satoshi / 100)
}

// SizeToDecimalString converts int64 (size) value to decimal string
func SizeToDecimalString(value int64) string {
	valueStr := strconv.FormatInt(value, 10)
	length := len(valueStr)

	if length < 6 {
		// Prepend zeros
		for i := (6 - length); i > 0; i-- {
			valueStr = "0" + valueStr
		}

		valueStr = ("0." + valueStr)
	} else if length == 6 {
		valueStr = ("0." + valueStr)
	} else {
		valueStr = valueStr[:length-6] + "." + valueStr[length-6:]
	}

	return valueStr
}

// DecimalStringToPercentageInt parses decimal string to int32 (Size)
func DecimalStringToPercentageInt(value string) int32 {
	satoshi := DecimalStringToSatoshi(value)

	if satoshi == 0 {
		return 0
	}

	return int32(satoshi / 10000)
}

// PercentageIntToDecimalString converts int32 (percentage) value to decimal string
func PercentageIntToDecimalString(value int32) string {
	valueStr := strconv.FormatInt(int64(value), 10)
	length := len(valueStr)

	if length < 4 {
		// Prepend zeros
		for i := (4 - length); i > 0; i-- {
			valueStr = "0" + valueStr
		}

		valueStr = ("0." + valueStr)
	} else if length == 4 {
		valueStr = ("0." + valueStr)
	} else {
		valueStr = valueStr[:length-4] + "." + valueStr[length-4:]
	}

	return valueStr
}

// DecimalToStatoshi converts (shifts) decimal.Decimal to int64 (satoshi)
func DecimalToStatoshi(value decimal.Decimal) int64 {
	dv := value.Shift(8)
	return dv.IntPart()
}

// DecimalToSize converts (shifts) decimal.Decimal to int64 (size)
func DecimalToSize(value decimal.Decimal) int64 {
	dv := value.Shift(6)
	return dv.IntPart()
}

// DecimalToPercentage converts (shifts) decimal.Decimal to int32 (percentage)
func DecimalToPercentage(value decimal.Decimal) int32 {
	dv := value.Shift(4)
	return int32(dv.IntPart())
}

// StatoshiToDecimal converts (shifts) decimal.Decimal to int64 (satoshi)
func StatoshiToDecimal(value int64) decimal.Decimal {
	return decimal.New(value, PriceDecimalExp)
}

// SizeToDecimal converts (shifts) decimal.Decimal to int64 (size)
func SizeToDecimal(value int64) decimal.Decimal {
	return decimal.New(value, SizeDecimalExp)
}

// PercentageToDecimal converts (shifts) decimal.Decimal to int32 (percentage)
func PercentageToDecimal(value int32) decimal.Decimal {
	return decimal.New(int64(value), PercentageDecimalExp)
}

// MathROI calculates ROI
func MathROI(spent, earned decimal.Decimal) decimal.Decimal {
	return earned.Sub(spent).Div(spent).Mul(oneHundred)
}

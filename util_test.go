package orderbook

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestDecimalStringToSatoshi(t *testing.T) {
	// Case 1
	var testValue string = "48.000012"
	var expectedResult int64 = 4800001200

	var result int64 = DecimalStringToSatoshi(testValue)

	if result != expectedResult {
		t.Errorf("Case 1. Parsing failed, expected: %d got: %d", expectedResult, result)
	}

	// Case 2
	testValue = "0.01633102"
	expectedResult = 1633102

	result = DecimalStringToSatoshi(testValue)

	if result != expectedResult {
		t.Errorf("Case 2. Parsing failed, expected: %d got: %d", expectedResult, result)
	}

	// Case 3
	testValue = "1633102"
	expectedResult = 0

	result = DecimalStringToSatoshi(testValue)

	if result != expectedResult {
		t.Errorf("Case 3. Parsing failed, expected: %d got: %d", expectedResult, result)
	}

	// Case 4
	testValue = "-48.000012"
	expectedResult = -4800001200

	result = DecimalStringToSatoshi(testValue)

	if result != expectedResult {
		t.Errorf("Case 4. Parsing failed, expected: %d got: %d", expectedResult, result)
	}
}

func TestSatoshiToDecimalString(t *testing.T) {
	// Case 1
	var testValue int64 = 4800001200
	var expectedResult string = "48.00001200"

	result := SatoshiToDecimalString(testValue)

	if result != expectedResult {
		t.Errorf("Case 1. Parsing failed, expected: %s got: %s", expectedResult, result)
	}

	// Case 2
	testValue = 10001200
	expectedResult = "0.10001200"

	result = SatoshiToDecimalString(testValue)

	if result != expectedResult {
		t.Errorf("Case 2. Parsing failed, expected: %s got: %s", expectedResult, result)
	}

	// Case 3
	testValue = 1200
	expectedResult = "0.00001200"

	result = SatoshiToDecimalString(testValue)

	if result != expectedResult {
		t.Errorf("Case 3. Parsing failed, expected: %s got: %s", expectedResult, result)
	}

	// Case 4
	testValue = -444800001200
	expectedResult = "-4448.00001200"

	result = SatoshiToDecimalString(testValue)

	if result != expectedResult {
		t.Errorf("Case 4. Parsing failed, expected: %s got: %s", expectedResult, result)
	}
}

func TestDecimalStringToSize(t *testing.T) {
	// Case 1
	var testValue string = "48.000012"
	var expectedResult int64 = 48000012

	var result int64 = DecimalStringToSize(testValue)

	if result != expectedResult {
		t.Errorf("Case 1. Parsing failed, expected: %d got: %d", expectedResult, result)
	}

	// Case 2
	testValue = "0.01633102"
	expectedResult = 16331

	result = DecimalStringToSize(testValue)

	if result != expectedResult {
		t.Errorf("Case 2. Parsing failed, expected: %d got: %d", expectedResult, result)
	}

	// Case 3
	testValue = "1633102"
	expectedResult = 0

	result = DecimalStringToSize(testValue)

	if result != expectedResult {
		t.Errorf("Case 3. Parsing failed, expected: %d got: %d", expectedResult, result)
	}

	// Case 4
	testValue = "-48.000012"
	expectedResult = -48000012

	result = DecimalStringToSize(testValue)

	if result != expectedResult {
		t.Errorf("Case 4. Parsing failed, expected: %d got: %d", expectedResult, result)
	}
}

func TestDecimalStringToPercentageInt(t *testing.T) {
	// Case 1
	var testValue string = "0.1"
	var expectedResult int32 = 1000

	var result int32 = DecimalStringToPercentageInt(testValue)

	if result != expectedResult {
		t.Errorf("Case 1. Parsing failed, expected: %d got: %d", expectedResult, result)
	}
}

func TestDecimalToStatoshi(t *testing.T) {
	// Case 1
	var testValue string = "48.000012"
	var expectedResult int64 = 4800001200

	dv, err := decimal.NewFromString(testValue)
	if err != nil {
		t.Error(err)
		return
	}

	var result int64 = DecimalToStatoshi(dv)

	if result != expectedResult {
		t.Errorf("Case 1. Parsing failed, expected: %d got: %d", expectedResult, result)
	}

	// Case 2
	testValue = "0.01633102"
	expectedResult = 1633102

	dv, err = decimal.NewFromString(testValue)
	if err != nil {
		t.Error(err)
		return
	}

	result = DecimalToStatoshi(dv)

	if result != expectedResult {
		t.Errorf("Case 2. Parsing failed, expected: %d got: %d", expectedResult, result)
	}

	// Case 3
	testValue = "1633102"
	expectedResult = 163310200000000

	dv, err = decimal.NewFromString(testValue)
	if err != nil {
		t.Error(err)
		return
	}

	result = DecimalToStatoshi(dv)

	if result != expectedResult {
		t.Errorf("Case 3. Parsing failed, expected: %d got: %d", expectedResult, result)
	}

	// Case 4
	testValue = "-48.000012"
	expectedResult = -4800001200

	dv, err = decimal.NewFromString(testValue)
	if err != nil {
		t.Error(err)
		return
	}

	result = DecimalToStatoshi(dv)

	if result != expectedResult {
		t.Errorf("Case 4. Parsing failed, expected: %d got: %d", expectedResult, result)
	}
}

func TestDecimalToSize(t *testing.T) {
	// Case 1
	var testValue string = "48.000012"
	var expectedResult int64 = 48000012

	dv, err := decimal.NewFromString(testValue)
	if err != nil {
		t.Error(err)
		return
	}

	var result int64 = DecimalToSize(dv)

	if result != expectedResult {
		t.Errorf("Case 1. Parsing failed, expected: %d got: %d", expectedResult, result)
	}
}

func TestDecimalToPercetage(t *testing.T) {
	// Case 1
	var testValue string = "48.000012"
	var expectedResult int32 = 480000

	dv, err := decimal.NewFromString(testValue)
	if err != nil {
		t.Error(err)
		return
	}

	var result int32 = DecimalToPercentage(dv)

	if result != expectedResult {
		t.Errorf("Case 1. Parsing failed, expected: %d got: %d", expectedResult, result)
	}

	// Case 2
	testValue = "0.1"
	expectedResult = 1000

	dv, err = decimal.NewFromString(testValue)
	if err != nil {
		t.Error(err)
		return
	}

	result = DecimalToPercentage(dv)

	if result != expectedResult {
		t.Errorf("Case 2. Parsing failed, expected: %d got: %d", expectedResult, result)
	}
}

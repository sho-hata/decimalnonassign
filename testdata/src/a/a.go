package a

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func calledDecimalMethodSimple() {
	d := decimal.Zero
	d.Add(d) // want "result is not assigned"
	d.Sub(d) // want "result is not assigned"
	d.Div(d) // want "result is not assigned"
	result := d.Add(decimal.Zero)
	fmt.Println(d.Add(decimal.Zero))
	fmt.Print(result)
}

func calledDecimalMethodInFor() {
	result := decimal.Zero

	for {
		result = result.Add(decimal.Zero)
		result.Add(decimal.Zero) // want "result is not assigned"
	}
}

func calledDecimalMethodInRange() {
	ds := []decimal.Decimal{decimal.Zero, decimal.Zero, decimal.Zero}

	result := decimal.Zero
	for _, d := range ds {
		result = result.Add(d)
		result.Add(d) // want "result is not assigned"
	}
}

func calledDecimalMethodInIf() {
	ds := []decimal.Decimal{decimal.Zero, decimal.Zero, decimal.Zero}

	result := decimal.Zero

	if true {
		for _, d := range ds {
			result = result.Add(d)
			result.Add(d) // want "result is not assigned"
		}
	} else {
		for _, d := range ds {
			result = result.Sub(d)
			result.Sub(d) // want "result is not assigned"
		}
	}
}

func calledDecimalMethodInSwitch() {
	var n int
	result := decimal.Zero
	switch n {
	case 1:
		result = result.Add(decimal.Zero)
		result.Add(decimal.Zero) // want "result is not assigned"
	case 2:
		result = result.Sub(decimal.Zero)
		result.Sub(decimal.Zero) // want "result is not assigned"
	case 3:
		result = result.Div(decimal.Zero)
		result.Div(decimal.Zero) // want "result is not assigned"
	default:
		result = result.Mul(decimal.Zero)
		result.Mul(decimal.Zero) // want "result is not assigned"
	}
}

func calledDecimalMethodInDefer() {
	result := decimal.Zero

	defer func() {
		result = result.Add(decimal.Zero)
		result.Add(decimal.Zero) // want "result is not assigned"
	}()
}

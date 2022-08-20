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

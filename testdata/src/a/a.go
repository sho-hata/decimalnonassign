package a

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func calledDecimalAdd() {
	d := decimal.Zero
	d.Add(d) // want "result is not assigned"
	d.Sub(d) // want "result is not assigned"
	d.Div(d) // want "result is not assigned"
	result := d.Add(decimal.Zero)
	fmt.Print(result)
}

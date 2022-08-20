package a

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

func calledDecimalMethodSimple() {
	d := decimal.Zero
	d.Add(d) // want "The result of 'Add' is not assigned"
	d.Sub(d) // want "The result of 'Sub' is not assigned"
	d.Div(d) // want "The result of 'Div' is not assigned"
	result := d.Add(decimal.Zero)
	fmt.Println(d.Add(decimal.Zero))
	fmt.Print(result)
}

func calledDecimalMethodInFor() {
	result := decimal.Zero

	for {
		result = result.Add(decimal.Zero)
		result.Add(decimal.Zero) // want "The result of 'Add' is not assigned"
	}
}

func calledDecimalMethodInRange() {
	ds := []decimal.Decimal{decimal.Zero, decimal.Zero, decimal.Zero}

	result := decimal.Zero
	for _, d := range ds {
		result = result.Add(d)
		result.Add(d) // want "The result of 'Add' is not assigned"
	}
}

func calledDecimalMethodInIf() {
	ds := []decimal.Decimal{decimal.Zero, decimal.Zero, decimal.Zero}

	result := decimal.Zero

	if true {
		for _, d := range ds {
			result = result.Add(d)
			result.Add(d) // want "The result of 'Add' is not assigned"
		}
	} else {
		for _, d := range ds {
			result = result.Sub(d)
			result.Sub(d) // want "The result of 'Sub' is not assigned"
		}
	}
}

func calledDecimalMethodInSwitch() {
	var n int
	result := decimal.Zero
	switch n {
	case 1:
		result = result.Add(decimal.Zero)
		result.Add(decimal.Zero) // want "The result of 'Add' is not assigned"
	case 2:
		result = result.Sub(decimal.Zero)
		result.Sub(decimal.Zero) // want "The result of 'Sub' is not assigned"
	case 3:
		result = result.Div(decimal.Zero)
		result.Div(decimal.Zero) // want "The result of 'Div' is not assigned"
	default:
		result = result.Mul(decimal.Zero)
		result.Mul(decimal.Zero) // want "The result of 'Mul' is not assigned"
	}
}

func calledDecimalMethodInDefer() {
	result := decimal.Zero

	defer func() {
		result = result.Add(decimal.Zero)
		result.Add(decimal.Zero) // want "The result of 'Add' is not assigned"
	}()
}

func calledDecimalMethodInGo() {
	result := decimal.Zero

	go func() {
		result = result.Add(decimal.Zero)
		result.Add(decimal.Zero) // want "The result of 'Add' is not assigned"
	}()
}

func calledDecimalMethodInSelect() {
	c1 := make(chan decimal.Decimal)
	c2 := make(chan decimal.Decimal)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- decimal.NewFromFloat(1)
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- decimal.NewFromFloat(2)
	}()

	result := decimal.Zero

	for i := 0; i < 2; i++ {
		select {
		case v := <-c1:
			result = result.Add(v)
			result.Add(v) // want "The result of 'Add' is not assigned"
		case v := <-c2:
			result = result.Add(v)
			result.Add(v) // want "The result of 'Add' is not assigned"
		}
	}

	fmt.Println(result.String()) // 3
}

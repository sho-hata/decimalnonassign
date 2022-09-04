# decimalnonassign
Go linter that checks if the result of a decimal operation is assigned

supported package: https://pkg.go.dev/github.com/shopspring/decimal

## Motivation
The arithmetic methods in ` github.com/shopspring/decimal`  are immutable. The result must therefore be assigned to a variable.
If no assignment is made, unexpected situations are more likely to occur.

```go
func sample() decimal.Decimal {
  result := decimal.Zero
  result = result.Add(decimal.NewFromFloat(2)) // ok
  result.Add(decimal.NewFromFloat(1)) // oops, Value not updated...💀

  return result
}
```
## Usage
```go
package main

import "github.com/shopspring/decimal"

func calledDecimalMethodInRange() {
	ds := []decimal.Decimal{decimal.NewFromFloat(2), decimal.NewFromFloat(2), decimal.NewFromFloat(2)}

	result := decimal.Zero
	for _, d := range ds {
		result = result.Add(d)
		result.Add(d)
	}
}
```
## Analysis
```bash
$ decimalnoassign ./...
./a.go:11:10: The result of 'Add' is not assigned
```

## Installation
```
go install github.com/sho-hata/decimalnonassign/cmd/decimalnonassign@latest
```

## Contribution
1. Fork (https://github.com/sho-hata/decimalnonassign/fork)
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the go `test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

## License
[MIT](https://github.com/sho-hata/decimalnonassign/blob/main/LICENSE)

## Author
[sho-hata](https://github.com/sho-hata)

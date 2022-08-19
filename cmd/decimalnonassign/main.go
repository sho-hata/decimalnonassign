package main

import (
	"github.com/sho-hata/decimalnonassign"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(decimalnonassign.Analyzer) }

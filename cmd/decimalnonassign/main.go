package main

import (
	"github.com/sho-hata/decimalnonassign"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(decimalnonassign.Analyzer) }

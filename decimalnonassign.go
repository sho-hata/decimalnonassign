package decimalnonassign

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "decimalnonassign is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "decimalnonassign",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	// List of methods that return decimal.Decimal
	// ref: https://pkg.go.dev/github.com/shopspring/decimal#pkg-types
	var decimalMethodNames = map[string]bool{
		"Abs":        true,
		"Add":        true,
		"Atan":       true,
		"Ceil":       true,
		"Copy":       true,
		"Cos":        true,
		"Div":        true,
		"DivRound":   true,
		"Floor":      true,
		"Mod":        true,
		"Mul":        true,
		"Neg":        true,
		"Pow":        true,
		"Round":      true,
		"RoundBank":  true,
		"RoundCash":  true,
		"RoundCeil":  true,
		"RoundDown":  true,
		"RoundFloor": true,
		"RoundUp":    true,
		"Shift":      true,
		"Sin":        true,
		"Sub":        true,
		"Tan":        true,
		"Truncate":   true,
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if n, ok := n.(*ast.FuncDecl); ok {
			for _, s := range n.Body.List {
				if s, ok := s.(*ast.ExprStmt); ok {
					if ex, ok := s.X.(*ast.CallExpr); ok {
						if s, ok := ex.Fun.(*ast.SelectorExpr); ok {
							if i, ok := s.X.(*ast.Ident); ok {
								if strings.HasSuffix(pass.TypesInfo.TypeOf(i).String(), "github.com/shopspring/decimal.Decimal") {
									if _, ok := decimalMethodNames[s.Sel.Name]; ok {
										pass.Reportf(i.Pos(), "result is not assigned")
									}
								}
							}
						}
					}
				}
			}
		}
	})

	return nil, nil
}

package decimalnonassign

import (
	"errors"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	doc        = "decimalnonassign is Go linter that checks if the result of a decimal operation is assigned"
	targetType = "github.com/shopspring/decimal.Decimal"
)

var Analyzer = &analysis.Analyzer{
	Name: "decimalnonassign",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
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

func run(pass *analysis.Pass) (interface{}, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errors.New("file inspection failed")
	}

	nodeFilter := []ast.Node{(*ast.FuncDecl)(nil)}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if n, ok := n.(*ast.FuncDecl); ok {
			report(pass, n.Body.List)
		}
	})

	return nil, nil
}

func report(pass *analysis.Pass, ss []ast.Stmt) {
	for _, s := range ss {
		switch s := s.(type) {
		case *ast.ExprStmt:
			if ex, ok := s.X.(*ast.CallExpr); ok {
				if s, ok := ex.Fun.(*ast.SelectorExpr); ok {
					if ok, pos := findDecimalMethodPos(pass, s); ok {
						pass.Reportf(pos, "result is not assigned")
					}
				}
			}
		case *ast.ForStmt:
			report(pass, s.Body.List)
		case *ast.RangeStmt:
			report(pass, s.Body.List)
		case *ast.IfStmt:
			report(pass, s.Body.List)

			switch s := s.Else.(type) {
			case *ast.BlockStmt:
				report(pass, s.List)
			case *ast.IfStmt:
				report(pass, s.Body.List)
			}
		case *ast.SwitchStmt:
			report(pass, s.Body.List)
		case *ast.CaseClause:
			report(pass, s.Body)
		case *ast.DeferStmt:
			if f, ok := s.Call.Fun.(*ast.FuncLit); ok {
				report(pass, f.Body.List)
			}
		case *ast.GoStmt:
			if f, ok := s.Call.Fun.(*ast.FuncLit); ok {
				report(pass, f.Body.List)
			}
		case *ast.SelectStmt:
			report(pass, s.Body.List)
		case *ast.CommClause:
			report(pass, s.Body)
		}
	}
}

func findDecimalMethodPos(pass *analysis.Pass, e *ast.SelectorExpr) (bool, token.Pos) {
	if i, ok := e.X.(*ast.Ident); ok {
		if strings.HasSuffix(pass.TypesInfo.TypeOf(i).String(), targetType) {
			if _, ok := decimalMethodNames[e.Sel.Name]; ok {
				return true, i.Pos()
			}
		}
	}

	return false, token.NoPos
}

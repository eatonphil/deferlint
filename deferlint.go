package deferlint

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "deferlint",
	Doc:      "reports defer in loops",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder(nil, func(n ast.Node) {
		var list []ast.Stmt

		if fs, ok := n.(*ast.ForStmt); ok {
			list = fs.Body.List
		} else if rs, ok := n.(*ast.RangeStmt); ok {
			list = rs.Body.List
		} else {
			return
		}

		for _, stmt := range list {
			ds, ok := stmt.(*ast.DeferStmt)
			if ok {
				pass.Reportf(ds.Pos(), "defer in loop found %q",
					prettyPrint(pass.Fset, ds))
				return
			}
		}
	})

	return nil, nil
}

func prettyPrint(fset *token.FileSet, node interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, node); err != nil {
		panic(err)
	}
	return buf.String()
}

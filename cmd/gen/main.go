package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	prog, err := parser.ParseFile(fset, "hello.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Inspect(prog, func(n ast.Node) bool {
		// Find Function Call Statements
		funcCall, ok := n.(*ast.CallExpr)
		if ok {
			fmt.Println(funcCall.Fun)
		}
		return true
	})
	ast.Print(fset, prog)
}

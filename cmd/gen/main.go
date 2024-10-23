package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	//Create a FileSet to work with
	fset := token.NewFileSet()
	//Parse the file and create an AST
	file, err := parser.ParseFile(fset, "hello.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Inspect(file, func(n ast.Node) bool {
		// Find Function Call Statements
		funcCall, ok := n.(*ast.CallExpr)
		if ok {
			fmt.Println(funcCall.Fun)
		}
		return true
	})
}

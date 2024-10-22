package main

import (
	"log"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// This example builds SSA code for a set of packages plus all their dependencies,
// using the [golang.org/x/tools/go/packages] API.
// This is what you'd typically use for a whole-program analysis.
func main() {
	// Load, parse, and type-check the whole program.
	cfg := packages.Config{Mode: packages.LoadAllSyntax}
	initial, err := packages.Load(&cfg, "fmt", "net/http")
	if err != nil {
		log.Fatal(err)
	}
	// Create SSA packages for well-typed packages and their dependencies.
	prog, pkgs := ssautil.AllPackages(initial, ssa.PrintPackages|ssa.InstantiateGenerics)
	_ = pkgs
	// Build SSA code for the whole program.
	prog.Build()
}

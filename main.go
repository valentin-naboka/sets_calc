package main

import (
	"fmt"
	"sets_calc/calcengine"
	"sets_calc/lexer"
	"sets_calc/parser"
	"strings"
)

func main() {

	r := strings.NewReader("[ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]")
	//r := strings.NewReader("[ LE 2 a.txt [ GR 1 b.txt c.txt ] ]")
	//r := strings.NewReader("[ GR 1 b.txt c.txt ]")
	//r := strings.NewReader("[ LE 2 a.txt ]")
	scanner := lexer.NewScanner(r)
	scanner.Scan()

	parser := parser.Parser{scanner}
	expr := parser.BuildTree()
	expr.Print()

	var engine calcengine.Calculator

	result := engine.Execute(expr)

	fmt.Printf("\n%v", result)

	// scanner.PrintTokens()
	// scanner.PrintValues()

	// file := ast.File{"a.txt"}
	// var op ast.Operator
	// op.Sets = make([]ast.Expr, 0, 3)

	// op2 := ast.Operator{OpType: ast.GreaterThan, N: 1}

	// op.Sets = append(op.Sets, &file, &ast.File{"b.txt"}, &op2)

	// op.N = ast.Int(2)
	// op.Print()
	// println()
}

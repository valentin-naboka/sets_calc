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

	parser := parser.Parser{Scanner: scanner}
	expr := parser.BuildTree()
	expr.Print()

	var engine calcengine.Calculator

	result := engine.Execute(expr)

	fmt.Printf("\n%v", result)
}

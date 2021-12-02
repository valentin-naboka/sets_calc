package main

import (
	"log"
	"os"
	"sets_calc/calcengine"
	"sets_calc/lexer"
	"sets_calc/parser"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		println("Expression is missing")
		os.Exit(1)
	}
	input := os.Args[1]

	r := strings.NewReader(input)

	scanner := lexer.NewScanner(r)
	parser := parser.Parser{Scanner: scanner}
	expr, err := parser.BuildExpression()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	engine := calcengine.NewCalculator(&calcengine.FileReader{})
	result, err := engine.Execute(expr)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	for _, i := range result {
		println(i)
	}
}

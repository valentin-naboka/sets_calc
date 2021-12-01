package main

import (
	"fmt"
	"log"
	"os"
	"sets_calc/calcengine"
	"sets_calc/lexer"
	"sets_calc/parser"
	"strings"
)

func main() {

	{
		r := strings.NewReader("[ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]")
		scanner := lexer.NewScanner(r)

		parser := parser.Parser{Scanner: scanner}
		expr, err := parser.BuildExpression()
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}

		expr.Print()

		var engine calcengine.Calculator

		result, _ := engine.Execute(expr)

		fmt.Printf("\n%v\n", result)
	}

	{
		r := strings.NewReader("[ LE 2 a.txt [ GR 1 b.txt c.txt ] ]")
		scanner := lexer.NewScanner(r)

		parser := parser.Parser{Scanner: scanner}
		expr, err := parser.BuildExpression()
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}

		expr.Print()

		var engine calcengine.Calculator

		result, _ := engine.Execute(expr)

		fmt.Printf("\n%v\n", result)
	}

	{
		r := strings.NewReader("[ GR 1 b.txt c.txt ]")
		scanner := lexer.NewScanner(r)

		parser := parser.Parser{Scanner: scanner}
		expr, err := parser.BuildExpression()
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		expr.Print()

		var engine calcengine.Calculator

		result, _ := engine.Execute(expr)

		fmt.Printf("\n%v\n", result)
	}

	{
		r := strings.NewReader("[ LE 2 a.txt ]")
		scanner := lexer.NewScanner(r)

		parser := parser.Parser{Scanner: scanner}
		expr, err := parser.BuildExpression()
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		expr.Print()

		var engine calcengine.Calculator

		result, _ := engine.Execute(expr)

		fmt.Printf("\n%v\n", result)
	}
}

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

	// TODO: Due to the remark to the provided solution,
	// the next lines have been added to match the exact specification
	// that was provided in the email.
	//
	// "it does not work, unfortunately. example from specification returns error:
	// $ ./scalc [ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]
	// 2021/12/03 11:45:01 unexpected token : whitespace is expected"

	var strBuilder strings.Builder
	if len(os.Args) > 2 {
		for _, arg := range os.Args[1 : len(os.Args)-1] {
			_, err := strBuilder.WriteString(arg)
			if err != nil {
				log.Print(err)
				os.Exit(1)
			}
			_, err = strBuilder.WriteString(" ")
			if err != nil {
				log.Print(err)
				os.Exit(1)
			}
		}
		_, err := strBuilder.WriteString(os.Args[len(os.Args)-1])
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
	} else {
		_, err := strBuilder.WriteString(os.Args[1])
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
	}

	input := strBuilder.String()

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

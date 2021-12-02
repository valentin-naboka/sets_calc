package parser_test

import (
	"reflect"
	"runtime"
	"sets_calc/ast"
	"sets_calc/lexer"
	"sets_calc/parser"
	"strings"
	"testing"
)

func printCaller(t *testing.T, depth int) {
	function, file, line, _ := runtime.Caller(depth)
	trimName := func(n string) string {
		i := strings.LastIndex(n, "/")
		if i == -1 {
			return n
		}
		return n[i+1:]
	}
	t.Logf("%s: line %d, function: %s\n", trimName(file), line, trimName(runtime.FuncForPC(function).Name()))
}

func checkValue(t *testing.T, input string, expected *ast.Expression) {
	p := parser.NewParser(lexer.NewScanner(strings.NewReader(input)))
	expr, err := p.BuildExpression()
	if err != nil {
		printCaller(t, 2)
		t.Fatal("unable to parse en expression: ", err)
	}

	if !reflect.DeepEqual(expected, expr) {
		printCaller(t, 2)
		t.Errorf("\nexpected: %v\nactual:   %v", expected, expr)
	}
}

//TODO: check on a specific error
func checkError(t *testing.T, input string) {
	p := parser.NewParser(lexer.NewScanner(strings.NewReader(input)))
	expr, err := p.BuildExpression()
	if err == nil {
		printCaller(t, 2)
		t.Errorf("unexpected expression: %v", expr)
	}
}

func TestParser(t *testing.T) {
	checkValue(t, "[ EQ 1 c.txt ]", &ast.Expression{OpType: ast.Equal, N: 1, Sets: []ast.FileOrExpression{ast.File{Name: "c.txt"}}})
	checkValue(t, "[ EQ 1 123c.txt ]", &ast.Expression{OpType: ast.Equal, N: 1, Sets: []ast.FileOrExpression{ast.File{Name: "123c.txt"}}})

	checkValue(t,
		"[ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]",
		&ast.Expression{OpType: ast.GreaterThan, N: 1, Sets: []ast.FileOrExpression{
			ast.File{Name: "c.txt"},
			&ast.Expression{OpType: ast.Equal, N: 3, Sets: []ast.FileOrExpression{
				ast.File{Name: "a.txt"},
				ast.File{Name: "a.txt"},
				ast.File{Name: "b.txt"},
			}}}})

	checkValue(t, "[ LE 2 a.txt [ GR 1 b.txt c.txt ] ]", &ast.Expression{OpType: ast.LessThan, N: 2, Sets: []ast.FileOrExpression{
		ast.File{Name: "a.txt"},
		&ast.Expression{OpType: ast.GreaterThan, N: 1, Sets: []ast.FileOrExpression{
			ast.File{Name: "b.txt"},
			ast.File{Name: "c.txt"},
		}}}})
}

func TestParserErrors(t *testing.T) {
	checkError(t, "[ EQ 1 ]")
	checkError(t, "")
	checkError(t, "[")
	checkError(t, "[ EQ 0 c.txt ]")
	checkError(t, "[ EQ -1 c.txt ]")
	checkError(t, "[ GR 1 c.txt [EQ 3 b.txt ] ]")
	checkError(t, "[ GR 1 c.txt [ EE 3 b.txt ] ]")
	checkError(t, "[ GR 1 c.txt [ EQ 3 b.txt] ]")
	checkError(t, "[ GR 1 c.txt [ EE 3 b.txt ] ]")
	checkError(t, "[ GR 1 c.txt [ EQ 3 b.txt ] ]]")
}

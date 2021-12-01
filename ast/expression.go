package ast

import "fmt"

type Operator uint8

const (
	Equal Operator = iota
	LessThan
	GreaterThan
)

type FileOrExpression interface{}

type Expression struct {
	OpType Operator
	N      Int
	Sets   []FileOrExpression
}

func (e *Expression) Print() {
	print("[ ")
	switch e.OpType {
	case Equal:
		fmt.Printf("EQ")
	case LessThan:
		fmt.Printf("LE")
	case GreaterThan:
		fmt.Printf("GE")
	default:
		panic("Unexpected op type")
	}

	fmt.Printf(" %v ", e.N)
	for _, v := range e.Sets {
		file, ok := v.(File)
		if ok {
			file.Print()
		} else {
			expr, ok := v.(*Expression)
			if ok {
				expr.Print()
			}
		}
		print(" ")
	}
	print("]")
}

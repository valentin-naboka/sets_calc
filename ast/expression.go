package ast

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

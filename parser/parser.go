package parser

import (
	"sets_calc/ast"
	"sets_calc/lexer"
	"strconv"

	"github.com/pkg/errors"
)

const whitespaceErrStr string = "unexpected token %s: whitespace is expected"

type Parser struct {
	Scanner *lexer.Scanner
}

func NewParser(s *lexer.Scanner) *Parser {
	return &Parser{Scanner: s}
}

// NOTE: The big depth of recursion may lead to stack overflow.
// In such a case, better to replace it with the stack data structure.
func (s *Parser) BuildExpression() (*ast.Expression, error) {
	expr, err := s.parseExpr(s.Scanner.NextToken())
	if err != nil {
		return expr, err
	}

	token := s.Scanner.NextToken()
	if token.Type != lexer.EOF {
		return nil, errors.Errorf("unexpected token: '%s' at the end of the expression", token.Value)
	}
	return expr, nil
}

func (s *Parser) parseExpr(token *lexer.Token) (*ast.Expression, error) {
	if token.Type != lexer.LeftBracket {
		return nil, errors.Errorf("unexpected token %s: '[' is expected", token.Value)
	}

	token = s.Scanner.NextToken()
	if token.Type != lexer.Whitespace {
		return nil, errors.Errorf(whitespaceErrStr, token.Value)
	}

	token = s.Scanner.NextToken()
	if token.Type != lexer.Operator {
		return nil, errors.Errorf("unexpected token %s: operator is expected", token.Value)
	}
	var expr ast.Expression

	opType, err := s.parseOperator(token.Value)
	if err != nil {
		return nil, err
	}

	expr.OpType = opType

	token = s.Scanner.NextToken()
	if token.Type != lexer.Whitespace {
		return nil, errors.Errorf(whitespaceErrStr, token.Value)
	}

	token = s.Scanner.NextToken()
	if token.Type != lexer.Integer {
		return nil, errors.Errorf("unexpected token %s: positive integer number is expected", token.Value)
	}

	integer, err := s.parseInteger(token.Value)
	if err != nil {
		return nil, err
	}
	expr.N = integer

	token = s.Scanner.NextToken()
	if token.Type != lexer.Whitespace {
		return nil, errors.Errorf(whitespaceErrStr, token.Value)
	}

	expr.Sets, err = s.parseSets()
	if err != nil {
		return nil, err
	}

	return &expr, nil
}

func (s *Parser) parseInteger(i string) (ast.Int, error) {
	integer, err := strconv.Atoi(i)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to convert N: %s", i)
	}

	if integer < 1 {
		return 0, errors.Errorf("N should be positive number: %s", i)
	}

	return ast.Int(integer), nil
}

func (s *Parser) parseOperator(op string) (ast.Operator, error) {
	switch op {
	case "EQ":
		return ast.Equal, nil
	case "GR":
		return ast.GreaterThan, nil
	case "LE":
		return ast.LessThan, nil
	default:
		return ast.Equal, errors.Errorf("unknown operator type %s", op)
	}
}

func (s *Parser) parseSets() ([]ast.FileOrExpression, error) {
	const unexpectedTokenErrStr = "unexpected token '%s': filename or expression statement is expected"
	token := s.Scanner.NextToken()

	if token.Type == lexer.RightBracket || token.Type == lexer.EOF {
		return nil, errors.Errorf(unexpectedTokenErrStr, token.Value)
	}

	var sets []ast.FileOrExpression
	for ; token.Type != lexer.RightBracket; token = s.Scanner.NextToken() {
		if token.Type == lexer.File {
			sets = append(sets, ast.File{Name: token.Value})
		} else if token.Type == lexer.LeftBracket {
			e, err := s.parseExpr(token)
			if err != nil {
				return nil, err
			}

			if e != nil {
				sets = append(sets, e)
			}
		} else {
			return nil, errors.Errorf(unexpectedTokenErrStr, token.Value)
		}

		if s.Scanner.NextToken().Type != lexer.Whitespace {
			return nil, errors.Errorf(whitespaceErrStr, token.Value)
		}
	}
	return sets, nil
}

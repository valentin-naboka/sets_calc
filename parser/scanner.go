package parser

import (
	"sets_calc/ast"
	"sets_calc/lexer"
	"strconv"

	"github.com/pkg/errors"
)

type Parser struct {
	Scanner *lexer.Scanner
}

func NewParser(s *lexer.Scanner) *Parser {
	return &Parser{Scanner: s}
}

func (s *Parser) BuildTree() *ast.Expression {
	return s.parseExpr()
}

func (s *Parser) parseExpr() *ast.Expression {

	token := s.Scanner.NextToken()
	if token.Type != lexer.LeftBracket {
		return nil
	}

	token = s.Scanner.NextToken()
	if token.Type != lexer.Whitespace {
		return nil
	}

	token = s.Scanner.NextToken()
	if token.Type != lexer.Operator {
		return nil
	}
	var expr ast.Expression

	opType, err := s.parseOperator(token.Value)
	if err != nil {
		return nil
	}

	expr.OpType = opType

	token = s.Scanner.NextToken()
	if token.Type != lexer.Whitespace {
		return nil
	}

	token = s.Scanner.NextToken()
	if token.Type != lexer.Integer {
		return nil
	}

	integer, err := s.parseInteger(token.Value)
	if err != nil {
		return nil
	}
	expr.N = integer

	token = s.Scanner.NextToken()
	if token.Type != lexer.Whitespace {
		return nil
	}

	for token := s.Scanner.PeekNextToken(); token.Type != lexer.RightBracket; token = s.Scanner.PeekNextToken() {
		if token.Type == lexer.File {
			expr.Sets = append(expr.Sets, ast.File{Name: token.Value})
			s.Scanner.NextToken()
		} else if token.Type == lexer.LeftBracket {
			e := s.parseExpr()
			if e != nil {
				expr.Sets = append(expr.Sets, e)
				s.Scanner.NextToken()
			}
		} else {
			return nil
		}

		token = s.Scanner.NextToken()
		//TODO: костыль
		if token.Type != lexer.Whitespace && token.Type != lexer.RightBracket {
			return nil
		}
	}
	return &expr
}

func (s *Parser) parseInteger(i string) (ast.Int, error) {
	//TODO: positive number
	integer, err := strconv.Atoi(i)
	return ast.Int(integer), err
}

//TODO: use just string
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

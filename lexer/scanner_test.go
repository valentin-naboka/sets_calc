package lexer_test

import (
	"sets_calc/lexer"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	{
		s := lexer.NewScanner(strings.NewReader("[ EQ 1 c.txt ]"))
		i := 0
		var expected = []lexer.TokenType{
			lexer.LeftBracket,
			lexer.Whitespace,
			lexer.Operator,
			lexer.Whitespace,
			lexer.Integer,
			lexer.Whitespace,
			lexer.File,
			lexer.Whitespace,
			lexer.RightBracket,
		}

		for token := s.NextToken(); token.Type != lexer.EOF; token = s.NextToken() {
			if token.Type != expected[i] {
				t.Errorf("unexpected token: '%s'", token.Value)
			}
			i++
		}
	}

	{
		s := lexer.NewScanner(strings.NewReader("[ EQ 1 123c.txt ]"))
		i := 0
		var expected = []lexer.TokenType{
			lexer.LeftBracket,
			lexer.Whitespace,
			lexer.Operator,
			lexer.Whitespace,
			lexer.Integer,
			lexer.Whitespace,
			lexer.File,
			lexer.Whitespace,
			lexer.RightBracket,
		}

		for token := s.NextToken(); token.Type != lexer.EOF; token = s.NextToken() {
			if token.Type != expected[i] {
				t.Errorf("unexpected token: '%s'", token.Value)
			}
			i++
		}
	}

	{
		s := lexer.NewScanner(strings.NewReader("[   EQ 1    c.txt    ]"))
		i := 0
		var expected = []lexer.TokenType{
			lexer.LeftBracket,
			lexer.Whitespace,
			lexer.Operator,
			lexer.Whitespace,
			lexer.Integer,
			lexer.Whitespace,
			lexer.File,
			lexer.Whitespace,
			lexer.RightBracket,
		}

		for token := s.NextToken(); token.Type != lexer.EOF; token = s.NextToken() {
			if token.Type != expected[i] {
				t.Errorf("unexpected token: '%s'", token.Value)
			}
			i++
		}
	}

	{
		s := lexer.NewScanner(strings.NewReader("[ EQ 1 ]"))
		i := 0
		var expected = []lexer.TokenType{
			lexer.LeftBracket,
			lexer.Whitespace,
			lexer.Operator,
			lexer.Whitespace,
			lexer.Integer,
			lexer.Whitespace,
			lexer.RightBracket,
		}

		for token := s.NextToken(); token.Type != lexer.EOF; token = s.NextToken() {
			if token.Type != expected[i] {
				t.Errorf("unexpected token: '%s'", token.Value)
			}
			i++
		}
	}

	{
		s := lexer.NewScanner(strings.NewReader(""))
		token := s.NextToken()
		if token.Type != lexer.EOF {
			t.Errorf("unexpected token: '%s'", token.Value)
		}
	}

	{
		s := lexer.NewScanner(strings.NewReader("["))
		token := s.NextToken()
		if token.Type != lexer.LeftBracket {
			t.Errorf("unexpected token: '%s'", token.Value)
		}

		token = s.NextToken()
		if token.Type != lexer.EOF {
			t.Errorf("unexpected token: '%s'", token.Value)
		}
	}

	{
		s := lexer.NewScanner(strings.NewReader("[ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]"))
		i := 0
		var expected = []lexer.TokenType{
			lexer.LeftBracket,
			lexer.Whitespace,
			lexer.Operator,
			lexer.Whitespace,
			lexer.Integer,
			lexer.Whitespace,
			lexer.File,
			lexer.Whitespace,
			lexer.LeftBracket,
			lexer.Whitespace,
			lexer.Operator,
			lexer.Whitespace,
			lexer.Integer,
			lexer.Whitespace,
			lexer.File,
			lexer.Whitespace,
			lexer.File,
			lexer.Whitespace,
			lexer.File,
			lexer.Whitespace,
			lexer.RightBracket,
			lexer.Whitespace,
			lexer.RightBracket,
		}

		for token := s.NextToken(); token.Type != lexer.EOF; token = s.NextToken() {
			if token.Type != expected[i] {
				t.Errorf("unexpected token: '%s'", token.Value)
			}
			i++
		}
	}

	{
		s := lexer.NewScanner(strings.NewReader("[ LE 2 a.txt [ GR 1 b.txt c.txt ] ]"))
		i := 0
		var expected = []lexer.TokenType{
			lexer.LeftBracket,
			lexer.Whitespace,
			lexer.Operator,
			lexer.Whitespace,
			lexer.Integer,
			lexer.Whitespace,
			lexer.File,
			lexer.Whitespace,
			lexer.LeftBracket,
			lexer.Whitespace,
			lexer.Operator,
			lexer.Whitespace,
			lexer.Integer,
			lexer.Whitespace,
			lexer.File,
			lexer.Whitespace,
			lexer.File,
			lexer.Whitespace,
			lexer.RightBracket,
			lexer.Whitespace,
			lexer.RightBracket,
		}

		for token := s.NextToken(); token.Type != lexer.EOF; token = s.NextToken() {
			if token.Type != expected[i] {
				t.Errorf("unexpected token: '%s'", token.Value)
			}
			i++
		}
	}
}

package lexer

import (
	"bufio"
	"io"
	"log"
	"strings"
	"unicode"
)

const (
	whitespace rune = ' '
	eof        rune = 0
)

func isWhitespaceOrEOF(ch rune) bool {
	return ch == whitespace || ch == eof
}

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) readNext() rune {
	ch, _, err := s.r.ReadRune()
	if err == io.EOF {
		return eof
	}

	if err != nil {
		log.Printf("unable to read next character: %s", err)
		return eof
	}

	return ch
}

func (s *Scanner) NextToken() *Token {
	ch := s.readNext()
	switch ch {
	case eof:
		return &Token{EOF, ""}
	case '[':
		return &Token{LeftBracket, string(ch)}
	case ']':
		return &Token{RightBracket, string(ch)}
	case whitespace:
		return s.readWhitespaces(ch)
	}

	token, stringBuilder := s.readDigit(ch)
	if token != nil {
		return token
	}

	if stringBuilder == nil {
		token, stringBuilder = s.readOperator(ch)
		if token != nil {
			return token
		}
	}

	return s.readFilename(stringBuilder)
}

func (s *Scanner) readWhitespaces(ch rune) *Token {
	for ; ch == whitespace; ch = s.readNext() {
	}
	s.r.UnreadRune()
	return &Token{Whitespace, string(whitespace)}
}

func (s *Scanner) readDigit(ch rune) (*Token, *strings.Builder) {

	if !unicode.IsDigit(ch) {
		return nil, nil
	}

	var number strings.Builder
	for ; unicode.IsDigit(ch); ch = s.readNext() {
		number.WriteRune(ch)
	}

	s.r.UnreadRune()
	if ch != whitespace {
		return nil, &number
	}

	return &Token{Integer, number.String()}, nil
}

func (s *Scanner) readOperator(ch rune) (*Token, *strings.Builder) {

	var op strings.Builder
	op.WriteRune(ch)
	op.WriteRune(s.readNext())

	//NOTE: peek the rune
	ch = s.readNext()
	s.r.UnreadRune()

	opStr := op.String()
	if isWhitespaceOrEOF(ch) && (opStr == "EQ" || opStr == "GR" || opStr == "LE") {
		return &Token{Operator, opStr}, nil
	}

	return nil, &op
}

func (s *Scanner) readFilename(b *strings.Builder) *Token {
	if b == nil {
		panic("b *strings.Builder varibale is nil")
	}

	for ch := s.readNext(); !isWhitespaceOrEOF(ch); ch = s.readNext() {
		b.WriteRune(ch)
	}

	s.r.UnreadRune()
	return &Token{File, b.String()}
}

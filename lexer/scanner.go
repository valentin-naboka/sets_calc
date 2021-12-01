//[ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ]
package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Scanner struct {
	r *bufio.Reader

	//TODO: remove, use chanel
	tokens   []Token
	tokenIdx int
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r), tokenIdx: -1}
}

func (s *Scanner) readNext() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		//TODO: resolve
		return rune(0)
	}
	return ch
}

func (s *Scanner) NextToken() (t Token) {
	s.tokenIdx++
	return s.tokens[s.tokenIdx]
}

func (s *Scanner) PeekNextToken() (t Token) {
	return s.tokens[s.tokenIdx+1]
}

func (s *Scanner) Scan() /*(t Token)*/ {

	//TODO: resolve
	for ch := s.readNext(); ch != 0; ch = s.readNext() {
		switch ch {
		case 0:
			{
				s.tokens = append(s.tokens, Token{EOF, ""})
				continue
			}
		case '[':
			{
				s.tokens = append(s.tokens, Token{LeftBracket, string(ch)})
				continue
			}
		case ']':
			{
				s.tokens = append(s.tokens, Token{RightBracket, string(ch)})
				continue
			}
		case ' ':
			{
				s.readWhitespaces(ch)
				continue
			}
		}

		ok, stringBuilder := s.readDigit(ch)
		if ok {
			continue
		}

		if stringBuilder == nil {
			ok, stringBuilder = s.readOperator(ch)
			if ok {
				continue
			}
		}

		s.readFilename(stringBuilder)

		//TODO: what do to with errors?
		//s.tokens = append(s.tokens, Token{Error, fmt.Sprintf("unknown symbol %c", ch)})
	}
}

func (s *Scanner) readWhitespaces(ch rune) {
	for ; ; ch = s.readNext() {
		//TODO: rewrite to loop body
		// or end of file
		if ch != ' ' {
			break
		}
	}
	s.r.UnreadRune()
	s.tokens = append(s.tokens, Token{Whitespace, string(' ')})
}

func (s *Scanner) readDigit(ch rune) (bool, *strings.Builder) {

	if !unicode.IsDigit(ch) {
		return false, nil
	}

	//TODO: optimize
	var number strings.Builder
	for ; unicode.IsDigit(ch); ch = s.readNext() {
		number.WriteRune(ch)
	}

	//TODO: whitesapce const
	if ch != ' ' {
		return false, &number
	}

	s.tokens = append(s.tokens, Token{Integer, number.String()})
	s.r.UnreadRune()
	return true, nil
}

func (s *Scanner) readOperator(ch rune) (bool, *strings.Builder) {

	var op strings.Builder
	op.WriteRune(ch)
	op.WriteRune(s.readNext())

	//TODO: replace by peek?
	ch = s.readNext()
	s.r.UnreadRune()

	opStr := op.String()
	if (ch == ' ' || ch == 0) && opStr == "EQ" || opStr == "GR" || opStr == "LE" {
		s.tokens = append(s.tokens, Token{Operator, opStr})

		return true, nil
	}

	return false, &op
}

func (s *Scanner) readFilename(b *strings.Builder) {
	if b == nil {
		b = new(strings.Builder)
	}

	//TODO: or end of file reached
	for ch := s.readNext(); ch != ' '; ch = s.readNext() {
		b.WriteRune(ch)
	}

	s.r.UnreadRune()
	s.tokens = append(s.tokens, Token{File, b.String()})
}

//TODO: remove
func (s *Scanner) PrintTokens() {
	for _, t := range s.tokens {
		fmt.Printf("%v\n", t)
	}
}

//TODO: remove
func (s *Scanner) PrintValues() {
	for _, t := range s.tokens {
		fmt.Printf("%v", t.Value)
	}
	println()
}

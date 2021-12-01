package lexer

type TokenType int

const (
	Error        TokenType = iota
	LeftBracket            // '['
	RightBracket           // ']'
	Whitespace             // ' '
	Integer                // '[1-9]+'
	Operator               // 'EQ' | 'GR' | 'LE'
	File                   // 'file.ext'
	EOF                    // End of file
)

type Token struct {
	Type  TokenType
	Value string
}

//TODO: remove
func (t *Token) String() string {
	return t.Value
}

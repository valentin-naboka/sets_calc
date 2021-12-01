package lexer

type TokenType int

const (
	LeftBracket  TokenType = iota // '['
	RightBracket                  // ']'
	Whitespace                    // ' '
	Integer                       // '[1-9]+'
	Operator                      // 'EQ' | 'GR' | 'LE'
	File                          // 'filename'
	EOF                           // End of file
)

type Token struct {
	Type  TokenType
	Value string
}

//TODO: remove
func (t *Token) String() string {
	return t.Value
}

package token

type Token struct {
	TokenType TokenType
	Value     string
	Line      int
	Column    int
}

func NewToken(tokenType TokenType, value string, line int, column int) *Token {
	return &Token{
		TokenType: tokenType,
		Value:     value,
		Line:      line,
		Column:    column,
	}
}

func NewErrorToken(value string, line int, column int) *Token {
	return &Token{
		TokenType: Error,
		Value:     value,
		Line:      line,
		Column:    column,
	}
}

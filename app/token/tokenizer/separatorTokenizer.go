package tokenizer

import (
	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
)

const (
	separatorStartState     = iota
	separatorFirstCharState = iota
)

type separatorTokenizer struct {
}

func (t *separatorTokenizer) Tokenize(input *fileSystem.Input, tokens *token.TokenCollection) error {
	beginLineNumber := input.GetLineNumber()
	beginCharIndex := input.GetCharIndex()

	char := input.Byte()
	switch true {
	case IsBracketSeparator(char):
		tokens.Add(token.NewToken(token.BracketSeparator, string(char), beginLineNumber, beginCharIndex))
	case IsQuoteSeparator(char):
		tokens.Add(token.NewToken(token.QuoteSeparator, string(char), beginLineNumber, beginCharIndex))
	case IsIdentifierSeparator(char):
		tokens.Add(token.NewToken(token.IdentifierSeparator, string(char), beginLineNumber, beginCharIndex))
	}

	input.ScanLn()
	return nil
}

func NewSeparatorTokenizer() Tokenizer {
	return &separatorTokenizer{}
}

func IsBracketSeparator(char byte) bool {
	return char == '{' || char == '}' || char == '[' || char == ']' || char == '(' || char == ')'
}

func IsQuoteSeparator(char byte) bool {
	return char == '`' || char == '"' || char == '\''
}

func IsIdentifierSeparator(char byte) bool {
	return char == ',' || char == ';'
}

func IsSeparator(char byte) bool {
	return IsIdentifierSeparator(char) || IsQuoteSeparator(char) || IsBracketSeparator(char)
}

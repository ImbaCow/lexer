package tokenizer

import (
	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
)

type spaceTokenizer struct {
	value string
}

func (t *spaceTokenizer) Tokenize(input *fileSystem.Input, tokens *token.TokenCollection) error {
	beginLineNumber := input.GetLineNumber()
	beginCharIndex := input.GetCharIndex()

	char := input.Byte()
	for !input.IsEOF() && IsSpace(char) {
		t.value += string(char)

		input.ScanLn()
		char = input.Byte()
	}

	tokens.Add(token.NewToken(token.Space, t.value, beginLineNumber, beginCharIndex))
	return nil
}

func NewSpaceTokenizer() Tokenizer {
	return &spaceTokenizer{
		value: "",
	}
}

func IsSpace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\r' || char == '\n'
}

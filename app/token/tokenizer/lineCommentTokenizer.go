package tokenizer

import (
	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
)

type lineCommentTokenizer struct {
	value           string
	beginLineNumber int
	beginCharIndex  int
}

func (t *lineCommentTokenizer) Tokenize(input *fileSystem.Input, tokens *token.TokenCollection) error {
	for input.Scan() {
		char := input.Byte()
		t.value += string(char)
	}

	tokens.Add(token.NewToken(token.LineComment, t.value, t.beginLineNumber, t.beginCharIndex))
	input.ScanLn()
	return nil
}

func NewLineCommentTokenizer(value string, beginLineNumber int, beginCharIndex int) Tokenizer {
	return &lineCommentTokenizer{
		value:           value,
		beginLineNumber: beginLineNumber,
		beginCharIndex:  beginCharIndex,
	}
}

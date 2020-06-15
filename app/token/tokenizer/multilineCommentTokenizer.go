package tokenizer

import (
	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
)

const (
	multilineCommentStartState   = iota
	multilineCommentEndCharState = iota
)

type multilineCommentTokenizer struct {
	currentState    int
	value           string
	beginLineNumber int
	beginCharIndex  int
}

func (t *multilineCommentTokenizer) Tokenize(input *fileSystem.Input, tokens *token.TokenCollection) error {
	for input.ScanLn() {
		char := input.Byte()
		t.value += string(char)

		switch t.currentState {
		case multilineCommentStartState:
			if IsMultilineCommentChar(char) {
				t.currentState = multilineCommentEndCharState
			}
		case multilineCommentEndCharState:
			if !IsMultilineCommentChar(char) {
				if !IsCommentBegin(char) {
					t.currentState = multilineCommentStartState
				} else {
					tokens.Add(token.NewToken(token.MultilineComment, t.value, t.beginLineNumber, t.beginCharIndex))
					input.ScanLn()
					return nil
				}
			}
		}
	}

	(*tokens).Add(token.NewToken(token.MultilineComment, t.value, t.beginLineNumber, t.beginCharIndex))
	return nil
}

func NewMultilineCommentTokenizer(value string, beginLineNumber int, beginCharIndex int) Tokenizer {
	return &multilineCommentTokenizer{
		currentState:    multilineCommentStartState,
		value:           value,
		beginLineNumber: beginLineNumber,
		beginCharIndex:  beginCharIndex,
	}
}

package tokenizer

import (
	"errors"

	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
)

const (
	commentStartState     = iota
	commentFirstCharState = iota
)

type commentTokenizer struct {
	currentState int
	value        string
}

func (t *commentTokenizer) Tokenize(input *fileSystem.Input, tokens *token.TokenCollection) error {
	beginLineNumber := input.GetLineNumber()
	beginCharIndex := input.GetCharIndex()

	for ; !input.IsEOLN(); input.Scan() {
		char := input.Byte()

		switch t.currentState {
		case commentStartState:
			if IsCommentBegin(char) {
				t.value += string(char)
				t.currentState = commentFirstCharState
			} else {
				return errors.New("First character is invalid: " + string(char))
			}

		case commentFirstCharState:
			if IsCommentBegin(char) {
				t.value += string(char)
				tokenizer := NewLineCommentTokenizer(t.value, beginLineNumber, beginCharIndex)
				return tokenizer.Tokenize(input, tokens)
			} else if IsMultilineCommentChar(char) {
				t.value += string(char)
				tokenizer := NewMultilineCommentTokenizer(t.value, beginLineNumber, beginCharIndex)
				return tokenizer.Tokenize(input, tokens)
			} else {
				tokens.Add(token.NewToken(token.ArithmeticOperation, t.value, beginLineNumber, beginCharIndex))
				return nil
			}
		}
	}

	if t.currentState != commentStartState {
		tokens.Add(token.NewToken(token.ArithmeticOperation, t.value, beginLineNumber, beginCharIndex))
		input.ScanLn()
	}
	return nil
}

func NewCommentTokenizer() Tokenizer {
	return &commentTokenizer{
		currentState: commentStartState,
		value:        "",
	}
}

func IsCommentBegin(char byte) bool {
	return char == '/'
}

func IsMultilineCommentChar(char byte) bool {
	return char == '*'
}

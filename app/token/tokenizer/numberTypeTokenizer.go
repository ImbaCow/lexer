package tokenizer

import (
	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
)

const (
	numberTypeStartState = iota
	numberTypeDigitState = iota
)

type byteCheckFunc func(char byte) bool

type numberTypeTokenizer struct {
	currentState    int
	value           string
	beginLineNumber int
	beginCharIndex  int
	digitCheckFunc  byteCheckFunc
	resultTokenType token.TokenType
}

func (t *numberTypeTokenizer) Tokenize(input *fileSystem.Input, tokens *token.TokenCollection) error {
	for input.Scan() {
		char := input.Byte()

		switch t.currentState {
		case numberTypeStartState:
			if t.digitCheckFunc(char) {
				t.value += string(char)
				t.currentState = numberTypeDigitState
			} else {
				tokens.Add(token.NewToken(token.Number10, t.value[:len(t.value)-1], t.beginLineNumber, t.beginCharIndex))
				input.ScanBack(1)
				return nil
			}
		case numberTypeDigitState:
			if t.digitCheckFunc(char) {
				t.value += string(char)
			} else {
				tokens.Add(token.NewToken(t.resultTokenType, t.value, t.beginLineNumber, t.beginCharIndex))
				return nil
			}
		}
	}

	switch t.currentState {
	case numberTypeStartState:
		tokens.Add(token.NewToken(token.Number10, t.value[:len(t.value)-1], t.beginLineNumber, t.beginCharIndex))
		input.ScanBack(1)
	case numberTypeDigitState:
		tokens.Add(token.NewToken(t.resultTokenType, t.value, t.beginLineNumber, t.beginCharIndex))
		input.ScanLn()
	}

	return nil
}

func NewNumberTypeTokenizer(value string, beginLineNumber int, beginCharIndex int, digitCheckFunc byteCheckFunc, resultTokenType token.TokenType) Tokenizer {
	return &numberTypeTokenizer{
		currentState:    numberTypeStartState,
		value:           value,
		beginLineNumber: beginLineNumber,
		beginCharIndex:  beginCharIndex,
		digitCheckFunc:  digitCheckFunc,
		resultTokenType: resultTokenType,
	}
}

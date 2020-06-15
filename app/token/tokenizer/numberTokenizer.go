package tokenizer

import (
	"errors"

	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
)

const (
	numberStartState   = iota
	numberFirst0State  = iota
	numberIntegerState = iota
)

type numberTokenizer struct {
	currentState int
	value        string
}

func (t *numberTokenizer) Tokenize(input *fileSystem.Input, tokens *token.TokenCollection) error {
	beginLineNumber := input.GetLineNumber()
	beginCharIndex := input.GetCharIndex()

	for ; !input.IsEOLN(); input.Scan() {
		char := input.Byte()

		switch t.currentState {
		case numberStartState:
			if char == '0' {
				t.currentState = numberFirst0State
			} else if IsDigit10(char) {
				t.currentState = numberIntegerState
			} else if IsFloatPartSeparator(char) {
				t.value += string(char)
				tokenizer := NewNumberTypeTokenizer(t.value, beginLineNumber, beginCharIndex, IsDigit10, token.Float)
				return tokenizer.Tokenize(input, tokens)
			} else {
				return errors.New("First character is invalid: " + string(char))
			}
			t.value += string(char)
		case numberIntegerState:
			if IsDigit10(char) {
				t.value += string(char)
			} else if IsFloatPartSeparator(char) {
				t.value += string(char)
				tokenizer := NewNumberTypeTokenizer(t.value, beginLineNumber, beginCharIndex, IsDigit10, token.Float)
				return tokenizer.Tokenize(input, tokens)
			} else {
				tokens.Add(token.NewToken(token.Number10, t.value, beginLineNumber, beginCharIndex))
				return nil
			}
		case numberFirst0State:
			switch true {
			case IsDigit10(char):
				t.value += string(char)
				t.currentState = numberIntegerState
			case IsDigit16PartSeparator(char):
				t.value += string(char)
				tokenizer := NewNumberTypeTokenizer(t.value, beginLineNumber, beginCharIndex, IsDigit16, token.Number16)
				return tokenizer.Tokenize(input, tokens)
			case IsDigit2PartSeparator(char):
				t.value += string(char)
				tokenizer := NewNumberTypeTokenizer(t.value, beginLineNumber, beginCharIndex, IsDigit2, token.Number2)
				return tokenizer.Tokenize(input, tokens)
			case IsFloatPartSeparator(char):
				t.value += string(char)
				tokenizer := NewNumberTypeTokenizer(t.value, beginLineNumber, beginCharIndex, IsDigit10, token.Float)
				return tokenizer.Tokenize(input, tokens)
			default:
				tokens.Add(token.NewToken(token.Number10, t.value, beginLineNumber, beginCharIndex))
				return nil
			}
		}
	}

	if t.currentState != numberStartState {
		tokens.Add(token.NewToken(token.Number10, t.value, beginLineNumber, beginCharIndex))
		input.ScanLn()
	}
	return nil
}

func NewNumberTokenizer() Tokenizer {
	return &numberTokenizer{
		currentState: numberStartState,
		value:        "",
	}
}

func IsDigit10(char byte) bool {
	return char >= '0' && char <= '9'
}

func IsDigit16PartSeparator(char byte) bool {
	return char == 'x'
}

func IsDigit2PartSeparator(char byte) bool {
	return char == 'b'
}

func IsFloatPartSeparator(char byte) bool {
	return char == '.'
}

func IsDigit2(char byte) bool {
	return char == '0' || char == '1'
}

func IsDigit16(char byte) bool {
	return IsDigit10(char) || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')
}

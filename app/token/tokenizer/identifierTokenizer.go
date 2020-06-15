package tokenizer

import (
	"errors"

	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
)

const (
	identifierStartState = iota
	identifierCharState  = iota
)

type identifierTokenizer struct {
	currentState int
	value        string
}

func (t *identifierTokenizer) Tokenize(input *fileSystem.Input, tokens *token.TokenCollection) error {
	beginLineNumber := input.GetLineNumber()
	beginCharIndex := input.GetCharIndex()

	for ; !input.IsEOLN(); input.Scan() {
		char := input.Byte()

		switch t.currentState {
		case identifierStartState:
			if IsLetter(char) {
				t.value += string(char)
				t.currentState = identifierCharState
			} else {
				return errors.New("First character is invalid: " + string(char))
			}
		case identifierCharState:
			if IsIdentifier(char) {
				t.value += string(char)
			} else {
				if IsKeyword(t.value) {
					tokens.Add(token.NewToken(token.Keyword, t.value, beginLineNumber, beginCharIndex))
				} else {
					tokens.Add(token.NewToken(token.Identifier, t.value, beginLineNumber, beginCharIndex))
				}
				return nil
			}
		}
	}

	if t.currentState != identifierStartState {
		if IsKeyword(t.value) {
			tokens.Add(token.NewToken(token.Keyword, t.value, beginLineNumber, beginCharIndex))
		} else if t.currentState != identifierStartState {
			tokens.Add(token.NewToken(token.Identifier, t.value, beginLineNumber, beginCharIndex))
		}
		input.ScanLn()
	}

	return nil
}

func NewIdentifierTokenizer() Tokenizer {
	return &identifierTokenizer{
		currentState: identifierStartState,
		value:        "",
	}
}

func IsLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func IsIdentifier(char byte) bool {
	return IsLetter(char) || IsDigit10(char) || char == '_'
}

func IsKeyword(word string) bool {
	for _, keyword := range getKeywords() {
		if keyword == word {
			return true
		}
	}
	return false
}

func getKeywords() []string {
	return []string{
		"if",
		"else",
		"switch",
		"case",
		"default",
		"break",
		"int",
		"float",
		"char",
		"for",
		"while",
		"do",
		"void",
		"return",
		"const",
		"continue",
	}
}

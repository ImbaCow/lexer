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

const identifierMaxLength = 64

type identifierTokenizer struct {
	currentState    int
	value           string
	beginLineNumber int
	beginCharIndex  int
}

func (t *identifierTokenizer) Tokenize(input *fileSystem.Input, tokens *token.TokenCollection) error {
	t.beginLineNumber = input.GetLineNumber()
	t.beginCharIndex = input.GetCharIndex()

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
			if IsIdentifier(char) && len(t.value) < identifierMaxLength {
				t.value += string(char)
			} else {
				return validateIdentifier(t.makeIdentifier(), tokens)
			}
		}
	}

	if t.currentState != identifierStartState {
		input.ScanLn()
		return validateIdentifier(t.makeIdentifier(), tokens)
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

func (t *identifierTokenizer) makeIdentifier() *token.Token {
	if IsKeyword(t.value) {
		return token.NewToken(token.Keyword, t.value, t.beginLineNumber, t.beginCharIndex)
	}
	return token.NewToken(token.Identifier, t.value, t.beginLineNumber, t.beginCharIndex)
}

func validateIdentifier(tokenItem *token.Token, tokens *token.TokenCollection) error {
	if tokenItem.TokenType != token.Keyword || len(tokenItem.Value) <= 256 {
		return tokens.Add(tokenItem)
	}
	return errors.New("Identifier too long: " + tokenItem.Value)
}

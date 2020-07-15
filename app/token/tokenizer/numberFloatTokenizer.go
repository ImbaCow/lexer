package tokenizer

import (
	"errors"

	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
)

const (
	numberFloatStartState             = iota
	numberFloatDigitState             = iota
	numberFloatOneDigitState          = iota
	numberFloatENotationBeginState    = iota
	numberFloatENotationState         = iota
	numberFloatENotationNegativeState = iota
)

type numberFloatTokenizer struct {
	currentState    int
	value           string
	floatValue      string
	scientificValue string
	isENegative     bool
	beginLineNumber int
	beginCharIndex  int
}

func (t *numberFloatTokenizer) Tokenize(input *fileSystem.Input, tokens *token.TokenCollection) error {
	for input.Scan() {
		char := input.Byte()

		switch t.currentState {
		case numberFloatStartState:
			if IsDigit10(char) {
				t.floatValue += string(char)
				if len(t.value) <= 2 {
					t.currentState = numberFloatOneDigitState
				} else {
					t.currentState = numberFloatDigitState
				}
			} else {
				tokens.Add(token.NewToken(token.Number10, t.value[:len(t.value)-1], t.beginLineNumber, t.beginCharIndex))
				input.ScanBack(1)
				return nil
			}
		case numberFloatOneDigitState:
			if IsDigit10(char) {
				t.floatValue += string(char)
				t.currentState = numberFloatDigitState
			} else if IsENotationPartBegin(char) && t.value[0] == '0' {
				t.currentState = numberFloatENotationBeginState
			} else {
				tokens.Add(token.NewToken(token.Float, t.value+t.floatValue, t.beginLineNumber, t.beginCharIndex))
				return nil
			}
		case numberFloatENotationBeginState:
			if char == '-' {
				t.isENegative = true
				t.currentState = numberFloatENotationNegativeState
			} else if IsDigit10(char) {
				t.scientificValue += string(char)
				t.currentState = numberFloatENotationState
			} else {
				tokens.Add(token.NewToken(token.Float, t.value+t.floatValue, t.beginLineNumber, t.beginCharIndex))
				input.ScanBack(1)
				return nil
			}
		case numberFloatENotationNegativeState:
			if IsDigit10(char) {
				t.scientificValue += string(char)
				t.currentState = numberFloatENotationState
			} else {
				tokens.Add(token.NewToken(token.Float, t.value+t.floatValue, t.beginLineNumber, t.beginCharIndex))
				input.ScanBack(2)
				return nil
			}
		case numberFloatENotationState:
			if IsDigit10(char) {
				t.scientificValue += string(char)
			} else {
				tokens.Add(token.NewToken(token.ENotationNumber, t.value+t.floatValue+t.scientificValue, t.beginLineNumber, t.beginCharIndex))
				return nil
			}
		case numberFloatDigitState:
			if IsDigit10(char) {
				t.floatValue += string(char)
			} else {
				tokens.Add(token.NewToken(token.Float, t.value+t.floatValue, t.beginLineNumber, t.beginCharIndex))
				return nil
			}
		}
	}

	switch t.currentState {
	case numberFloatStartState:
		tokens.Add(token.NewToken(token.Number10, t.value[:len(t.value)-1], t.beginLineNumber, t.beginCharIndex))
		input.ScanBack(1)
	case numberFloatENotationBeginState:
		tokens.Add(token.NewToken(token.Float, t.value+t.floatValue, t.beginLineNumber, t.beginCharIndex))
		input.ScanBack(1)
	case numberFloatENotationNegativeState:
		tokens.Add(token.NewToken(token.Float, t.value+t.floatValue, t.beginLineNumber, t.beginCharIndex))
		input.ScanBack(2)
	case numberFloatENotationState:
		tokens.Add(token.NewToken(token.ENotationNumber, t.value+t.floatValue+t.scientificValue, t.beginLineNumber, t.beginCharIndex))
		input.ScanLn()
	case numberFloatDigitState, numberFloatOneDigitState:
		tokens.Add(token.NewToken(token.Float, t.value+t.floatValue, t.beginLineNumber, t.beginCharIndex))
		input.ScanLn()
	}

	return nil
}

func NewNumberFloatTokenizer(value string, beginLineNumber int, beginCharIndex int) Tokenizer {
	return &numberFloatTokenizer{
		currentState:    numberFloatStartState,
		value:           value,
		scientificValue: "",
		beginLineNumber: beginLineNumber,
		beginCharIndex:  beginCharIndex,
	}
}

func (t *numberFloatTokenizer) makeFloat(input *fileSystem.Input) *token.Token {
	switch t.currentState {
	case numberFloatStartState:
		input.ScanBack(1)
		return token.NewToken(token.Number10, t.value[:len(t.value)-1], t.beginLineNumber, t.beginCharIndex)
	case numberFloatENotationBeginState:
		input.ScanBack(1)
		return token.NewToken(token.Float, t.value+t.floatValue, t.beginLineNumber, t.beginCharIndex)
	case numberFloatENotationNegativeState:
		input.ScanBack(2)
		return token.NewToken(token.Float, t.value+t.floatValue, t.beginLineNumber, t.beginCharIndex)
	case numberFloatENotationState:
		input.ScanLn()
		return token.NewToken(token.ENotationNumber, t.value+t.floatValue+t.scientificValue, t.beginLineNumber, t.beginCharIndex)
	default:
		input.ScanLn()
		return token.NewToken(token.Float, t.value+t.floatValue, t.beginLineNumber, t.beginCharIndex)
	}
}

func validateFloat(tokenItem *token.Token, tokens *token.TokenCollection) error {
	if tokenItem.TokenType == token.Number10 {
		return validateNumber10(tokenItem, tokens)
	} else if tokenItem.TokenType == token.Float {

	} else if tokenItem.TokenType == token.ENotationNumber {

	}
	return errors.New("Identifier too long: " + tokenItem.Value)
}

func IsENotationPartBegin(char byte) bool {
	return char == 'e'
}

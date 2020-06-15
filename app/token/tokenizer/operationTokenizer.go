package tokenizer

import (
	"errors"

	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
)

const (
	operationStartState     = iota
	operationFirstCharState = iota
)

type operationTokenizer struct {
	currentState int
	firstChar    byte
}

func (t *operationTokenizer) Tokenize(input *fileSystem.Input, tokens *token.TokenCollection) error {
	beginLineNumber := input.GetLineNumber()
	beginCharIndex := input.GetCharIndex()

	for ; !input.IsEOLN(); input.Scan() {
		char := input.Byte()

		switch t.currentState {
		case operationStartState:
			if IsOperation(char) {
				t.firstChar += char
				t.currentState = operationFirstCharState
			} else {
				return errors.New("First character is invalid: " + string(char))
			}
		case operationFirstCharState:
			switch t.firstChar {
			case '|':
				if char == '|' {
					tokens.Add(token.NewToken(token.ArithmeticOperation, "||", beginLineNumber, beginCharIndex))
					input.ScanLn()
					return nil
				} else {
					tokens.Add(token.NewToken(token.ArithmeticOperation, "|", beginLineNumber, beginCharIndex))
					return nil
				}
			case '!':
				if char == '=' {
					tokens.Add(token.NewToken(token.CompareOperation, "!=", beginLineNumber, beginCharIndex))
					input.ScanLn()
					return nil
				} else {
					tokens.Add(token.NewToken(token.BitOperation, "!", beginLineNumber, beginCharIndex))
					return nil
				}
			case '+':
				if char == '+' {
					tokens.Add(token.NewToken(token.ArithmeticOperation, "++", beginLineNumber, beginCharIndex))
					input.ScanLn()
					return nil
				} else {
					tokens.Add(token.NewToken(token.ArithmeticOperation, "+", beginLineNumber, beginCharIndex))
					return nil
				}
			case '-':
				if char == '-' {
					tokens.Add(token.NewToken(token.ArithmeticOperation, "--", beginLineNumber, beginCharIndex))
					input.ScanLn()
					return nil
				} else {
					tokens.Add(token.NewToken(token.ArithmeticOperation, "-", beginLineNumber, beginCharIndex))
					return nil
				}
			case '%':
				tokens.Add(token.NewToken(token.ArithmeticOperation, "%", beginLineNumber, beginCharIndex))
				return nil
			case '^':
				tokens.Add(token.NewToken(token.ArithmeticOperation, "^", beginLineNumber, beginCharIndex))
				return nil
			case '=':
				if char == '=' {
					tokens.Add(token.NewToken(token.CompareOperation, "==", beginLineNumber, beginCharIndex))
					input.ScanLn()
					return nil
				} else {
					tokens.Add(token.NewToken(token.AssignOperation, "=", beginLineNumber, beginCharIndex))
					return nil
				}
			case '<':
				if char == '=' {
					tokens.Add(token.NewToken(token.CompareOperation, "<=", beginLineNumber, beginCharIndex))
					input.ScanLn()
					return nil
				} else {
					tokens.Add(token.NewToken(token.CompareOperation, "<", beginLineNumber, beginCharIndex))
					return nil
				}
			case '>':
				if char == '=' {
					tokens.Add(token.NewToken(token.CompareOperation, ">=", beginLineNumber, beginCharIndex))
					input.ScanLn()
					return nil
				} else {
					tokens.Add(token.NewToken(token.CompareOperation, ">", beginLineNumber, beginCharIndex))
					return nil
				}
			case '&':
				if char == '&' {
					tokens.Add(token.NewToken(token.BooleanOperation, "&&", beginLineNumber, beginCharIndex))
					input.ScanLn()
					return nil
				} else {
					tokens.Add(token.NewToken(token.BitOperation, "&", beginLineNumber, beginCharIndex))
					return nil
				}
			case '/':
				tokens.Add(token.NewToken(token.ArithmeticOperation, "/", beginLineNumber, beginCharIndex))
				return nil
			case '*':
				tokens.Add(token.NewToken(token.ArithmeticOperation, "*", beginLineNumber, beginCharIndex))
				return nil
			}
		}
	}

	if t.currentState != operationStartState {
		var tokenType token.TokenType = token.Error

		switch true {
		case IsArithmeticOperation(t.firstChar):
			tokenType = token.ArithmeticOperation
		case IsCompareOperation(t.firstChar):
			tokenType = token.CompareOperation
		case IsBitOperation(t.firstChar):
			tokenType = token.BitOperation
		case IsAssignOperation(t.firstChar):
			tokenType = token.AssignOperation
		default:
			tokens.Add(token.NewErrorToken(string(t.firstChar), beginLineNumber, beginCharIndex))
			return errors.New("Unhandled operation: " + string(t.firstChar))
		}
		tokens.Add(token.NewToken(tokenType, string(t.firstChar), beginLineNumber, beginCharIndex))
		input.ScanLn()
	}
	return nil
}

func NewOperationTokenizer() Tokenizer {
	return &operationTokenizer{
		currentState: commentStartState,
		firstChar:    0,
	}
}

func IsOperation(char byte) bool {
	return char == '|' ||
		char == '!' ||
		char == '+' ||
		char == '%' ||
		char == '^' ||
		char == '-' ||
		char == '=' ||
		char == '<' ||
		char == '>' ||
		char == '&' ||
		char == '/' ||
		char == '*'
}

func IsArithmeticOperation(char byte) bool {
	return char == '+' ||
		char == '%' ||
		char == '-' ||
		char == '/' ||
		char == '*'
}

func IsCompareOperation(char byte) bool {
	return char == '>' ||
		char == '<'
}

func IsBitOperation(char byte) bool {
	return char == '&' ||
		char == '|' ||
		char == '!'
}

func IsAssignOperation(char byte) bool {
	return char == '='
}

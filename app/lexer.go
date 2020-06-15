package app

import (
	"errors"
	"io"

	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
	"github.com/ImbaCow/lexer/app/token/tokenizer"
)

type Lexer struct {
}

func (l *Lexer) TokenizeFile(inputReader io.Reader, outputWriter io.Writer) error {
	input := fileSystem.CreateInput(inputReader)
	tokens := token.NewTokenCollection(outputWriter)

	for input.ScanLn(); !input.IsEOF(); {
		char := input.Byte()
		tokenizer := getTokenizer(char)
		if tokenizer == nil {
			tokens.Add(token.NewErrorToken(string(char), input.GetLineNumber(), input.GetCharIndex()))
			return errors.New("Undefined character: " + string(char))
		}
		err := tokenizer.Tokenize(input, tokens)
		if err != nil {
			return err
		}
	}
	return nil
}

func getTokenizer(char byte) tokenizer.Tokenizer {
	switch true {
	case tokenizer.IsDigit10(char):
		return tokenizer.NewNumberTokenizer()
	case tokenizer.IsCommentBegin(char):
		return tokenizer.NewCommentTokenizer()
	case tokenizer.IsOperation(char):
		return tokenizer.NewOperationTokenizer()
	case tokenizer.IsLetter(char):
		return tokenizer.NewIdentifierTokenizer()
	case tokenizer.IsSpace(char):
		return tokenizer.NewSpaceTokenizer()
	case tokenizer.IsSeparator(char):
		return tokenizer.NewSeparatorTokenizer()
	default:
		return nil
	}
}

func NewLexer() *Lexer {
	return &Lexer{}
}

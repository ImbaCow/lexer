package token

import (
	"fmt"
	"io"
)

type TokenCollection struct {
	tokens       []*Token
	outputWriter io.Writer
}

func (tc *TokenCollection) Add(tokenItem *Token) error {
	tc.tokens = append(tc.tokens, tokenItem)
	message := fmt.Sprintf("Token: type %v value '%v' line: %v position: %v \n", TokenTypeToString(tokenItem.TokenType), tokenItem.Value, tokenItem.Line, tokenItem.Column)
	_, err := tc.outputWriter.Write([]byte(message))
	return err
}

func (tc *TokenCollection) Iterate() []*Token {
	return tc.tokens
}

func NewTokenCollection(outputWriter io.Writer) *TokenCollection {
	return &TokenCollection{
		tokens:       make([]*Token, 0),
		outputWriter: outputWriter,
	}
}

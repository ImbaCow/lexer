package token

import (
	"fmt"
	"io"
)

type TokenCollection struct {
	tokens       []*Token
	outputWriter io.Writer
}

func (tc *TokenCollection) Add(token *Token) {
	tc.tokens = append(tc.tokens, token)
	message := fmt.Sprintf("Token: type %v value '%v' line: %v position: %v \n", TokenTypeToString(token.TokenType), token.Value, token.Line, token.Column)
	tc.outputWriter.Write([]byte(message))
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

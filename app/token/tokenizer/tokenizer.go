package tokenizer

import (
	"github.com/ImbaCow/lexer/app/fileSystem"
	"github.com/ImbaCow/lexer/app/token"
)

type Tokenizer interface {
	Tokenize(*fileSystem.Input, *token.TokenCollection) error
}

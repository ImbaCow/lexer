package token

import "fmt"

type TokenType int

const (
	Operator            = iota
	Number10            = iota
	Number16            = iota
	Number2             = iota
	Float               = iota
	Identifier          = iota
	Keyword             = iota
	LineComment         = iota
	MultilineComment    = iota
	ArithmeticOperation = iota
	BooleanOperation    = iota
	BitOperation        = iota
	CompareOperation    = iota
	AssignOperation     = iota
	Error               = iota
	BracketSeparator    = iota
	QuoteSeparator      = iota
	IdentifierSeparator = iota
	Space               = iota
)

func TokenTypeToString(tokenType TokenType) string {
	switch tokenType {
	case Operator:
		return "Operator"
	case Number10:
		return "Number10"
	case Number16:
		return "Number16"
	case Number2:
		return "Number2"
	case Float:
		return "Float"
	case Identifier:
		return "Identifier"
	case Keyword:
		return "Keyword"
	case LineComment:
		return "LineComment"
	case MultilineComment:
		return "MultilineComment"
	case ArithmeticOperation:
		return "ArithmeticOperation"
	case BooleanOperation:
		return "BooleanOperation"
	case BitOperation:
		return "BitOperation"
	case CompareOperation:
		return "CompareOperation"
	case AssignOperation:
		return "AssignOperation"
	case Error:
		return "Error"
	case BracketSeparator:
		return "BracketSeparator"
	case QuoteSeparator:
		return "QuoteSeparator"
	case IdentifierSeparator:
		return "IdentifierSeparator"
	case Space:
		return "Space"
	default:
		return fmt.Sprint(tokenType)
	}
}

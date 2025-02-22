package token

// explicit type instead of directly using string as it offers a way to limit possibilities of tokentype, unless author explicitly typecasts
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN    = "ASSIGN"
	PLUS      = "PLUS"
	MINUS     = "MINUS"
	MULTIPLY  = "MULTIPLY"
	DIVIDE    = "DIVIDE"
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	// Parenthesis
	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

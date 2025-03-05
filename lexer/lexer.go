package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input    string
	position int
	nextPos  int
	char     byte
}

var mapTokenType = map[byte]token.TokenType{
	//parenthesis
	'{': token.LBRACE,
	'}': token.RBRACE,
	'(': token.LPAREN,
	')': token.RPAREN,
	'/': token.DIVIDE,

	//operators
	'=': token.ASSIGN,
	'*': token.MULTIPLY,
	'+': token.PLUS,
	'-': token.MINUS,
	'!': token.NOT,
	'<': token.LESSTHAN,
	'>': token.MORETHAN,

	// separators
	',': token.COMMA,
	';': token.SEMICOLON,
}

func (lex *Lexer) NextToken() token.Token {
	lex.skipWhitespace()

	var tok token.Token
	if tt, ok := mapTokenType[lex.char]; ok {
		switch {
		default:
			tok = newToken(tt, lex.char)
		case token.ASSIGN == tt && token.ASSIGN == mapTokenType[lex.peekChar()]:
			tok = token.Token{Type: token.EQUALITY, Literal: lex.readTwice()}
		case token.NOT == tt && token.ASSIGN == mapTokenType[lex.peekChar()]:
			tok = token.Token{Type: token.NEQUALITY, Literal: lex.readTwice()}
		case token.LESSTHAN == tt && token.ASSIGN == mapTokenType[lex.peekChar()]:
			tok = token.Token{Type: token.EQ_OR_LESS, Literal: lex.readTwice()}
		case token.MORETHAN == tt && token.ASSIGN == mapTokenType[lex.peekChar()]:
			tok = token.Token{Type: token.EQ_OR_MORE, Literal: lex.readTwice()}
		}
	} else {
		switch lex.char {
		case 0:
			tok.Type, tok.Literal = token.EOF, ""
		default:
			if isLetter(lex.char) {
				tok.Literal = lex.readIdentifier()
				tok.Type = checkIfKeyword(tok.Literal)
				return tok
			} else if isDigit(lex.char) {
				tok.Literal = lex.readDigit()
				tok.Type = token.INT
				return tok
			} else {
				tok = newToken(token.ILLEGAL, lex.char)
			}
		}
	}

	lex.readChar()
	return tok
}

func (lex *Lexer) readTwice() string {
	first := lex.char
	lex.readChar()
	second := lex.char
	return string(first) + string(second)
}

func (lex *Lexer) readDigit() string {
	position := lex.position
	for isDigit(lex.char) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func (lex *Lexer) skipWhitespace() {
	for lex.char == ' ' || lex.char == '\t' || lex.char == '\n' || lex.char == '\r' {
		lex.readChar()
	}
}

// map of known keywords
var keywords = map[string]token.TokenType{
	"let":       token.LET,
	"fn":        token.FUNCTION,
	"yes":       token.YES,
	"no":        token.NO,
	"when":      token.WHEN,
	"otherwise": token.OTHERWISE,
	"send":      token.SEND,
}

// returns whether the string is among known keywords
func checkIfKeyword(s string) token.TokenType {
	if toktype, ok := keywords[s]; ok {
		return toktype
	}
	return token.IDENT
}

func (lex *Lexer) readIdentifier() string {
	position := lex.position
	for isLetter(lex.char) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

func isLetter(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z' || (b == '_'))
}

func newToken(tt token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tt,
		Literal: string(char),
	}
}

func (lex *Lexer) readChar() {
	if lex.nextPos >= len(lex.input) {
		lex.char = 0
	} else {
		lex.char = lex.input[lex.nextPos]
	}
	lex.position = lex.nextPos
	lex.nextPos++
}

func (lex *Lexer) peekChar() byte {
	if lex.nextPos >= len(lex.input) {
		return 0
	}
	return lex.input[lex.nextPos]
}

func NewLexer(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

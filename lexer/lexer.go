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

func (lex *Lexer) NextToken() token.Token {
	lex.skipWhitespace()

	var tok token.Token
	switch lex.char {
	case '=':
		tok = newToken(token.ASSIGN, lex.char)
	case '{':
		tok = newToken(token.LBRACE, lex.char)
	case '}':
		tok = newToken(token.RBRACE, lex.char)
	case '(':
		tok = newToken(token.LPAREN, lex.char)
	case ')':
		tok = newToken(token.RPAREN, lex.char)
	case '/':
		tok = newToken(token.DIVIDE, lex.char)
	case '*':
		tok = newToken(token.MULTIPLY, lex.char)
	case '+':
		tok = newToken(token.PLUS, lex.char)
	case '-':
		tok = newToken(token.MINUS, lex.char)
	case ',':
		tok = newToken(token.COMMA, lex.char)
	case ';':
		tok = newToken(token.SEMICOLON, lex.char)
	case 0:
		tok.Type, tok.Literal = token.EOF, ""
	default:
		if isLetter(lex.char) {
			tok.Literal = lex.readIdentifier()
			tok.Type = lookupType(tok.Literal)
			return tok
		} else if isDigit(lex.char) {
			tok.Literal = lex.readDigit()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.char)
		}
	}
	lex.readChar()
	return tok
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

func lookupType(s string) token.TokenType {
	if toktype, ok := keywords[s]; ok {
		return toktype
	}
	return token.IDENT
}

var keywords = map[string]token.TokenType{
	"let": token.LET,
	"fn":  token.FUNCTION,
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
	return token.Token{Type: tt, Literal: string(char)}
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

func NewLexer(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

package lexer

import (
	"monkey/token"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

type validation struct {
	expectedType    token.TokenType
	expectedLiteral string
}

var l = &logrus.Logger{
	Out:   os.Stderr,
	Level: logrus.InfoLevel,
	// Level: logrus.DebugLevel,
	Formatter: &logrus.TextFormatter{
		// DisableColors:   true,
		ForceColors:     true,
		TimestampFormat: "2006",
		FullTimestamp:   true,
	},
}

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	when (5<10) {
		send yes;
	} otherwise {
		send no;
	}
	`

	validations := []validation{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.WHEN, "when"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LESSTHAN, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.SEND, "send"},
		{token.YES, "yes"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.OTHERWISE, "otherwise"},
		{token.LBRACE, "{"},
		{token.SEND, "send"},
		{token.NO, "no"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}

	lex := NewLexer(input)
	for i, v := range validations {
		tok := lex.NextToken()
		l.Debugf("parsing token %v", tok)
		if tok.Type != v.expectedType {
			t.Fatalf("tests[%v] - INVALID tokenType want=%v have=%v for token=%v", i, v.expectedType, tok.Type, tok)
		}
		if tok.Literal != v.expectedLiteral {
			t.Fatalf("tests[%v] - INVALID tokenLiteral want=%v have=%v", i, v.expectedLiteral, tok.Literal)
		}
		l.Infof("parsed token %v", tok)
	}
}

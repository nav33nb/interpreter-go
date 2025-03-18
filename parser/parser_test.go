package parser

import (
	"monkey/lexer"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

var l = &logrus.Logger{
	Out: os.Stderr,
	// Level: logrus.InfoLevel,
	Level: logrus.DebugLevel,
	Formatter: &logrus.TextFormatter{
		// DisableColors:   true,
		ForceColors:     true,
		TimestampFormat: "2006",
		FullTimestamp:   true,
	},
}

func TestLetStatements(t *testing.T) {
	input := `
	let x 5;
	let 6;
	let somefoo = 4
	send 56;`
	expected := 3

	lex := lexer.NewLexer(input)
	p := New(lex)

	program := p.ParseProgram()
	// check length first
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	} else if len(program.Statements) != expected {
		t.Fatalf("parsed program has invalid num of statements: EXPECTED %v, HAVE %v", expected, len(program.Statements))
	}

	l.Debugf("Parsed %v statements", len(program.Statements))
	// check statement, once length is correct
	tests := []struct{ expectedIdentifier string }{
		{"x"},
		{"y"},
		{"somefoo"},
		// {"send"},
	}
	for i := range tests {
		stmt := program.Statements[i]
		err := program.Errors[i]
		l.Debugf("line:%v, parsed: %v, err: %v", i, stmt, err)
		if err != nil {
			t.Fail()
		}
		// if stmt.TokenLiteral() != ei.expectedIdentifier {
		// 	t.Errorf("invalid statment ei: WANT %v, HAVE %v", ei.expectedIdentifier, stmt.TokenLiteral())
		// }
	}
}

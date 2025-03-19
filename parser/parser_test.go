package parser

import (
	"monkey/ast"
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
	let x = 5;
	let y=7;
	let somefoo = 4;` // second will fail, no assign operator in LET
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

	l.Debugf("Parsed %v statements a", len(program.Statements))
	// check statement, once length is correct
	tests := []struct{ expectedIdentifier string }{
		{"x"},
		{"y"},
		{"somefoo"},
	}
	for i := range tests {
		stmt := program.Statements[i]
		err := program.Errors[i]
		printer(i, stmt.ToString(), err)
		if err != nil {
			t.Fail()
			continue
		}
	}
}

func TestSendStatements(t *testing.T) {
	input := `
	send 400;
	send 4;` // send <expression>
	expected := 2

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
	for i, stmt := range program.Statements {
		err := program.Errors[i]
		printer(i, stmt.ToString(), err)
		stmt = stmt.(*ast.SendStatement)
		//fmt.Printf("%T\n", stmt)
		if err != nil {
			t.Fail()
			continue
		}
		// sendStmt := stmt.(*ast.SendStatement)
		// l.Debugf("%T", sendStmt)
	}
}

func printer(line int, content string, err error) {
	l.Debugf("LINE:%v, CONTENT: %v, ERR: %v", line, content, err)
}

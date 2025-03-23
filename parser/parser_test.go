package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
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

func TestString(t *testing.T) {

	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.ToString() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.ToString())
	}
}

func printer(line int, content string, err error) {
	l.Debugf("LINE:%v, CONTENT: %v, ERR: %v", line, content, err)
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;" // this is an expression
	expected := 1

	lex := lexer.NewLexer(input)
	p := New(lex)

	program := p.ParseProgram()
	// check length first
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	} else if len(program.Statements) != expected {
		t.Fatalf("parsed program has invalid num of statements: want %v, got %v", expected, len(program.Statements))
	}

	l.Debugf("Parsed %v statements", len(program.Statements))
	for i, stmt := range program.Statements {
		err := program.Errors[i] // some error in parsing the tokens themselves
		printer(i, stmt.ToString(), err)
		if err != nil {
			t.Fail()
			continue
		}

		s, ok := stmt.(*ast.ExpressionStatement) // type assertion to expression statement
		if !ok {
			t.Errorf("parsed statement is not ExpressionStatement. got=%T", stmt)
			continue
		}

		id, ok := s.Expression.(*ast.Identifier) // expression statement should have valid identifier
		if !ok {
			t.Errorf("exp not *ast.Identifier. got=%T", s.Expression)
			continue
		}
		if id.Value != "foobar" { // the valid identifier type should be correct
			t.Errorf("ident.Value not %s. got=%s", "foobar", id.Value)
			continue
		}
		if id.TokenLiteral() != "foobar" { // token literal should be correct
			t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
				id.TokenLiteral())
			continue
		}
		// sendStmt := s.(*ast.SendStatement)
		// l.Debugf("%T", sendStmt)
	}

}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;" // this is an expression, RIGHT SIDE ?
	expected := 1

	lex := lexer.NewLexer(input)
	p := New(lex)

	program := p.ParseProgram()
	// check length first
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	} else if len(program.Statements) != expected {
		t.Fatalf("parsed program has invalid num of statements: want %v, got %v", expected, len(program.Statements))
	}

	l.Debugf("Parsed %v statements", len(program.Statements))
	for i, stmt := range program.Statements {
		err := program.Errors[i] // some error in parsing the tokens themselves
		printer(i, stmt.ToString(), err)
		if err != nil {
			t.Fail()
			continue
		}

		s, ok := stmt.(*ast.ExpressionStatement) // type assertion to expression statement
		if !ok {
			t.Errorf("parsed statement is not ExpressionStatement. got=%T", stmt)
			continue
		}

		//fmt.Printf("s is %T\n", s)
		literal, ok := s.Expression.(*ast.IntegerLiteral) // expression statement should have valid identifier
		if !ok {
			t.Errorf("exp not *ast.IntegerLiteral. got=%T", s.Expression)
			continue
		}
		if literal.Value != 5 { // the valid identifier type should be correct
			t.Errorf("ident.Value not %v. got=%v", "foobar", literal.Value)
			continue
		}
		if literal.TokenLiteral() != "5" { // token literal should be correct
			t.Errorf("ident.TokenLiteral not %v. got=%v", "5",
				literal.TokenLiteral())
			continue
		}
		// sendStmt := s.(*ast.SendStatement)
		// l.Debugf("%T", sendStmt)
	}

}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operatar string
		intValue int
	}{
		{"!5", "!", 5},
		{"-14", "-", 14},
	}
	for _, tt := range prefixTests {
		lex := lexer.NewLexer(tt.input)
		p := New(lex)
		program := p.ParseProgram()
		// check length first
		l.Debugf("Parsed %v statements: %v", len(program.Statements), program.Statements)
		if program == nil {
			t.Fatalf("ParseProgram() returned nil")
		} else if len(program.Statements) != 1 {
			t.Fatalf("parsed program has invalid num of statements: want %v, got %v", 1, len(program.Statements))
		}
		for i, stmt := range program.Statements {
			err := program.Errors[i] // some error in parsing the tokens themselves
			printer(i, stmt.ToString(), err)
			if err != nil {
				t.Fail()
				continue
			}

			s, ok := stmt.(*ast.ExpressionStatement) // type assertion to expression statement
			if !ok {
				t.Errorf("parsed statement is not ExpressionStatement. got=%T", stmt)
				continue
			}

			//fmt.Printf("s is %T\n", s)
			exp, ok := s.Expression.(*ast.PrefixExpression) // expression statement should have valid identifier
			if !ok {
				t.Errorf("exp not *ast.PrefixExpression. got=%T", s.Expression)
				continue
			}
			if exp.Operator != tt.operatar {
				t.Fatalf("exp.Operator is not '%s'. got=%s",
					tt.operatar, exp.Operator)
			}
			if !testIntegerLiteral(t, exp.Right, tt.intValue) {
				return
			}
			// sendStmt := s.(*ast.SendStatement)
			// l.Debugf("%T", sendStmt)
		}

	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int) bool {
	intlit, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il is not *ast.integerLiteral. got=%T", il)
		return false
	}

	if intlit.Value != value {
		t.Errorf("intlit value not %v, got %v", value, intlit.Value)
		return false
	}

	if intlit.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("intlit.TokenLiteral not %v. got=%s", value, intlit.TokenLiteral())
		return false
	}

	return true

}

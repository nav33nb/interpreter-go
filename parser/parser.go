package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// parses some statement(node) based on what kind of node it is
func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.curToken.Type {
	case token.LET: // let [HERE] x = 5
		return p.parseLetStatement()
	case token.SEND: // let [HERE] x = 5
		return p.parseSendStatement()
	default:
		return nil, nil
	}
}

func (p *Parser) parseSendStatement() (*ast.SendStatement, error) {
	// at this point we know statemetn is "SEND *****", curToken is on SEND
	// start making SEND statment
	send := &ast.SendStatement{Token: p.curToken}

	//todo))  expression code here

	// fixme)) below skips expressions
	for p.curToken.Type != token.SEMICOLON {
		p.NextToken()
	}
	// fixme))

	return send, nil
}

func (p *Parser) parseLetStatement() (*ast.LetStatement, error) {
	// start making let statement
	// at this point we know statement is "LET *****" nothing beyond LET
	let := &ast.LetStatement{Token: p.curToken}

	// first after "let" should be identifier, fail otherwise
	if ok, err := p.peekTokenTypeIs(token.IDENT); !ok {
		fmt.Println("ok,err", ok, err)
		return nil, err
	}
	let.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.peekToken.Literal,
	}
	p.NextToken()

	// at this point we know, "LET x ****" begins with let and has valid identifier x
	// next should be assignment operator, fail otherwise
	if ok, err := p.peekTokenTypeIs(token.ASSIGN); !ok {
		return nil, err
	}
	// fmt.Println("Passing ASSIGN")
	p.NextToken()

	//todo))  expression code here

	// fixme)) below skips expressions
	for p.curToken.Type != token.SEMICOLON {
		p.NextToken()
	}
	// fixme))

	return let, nil
}

// func (p *Parser) appendError(tt token.TokenType) {
// 	p.errors = append(p.errors, fmt.Sprintf()
// }

// func (p *Parser) expectedPeekTypeIs(s string) bool {
// 	if string(p.curToken.Type) == s {
// 		p.NextToken()
// 		return true
// 	}
// 	return false
// }

func (p *Parser) currTokenTypeIs(tt token.TokenType) bool {
	return p.curToken.Type == tt
}

func (p *Parser) peekTokenTypeIs(tt token.TokenType) (bool, error) {
	if p.peekToken.Type == tt {
		return true, nil
	} else {
		return false, fmt.Errorf("parser error: AFTER %v WANT %v, HAVE %v", p.curToken, tt, p.peekToken.Type)
	}
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// read twice to fill current and peek
	p.NextToken()
	p.NextToken()
	return p
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currTokenTypeIs(token.EOF) {
		stmt, err := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
			program.Errors = append(program.Errors, err)
		}
		p.NextToken()
	}

	return program
}

package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token

	// With these maps in place, we can just check if the appropriate map
	// (infix or prefix) has a parsing function associated with curToken.Type.
	prefix map[token.TokenType]prefixParseFunc
	infix  map[token.TokenType]infixParseFunc
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFunc) {
	p.prefix[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFunc) {
	p.infix[tokenType] = fn
}

// parses some statement(node) based on what kind of node it is
func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.curToken.Type {
	case token.LET: // let [HERE] x = 5
		return p.parseLetStatement()
	case token.SEND: // let [HERE] x = 5
		return p.parseSendStatement()
	default:
		return p.parseExpressionStatement()
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

	fmt.Println("Registering prefixes for parser")
	p.prefix = make(map[token.TokenType]prefixParseFunc)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.NOT, p.parsePrefixExpression)   // like for !5
	p.registerPrefix(token.MINUS, p.parsePrefixExpression) // like for -15

	return p
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.NextToken()
	var err error
	expression.Right, err = p.parseExpression(PREFIX)
	if err != nil {
		fmt.Errorf("parser error: %v", err)
	}
	return expression
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	intlit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		fmt.Printf("could not parse %q as integer", p.curToken.Literal)
		return nil
	}

	intlit.Value = value
	return intlit
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

type (
	prefixParseFunc func() ast.Expression               //something like ++5, here there is nothing to pass as argument
	infixParseFunc  func(ast.Expression) ast.Expression //something like add(1,5) + 5, there IS A LEFT side
)

func (p *Parser) parseExpressionStatement() (*ast.ExpressionStatement, error) {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression, _ = p.parseExpression(LOWEST)

	if ok, _ := p.peekTokenTypeIs(token.SEMICOLON); ok {
		p.NextToken()
	}

	return stmt, nil
}

func (p *Parser) parseExpression(precedence int) (ast.Expression, error) {
	prefix := p.prefix[p.curToken.Type]
	if prefix == nil {
		return nil, p.noPrefix(p.curToken.Type)
	}
	leftExpression := prefix()
	return leftExpression, nil
}

func (p *Parser) noPrefix(tokenType token.TokenType) error {
	return fmt.Errorf("parser error: unknown prefix type %v", p.curToken.Type)
}

const (
	_ int = iota
	LOWEST
	EQUALITY
	LESSMORE
	PLUS
	MULTIPLY
	PREFIX
	CALL
)

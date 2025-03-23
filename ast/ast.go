package ast

import (
	"bytes"
	"monkey/token"
)

// Program
// ├── Statement 1 (LetStatement)	 <------ Node#1
// │   ├── Token: LET
// │   ├── Name: Identifier ("x")		<-- *identifier
// │   └── expression: Literal (5)			<-- expression
// │
// ├── Statement 2 (LetStatement)	 <------ Node#2
// │   ├── Token: LET
// │   ├── Name: Identifier ("y")		<-- *identifier
// │   └── expression: Literal (6)			<-- expression
// │
// └── Statement 3 (AssignStatement) <------ Node#3
//	   ├── Token: ASSIGN
//	   ├── Name: Identifier ("z")		<-- *identifier
//	   └── expression: Add					<-- expression (composite)
//		   ├── Left: Identifier ("x")
//		   └── Right: Identifier ("y")

type Node interface {
	TokenLiteral() string
	ToString() string
}

// Program = Sequence of Statements, This will be the top-level
//
//	A program is a series of statements.
type Program struct {
	Statements []Statement
	Errors     []error
}

// Statement = representation of each Node
//
//	Each statement is a node in the AST and is some kind of keyword operation (like let).
//	The statement may contain an Identifier and an Expression.
type Statement interface {
	Node
	statementNode() // dummy method
}

// Expression = Recursive Node Structure, THE RIGHT SIDE
//
//	An expression can be a simple token (like a literal or identifier) or a recursive combination of expressions (like binary operations). IDEA IS, expression can recursively be other expression or end up in some identifier. Expressions like x * (y+z) have an operation node with sub-nodes (the operands).
//
// so identifier should implement expression, since an identifier is just "END OF AN EXPRESSION"
type Expression interface {
	Node
	expressionNode() // dummy method
}

// Identifier
// Tokens, THE LEFT SIDE
// identifier should implement expression, since an identifier is just "KIND/END OF AN EXPRESSION"
type Identifier struct {
	Token token.Token
	Value string
}

func (id *Identifier) expressionNode() {}
func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}
func (id *Identifier) ToString() string {
	return id.Value
}

// LetStatement Structural representation of let statement
// Since statement=node, it must implement node methods
type LetStatement struct {
	Token token.Token // LET
	Name  *Identifier // x
	Value Expression  // 5
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) ToString() string {
	//fmt.Println(ls, ls.Name, ls.Value)
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.ToString())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.ToString())
	}
	out.WriteString(";")
	return out.String()
}

// SendStatement Structural representation of send statement
// Since statement=node, it must implement node methods
type SendStatement struct {
	Token token.Token // SEND
	Value Expression  // 5
}

func (ss *SendStatement) statementNode()       {}
func (ss *SendStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SendStatement) ToString() string {
	var out bytes.Buffer
	out.WriteString(ss.TokenLiteral() + " ")
	if ss.Value != nil {
		out.WriteString(ss.Value.ToString())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement these are the statements without any LEFT, entire statement is an expression
// ExpressionStatement is like a wrapper, which contains only single expression
type ExpressionStatement struct {
	Token      token.Token // first token in the expression, like 55*10 has 55
	Expression Expression  // 5
}

func (es *ExpressionStatement) expressionNode() {
	//TODO implement me
	panic("implement me")
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) ToString() string {
	var out bytes.Buffer
	out.WriteString(es.TokenLiteral() + " ")
	if es.Expression != nil {
		out.WriteString(es.Expression.ToString())
	}
	out.WriteString(";")
	return out.String()
}

// IntegerLiteral Literal representation of integers
type IntegerLiteral struct {
	Token token.Token // token
	Value int         // integer value
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) ToString() string     { return il.Token.Literal }

// func (p *Program) TokenLiteral() string {
// 	if len(p.Statements) > 0 {
// 		return p.Statements[0].TokenLiteral()
// 	} else {
// 		return ""
// 	}
// }

func (p *Program) ToString() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.ToString())
	}

	return out.String()
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) ToString() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.ToString())
	out.WriteString(")")
	return out.String()
}

package ast

import "monkey/token"

// Program
// ├── Statement 1 (LetStatement)	 <------ Node#1
// │   ├── Token: LET
// │   ├── Name: Identifier ("x")		<-- *identifier
// │   └── Value: Literal (5)			<-- expression
// │
// ├── Statement 2 (LetStatement)	 <------ Node#2
// │   ├── Token: LET
// │   ├── Name: Identifier ("y")		<-- *identifier
// │   └── Value: Literal (6)			<-- expression
// │
// └── Statement 3 (AssignStatement) <------ Node#3
//	   ├── Token: ASSIGN
//	   ├── Name: Identifier ("z")		<-- *identifier
//	   └── Value: Add					<-- expression (composite)
//		   ├── Left: Identifier ("x")
//		   └── Right: Identifier ("y")

type Node interface {
	TokenLiteral() string
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
// so identifier should implment expression, since an identifier is is just "END OF AN EXPRESSION"
type Expression interface {
	Node
	expressionNode() // dummy method
}

// Identifiers = Tokens, THE LEFT SIDE
//
// identifier should implment expression, since an identifier is is just "KIND/END OF AN EXPRESSION"
type Identifier struct {
	Token token.Token
	Value string
}

func (id *Identifier) expressionNode() {}
func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}

// LET STATEMENT: Structural represenation of let statement
// Since statement=node, it must implement node methods
type LetStatement struct {
	Token token.Token // LET
	Name  *Identifier // x
	Value Expression  // 5
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// func (p *Program) TokenLiteral() string {
// 	if len(p.Statements) > 0 {
// 		return p.Statements[0].TokenLiteral()
// 	} else {
// 		return ""
// 	}
// }

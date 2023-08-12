// Builds the nodes for the Abstract Syntax Tree (AST).

package ast

import "github.com/GStechschulte/go-interpreter/src/monkey/token"

// Node is the basic interface for all nodes in the AST
type Node interface {
	TokenLiteral() string
}

// Statement is a node that doesn't produce a value
type Statement interface {
	Node
	statementNode()
}

// Expression is a node that produces a value
type Expression interface {
	Node
	expressionNode()
}

// Monkey Program is a series of Statements
type Program struct {
	Statements []Statement
}

// root node of every AST our parser produces
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement contains the three fields we need for variable binding
type LetStatement struct {
	Token token.Token
	Name  *Identifier // variable name
	Value Expression  // "produces" the value
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

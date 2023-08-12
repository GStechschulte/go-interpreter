// Parses the tokens from the lexer and constructs the Abstract Syntax
// Tree (AST) for the Monkey program.

package parser

import (
	"github.com/GStechschulte/go-interpreter/src/monkey/ast"
	"github.com/GStechschulte/go-interpreter/src/monkey/lexer"
	"github.com/GStechschulte/go-interpreter/src/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	// point to current and next token to decide what to do "next"
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read two tokens so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{} // root node of AST
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			// if the statement is not nil, append the elements to the
			// slice of statements in the program
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() *ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skipping the expressions until we encounter a semicolon
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	// check the type of the next token
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	// if the next token is of the expected type, advance the tokens
	// and return true
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

package main

import (
	"fmt"

	"github.com/GStechschulte/go-interpreter/src/monkey/lexer"
	"github.com/GStechschulte/go-interpreter/src/monkey/token"
)

func main() {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Printf("%d, expected: %s, got: %s \n", i, tt.expectedLiteral, tok.Literal)
	}
}

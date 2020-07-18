package lexer

import (
	"github.com/Asalle/monkey-interpreter/token"
	"testing"
)

func Test_NextToken(t *testing.T) {
	input := "=+{}();,"

	cases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
	}

	l := New(input)

	for _, tt := range cases {
		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Errorf("Wrong token type: expected: %v, got: %v", tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Errorf("Wrong token literal: expected: %v, got: %v", tt.expectedLiteral, token.Literal)
		}
	}
}

func Test_NextToken_Monkey(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
}

let result = add(five, ten);
`

	cases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for _, tt := range cases {
		token := l.NextToken()
		if token.Type != tt.expectedType {
			t.Errorf("Wrong token type: expected: '%v', got: '%v'", tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Errorf("Wrong token literal: expected: '%v', got: '%v'", tt.expectedLiteral, token.Literal)
		}
	}
}

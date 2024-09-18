package lexer

import (
	"schemego/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `(define (square x) (* x x))`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPARENT, "("},
		{token.DEFINE, "define"},
		{token.LPARENT, "("},
		{token.IDENT, "square"},
		{token.IDENT, "x"},
		{token.RPARENT, ")"},
		{token.LPARENT, "("},
		{token.ASTERISK, "*"},
		{token.IDENT, "x"},
		{token.IDENT, "x"},
		{token.RPARENT, ")"},
		{token.RPARENT, ")"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected %q, got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

package lexer

import "schemego/token"

type Lexer struct {
	input        string
	ch           byte
	position     int
	readPosition int
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	var tok token.Token

	switch l.ch {
	case '(':
		tok = NewToken(token.LPARENT, l.ch)
	case ')':
		tok = NewToken(token.RPARENT, l.ch)
	case '*':
		tok = NewToken(token.ASTERISK, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isNumber(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = NewToken(token.ILLEGAL, l.ch)
		}
	}
  l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func NewToken(ttype token.TokenType, ch byte) token.Token {
	return token.Token{Type: ttype, Literal: string(ch)}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

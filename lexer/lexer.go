package lexer

import (
	"github.com/shozawa/monkey2/token"
)

type Lexer struct {
	input    string
	position int
	ch       byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	// FIXME: input が空文字のときも考慮すべき
	l.ch = input[0]
	return l
}

func (l *Lexer) NextToken() (tok token.Token) {
	l.skipWhitespace()
	if isDigit(l.ch) {
		return l.readDigit()
	}
	switch l.ch {
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: "+"}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: ";"}
	case 0:
		return token.Token{Type: token.EOF, Literal: ""}
	default:
		return token.Token{}
	}
	l.readChar()
	return
}

func (l *Lexer) readChar() {
	l.position += 1
	if l.position >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
	}
}

func (l *Lexer) readDigit() token.Token {
	var result []byte
	tok := token.Token{Type: token.INT}
	for isDigit(l.ch) {
		result = append(result, l.ch)
		l.readChar()
	}
	tok.Literal = string(result)
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

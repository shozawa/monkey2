package lexer

import (
	"testing"

	"github.com/shozawa/monkey2/token"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		input string
		want  []token.Token
	}{
		{"5;", []token.Token{
			token.Token{Type: token.INT, Literal: "5"},
			token.Token{Type: token.SEMICOLON, Literal: ";"},
			token.Token{Type: token.EOF, Literal: ""},
		}},
		{"42;", []token.Token{
			token.Token{Type: token.INT, Literal: "42"},
			token.Token{Type: token.SEMICOLON, Literal: ";"},
			token.Token{Type: token.EOF, Literal: ""},
		}},
		{"10 + 25 * 3;", []token.Token{
			token.Token{Type: token.INT, Literal: "10"},
			token.Token{Type: token.PLUS, Literal: "+"},
			token.Token{Type: token.INT, Literal: "25"},
			token.Token{Type: token.ASTERISK, Literal: "*"},
			token.Token{Type: token.INT, Literal: "3"},
			token.Token{Type: token.SEMICOLON, Literal: ";"},
			token.Token{Type: token.EOF, Literal: ""},
		}},
	}

	for _, test := range tests {
		l := New(test.input)
		for _, tok := range test.want {
			got := l.NextToken()
			if got.Type != tok.Type {
				t.Errorf("token.Type is not %q. got=%q", tok.Type, got.Type)
			}
			if got.Literal != tok.Literal {
				t.Errorf("token.Literal is not %q. got=%q", tok.Literal, got.Literal)
			}
		}
	}
}

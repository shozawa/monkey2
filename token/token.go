package token

type TokenType string

const (
	INT       = "INT"
	PLUS      = "+"
	SEMICOLON = ";"
	EOF       = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}

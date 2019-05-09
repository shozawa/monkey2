package token

type TokenType string

const (
	IDNT      = "IDNT"
	INT       = "INT"
	PLUS      = "+"
	ASTERISK  = "*"
	SEMICOLON = ";"
	EOF       = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}

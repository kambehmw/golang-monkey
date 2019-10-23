package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF 	= "EOF"

	IDENT	= "IDENT"
	INT		= "INT"

	ASSIGN	= "="
	PLUS	= "+"
	MINUS   = "-"
	BANG    = "!"
	ASTERISK  = "*"
	SLASH   = "/"

	LT = "<"
	GT = ">"

	COMMA	= ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"

	EQ = "=="
	NOT_EQ = "!="
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIndent(idnet string) TokenType {
	if tok, ok := keywords[idnet]; ok {
		return tok
	}
	return IDENT
}
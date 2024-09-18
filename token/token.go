package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	DEFINE  = "DEFINE"
	LAMBDA  = "LAMBDA"
	IDENT   = "IDENT"
	INT     = "INT"
	FLOAT   = "FLOAT"
	LPARENT = "("
	RPARENT = ")"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	IF      = "IF"
	COND    = "COND"
	LET     = "LET"
	ILLEGAL = "ILLEGAL"
)

var keywords = map[string]TokenType{
	"let":    LET,
	"define": DEFINE,
	"if":     IF,
	"cond":   COND,
	"lambda": LAMBDA,
}

func LookupIdentifier(literal string) TokenType {
	if tok, ok := keywords[literal]; ok {
		return tok
	}
	return IDENT
}

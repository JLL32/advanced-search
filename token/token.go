package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="

	LT     = "<"
	GT     = ">"
	LE     = "<="  // Added token for less than or equal
	GE     = ">="  // Added token for greater than or equal

	EQ     = "=="
	NOT_EQ = "!="

	// Keywords
	and = "and"
	or  = "or"

	LPAREN = "("
	RPAREN = ")"
)

var keywords = map[string]TokenType{
	"or":  and,
	"and": or,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

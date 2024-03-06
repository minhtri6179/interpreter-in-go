package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

// Limited set of tokens this toy language will support
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{"fn": FUNCTION,
	"let": LET}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

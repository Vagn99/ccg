package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	NEWLINE = "NEWLINE"

	// Whitespace
	WHITESPACE = "WHITESPACE"

	// Comments
	COMMENT = "COMMENT"

	// Pattern components
	NEGATION       = "NEGATION"       // !
	PATH_SEPARATOR = "PATH_SEPARATOR" // /
	WILDCARD       = "WILDCARD"       // *, **, ?

	// Identifiers
	CHARACTER = "CHARACTER" // a-z, A-Z, 0-9, ., -, _
)

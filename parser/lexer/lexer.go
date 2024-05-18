package lexer

import "ccg/parser/token"

type Lexer struct {
	input    string
	position int  // current position in input (points to current char)
	readPos  int  // current reading position in input (after current char)
	ch       byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.position = l.readPos
	l.readPos++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '!':
		tok = newToken(token.NEGATION, l.ch)
	case '/':
		tok = newToken(token.PATH_SEPARATOR, l.ch)
	case '*':
		if l.peekChar() == '*' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.WILDCARD, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.WILDCARD, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.ReadChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

func (l *Lexer) readChar() byte {
	ch := l.ch
	l.ReadChar()
	return ch
}

// lexer package is to tokenize the input to tokens.
package lexer

import (
	"github.com/PetraZ/monkey/token"
)

// What is a lexer?
// A lexer or a tokenizer - takes in a string as the input(source code) and emits tokens that
// represents the source code.
type Lexer struct {
	input string
	// current position that points to char
	pos int
	// current char
	char byte
}

func New(input string) *Lexer {
	if len(input) == 0 {
		return nil
	}
	return &Lexer{
		input: input,
		pos:   0,
		char:  input[0],
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	// operators
	case '=':
		if l.peakChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			tok = token.Token{Type: token.ASSIGN, Literal: "="}
		}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: "+"}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: "-"}
	case '!':
		if l.peakChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			tok = token.Token{Type: token.BANG, Literal: "!"}
		}
	case '*':
		tok = token.Token{Type: token.ASTERISK, Literal: "*"}
	case '/':
		tok = token.Token{Type: token.SLASH, Literal: "/"}
	case '<':
		tok = token.Token{Type: token.LT, Literal: "<"}
	case '>':
		tok = token.Token{Type: token.GT, Literal: ">"}
	// delimiters
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: "("}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: ")"}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: "{"}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: "}"}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: ","}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: ";"}

	default:
		if l.isLetter() {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if l.isDigit() {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok.Type = token.ILLEGAL
			tok.Literal = string(l.char)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	start := l.pos
	for l.isDigit() {
		l.readChar()
	}
	return l.input[start:l.pos]
}

func (l *Lexer) readIdentifier() string {
	start := l.pos
	for l.isLetter() {
		l.readChar()
	}
	return l.input[start:l.pos]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) isLetter() bool {
	return (l.char >= 'a' && l.char <= 'z') || (l.char >= 'A' && l.char <= 'Z') || l.char == '_'
}

func (l *Lexer) isDigit() bool {
	return (l.char >= '0') && (l.char <= '9')
}

func (l *Lexer) readChar() {
	// move pos to next char
	l.pos += 1
	if l.pos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.pos]
	}
}

func (l *Lexer) peakChar() byte {
	if (l.pos + 1) >= len(l.input) {
		return 0
	}
	return l.input[l.pos+1]
}

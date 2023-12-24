package lexer

import (
	"strings"
	"unicode"
)

type Lexer struct {
	char   []rune
	Length int
	cursor int
}

func NewLexer(str string) Lexer {
	char :=  []rune(str)
	return Lexer{
		char:  char,
		Length: len(char),
		cursor: 0,
	}
}

func (l *Lexer) trim() {
	char := l.char[l.cursor]
	
	for((!unicode.IsLetter(char) && !unicode.IsDigit(char))){
		l.cursor++
		if(l.cursor >= l.Length - 1){
			break
		}
		char = l.char[l.cursor]
	}
}

func (l *Lexer) chop() {
	char := l.char[l.cursor]
	for((!unicode.IsSpace(char) && !unicode.IsPunct(char))){
		
		l.cursor++
		if( l.cursor >= l.Length - 1){
			break
		}
		char = l.char[l.cursor]
	}
}

func (l *Lexer) Get() (string, bool) {
	if l.cursor >= l.Length - 1 {
		return "", true
	}
	l.trim()
	start := l.cursor
	l.chop()
	chars := strings.ToLower(string(l.char[start:l.cursor]))
	return chars, false
}
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/PetraZ/monkey/lexer"
	"github.com/PetraZ/monkey/token"
)

// The main job of a repl is, given an input(string) to interactively produce tokens.
func main() {
	// os.Stdin reads what you typed in cmd line
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		s := scanner.Text()
		l := lexer.New(s)
		for {
			t := l.NextToken()
			fmt.Println(t.Type, t.Literal)
			if t.Type == token.EOF {
				break
			}
		}
		fmt.Println("")
		fmt.Println("Finished tokenization, feel free to enter next...")
	}
}

// example:
// go run repl.go
// let add = fn(x, y) {x +y;};

// will output
// LET let
// IDENT add
// = =
// FUNCTION fn
// ( (
// IDENT x
// , ,
// IDENT y
// ) )
// { {
// IDENT x
// + +
// IDENT y
// ; ;
// } }
// ; ;
// EOF

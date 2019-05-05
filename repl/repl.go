package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/shozawa/monkey2/evaluator"
	"github.com/shozawa/monkey2/lexer"
	"github.com/shozawa/monkey2/parser"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print("monkey> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		fmt.Print("=> ")

		l := lexer.New(scanner.Text())
		p := parser.New(l)
		program := p.Parse()
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

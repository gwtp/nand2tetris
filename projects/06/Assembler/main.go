// Binary main reads as input a text file Prog.asm, containing Hack assembly program,
// and produces as output a text file named Prog.hack, containing the translated Hack
// machine code.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gwtp/nand2tetris/projects/06/Assembler/code"
	"github.com/gwtp/nand2tetris/projects/06/Assembler/parser"
)

var file = flag.String("file", "", "Path to the assembly file, containing hack assembly program.")

func translate(in io.Reader, out io.Writer) error {
	p := parser.New(in)
	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case parser.LCommand:
		case parser.ACommand:
			i, err := strconv.Atoi(p.Symbol())
			if err != nil {
				return err
			}
			out.Write([]byte(fmt.Sprintf("0%s\n", fmt.Sprintf("%016b", i))[1:]))
		case parser.CCommand:
			comp, dest, jump := code.Comp(p.Comp()), code.Dest(p.Dest()), code.Jump(p.Jump())
			out.Write([]byte(fmt.Sprintf("111%s%s%s\n", comp, dest, jump)))
		}
	}
	if err := p.Error(); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()

	in, err := os.Open(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open error: %v\n", err)
		os.Exit(1)
	}
	defer in.Close()

	outFile := fmt.Sprintf("%s.hack", strings.Split(in.Name(), ".")[0])
	out, err := os.Create(outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create error: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	if err := translate(in, out); err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(1)
	}

	return
}

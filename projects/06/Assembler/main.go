// Binary main reads as input a text file Prog.asm, containing Hack assembly program,
// and produces as output a text file named Prog.hack, containing the translated Hack
// machine code.
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gwtp/nand2tetris/projects/06/Assembler/code"
	"github.com/gwtp/nand2tetris/projects/06/Assembler/parser"
)

var file = flag.String("file", "", "Path to the assembly file, containing hack assembly program.")

func parse(in, out *os.File) error {
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
			out.WriteString(fmt.Sprintf("0%s\n", fmt.Sprintf("%016b", i)[1:]))
		case parser.CCommand:
			out.WriteString(fmt.Sprintf("111%s%s%s\n", code.Comp(p.Comp()), code.Dest(p.Dest()), code.Jump(p.Jump())))
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
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer in.Close()

	outFile := fmt.Sprintf("%s.hack", strings.Split(in.Name(), ".")[0])
	out, err := os.Create(outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	if err := parse(in, out); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	return
}

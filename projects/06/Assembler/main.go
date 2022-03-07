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
	"github.com/gwtp/nand2tetris/projects/06/Assembler/symbol"
)

var (
	file = flag.String("file", "", "Path to the assembly file, containing hack assembly program.")
)

func labels(in io.Reader) (*symbol.Symbol, error) {
	var lineCount int

	s := symbol.New(symbol.BuiltIn)
	p := parser.New(in)
	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case parser.ACommand, parser.CCommand:
			lineCount++
		case parser.LCommand:
			s.AddEntry(p.Symbol(), lineCount)
		}
		if err := p.Error(); err != nil {
			return s, err
		}
	}
	return s, nil
}

func translate(in io.Reader, out io.Writer, s *symbol.Symbol) error {
	memStart := 16

	p := parser.New(in)
	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case parser.ACommand:
			var addr int
			i, err := strconv.Atoi(p.Symbol())
			switch {
			case err == nil:
				addr = i
			case err != nil:
				if !s.Contains(p.Symbol()) {
					s.AddEntry(p.Symbol(), memStart)
					memStart++
				}
				addr = s.GetAddress(p.Symbol())
			}
			binAddr := fmt.Sprintf("%016b", addr)[1:]
			aTempl := "0%s\n"
			out.Write([]byte(fmt.Sprintf(aTempl, binAddr)))
		case parser.CCommand:
			cTempl := "111%s%s%s\n"
			binComp, binDest, binJump := code.Comp(p.Comp()), code.Dest(p.Dest()), code.Jump(p.Jump())
			out.Write([]byte(fmt.Sprintf(cTempl, binComp, binDest, binJump)))
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

	s, err := labels(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "labels error: %v\n", err)
	}

	if _, err := in.Seek(0, 0); err != nil {
		fmt.Fprintf(os.Stderr, "seek error: %v\n", err)
	}

	if err := translate(in, out, s); err != nil {
		fmt.Fprintf(os.Stderr, "translate error: %v\n", err)
		os.Exit(1)
	}

	return
}

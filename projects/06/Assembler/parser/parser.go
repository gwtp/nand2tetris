// Package parser reads an assembly language command from file and provides convenient access to the commands components.
package parser

import (
	"bufio"
	"io"
	"strings"
)

type CommandType uint8

const (
	UnknownCommand CommandType = iota
	ACommand
	CCommand
	LCommand
)

type Parser struct {
	scanner      *bufio.Scanner
	moreCommands bool
	err          error

	commandType CommandType
	comp        string
	dest        string
	jump        string
	symbol      string
}

func New(f io.Reader) *Parser {
	return &Parser{scanner: bufio.NewScanner(f)}
}

func (p *Parser) Error() error {
	return p.err
}

// HasMoreCommands checks if there are any more commands in the input.
func (p *Parser) HasMoreCommands() bool {
	s := p.scanner.Scan()

	if err := p.scanner.Err(); err != nil {
		p.err = err
		return false
	}

	return s
}

// reset sets parser variables back to zero value.
func (p *Parser) reset() {
	p.commandType = UnknownCommand
	p.comp = ""
	p.dest = ""
	p.jump = ""
	p.symbol = ""
}

// Advance reads the next command from the input. Should only be called if HasMoreCommands() is true.
func (p *Parser) Advance() {
	p.reset()
	switch line := strings.TrimSpace(p.scanner.Text()); {
	// Comment
	case strings.HasPrefix(line, "//"):
		p.commandType = UnknownCommand
	// Label
	case strings.HasPrefix(line, "("):
		p.commandType = LCommand
		p.symbol = line[1 : len(line)-1]
	// A Command
	case strings.HasPrefix(line, "@"):
		p.commandType = ACommand
		p.symbol = line[1:]
	// C Command - Comp
	case strings.Contains(line, "="):
		p.commandType = CCommand
		parts := strings.Split(line, "=")
		p.dest, p.comp = parts[0], parts[1]
	// C Command - Goto (Jump)
	case strings.Contains(line, ";"):
		p.commandType = CCommand
		parts := strings.Split(line, ";")
		p.comp, p.jump = parts[0], parts[1]
	}
}

// CommandType returns the type of the current command.
func (p *Parser) CommandType() CommandType {
	return p.commandType
}

// Comp returns the comp mnemonic. Should only be called when CommandType() is CCommand.
func (p *Parser) Comp() string {
	return p.comp
}

// Dest return the dest mnemonic. Should only be called when CommandType() is CCommand.
func (p *Parser) Dest() string {
	return p.dest
}

// Jump returns the jump mnemonic. Should only be called when CommandType() is CCommand.
func (p *Parser) Jump() string {
	return p.jump
}

// Symbol returns the symbol or decimal of the current command. Should only be called when command type is ACommand or LCommand.
func (p *Parser) Symbol() string {
	return p.symbol
}

// Package parser reads an assembly language command from file and provides convenient access to the commands components.
package parser

import (
	"bufio"
	"os"
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

	commandType CommandType
	comp        string
	dest        string
	jump        string
	symbol      string
}

func New(path string) (*Parser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return &Parser{scanner: bufio.NewScanner(f)}, nil
}

// HasMoreCommands checks if there are any more commands in the input.
func (p *Parser) HasMoreCommands() (bool, error) {
	s := p.scanner.Scan()

	if err := p.scanner.Err(); err != nil {
		return false, err
	}

	return s, nil
}

// Advance reads the next command from the input. Should only be called if HasMoreCommands() is true.
func (p *Parser) Advance() {
	line := p.scanner.Text()
	switch {
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

// Symbol returns the symbol or decimal of the current command. Should only be called when command type is ACommand or LCommand.
func (p *Parser) Symbol() string {
	return p.symbol
}

// Dest return the dest mnemonic. Should only be called when CommandType() is CCommand.
func (p *Parser) Dest() string {
	return p.dest
}

// Comp returns the comp mnemonic. Should only be called when CommandType() is CCommand.
func (p *Parser) Comp() string {
	return p.comp
}

// Jump returns the jump mnemonic. Should only be called when CommandType() is CCommand.
func (p *Parser) Jump() string {
	return p.jump
}

// Package parser reads an assembly language command from file and provides convenient access to the commands components.
package parser

type CommandType uint8

const (
	ACommand CommandType = iota
	CCommand
	LCommand
)

type Parser struct{}

func New(path string) *Parser {
	return &Parser{}
}

// HasMoreCommands checks if there are any more commands in the input.
func (p *Parser) HasMoreCommands() bool {
	return false
}

// Advance reads the next command from the input. Should only be called if HasMoreCommands() is true.
func (p *Parser) Advance() {
	return
}

// CommandType returns the type of the current command.
func (p *Parser) CommandType() CommandType {
	return nil
}

// Symbol returns the symbol or decimal of the current command. Should only be called when command type is ACommand or LCommand.
func (p *Parser) Symbol() string {
	return ""
}

// Dest return the dest mnemonic. Should only be called when CommandType() is CCommand.
func (p *Parser) Dest() string {
	return ""
}

// Comp returns the comp mnemonic. Should only be called when CommandType() is CCommand.
func (p *Parser) Comp() string {
	return ""
}

// Jump returns the jump mnemonic. Should only be called when CommandType() is CCommand.
func (p *Parser) Jump() string {
	return ""
}

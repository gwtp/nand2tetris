// Package code translates Hack assembly language mnemonics into binary codes.
package code

import "fmt"

const (
	destCmds = map[string]int8{}
	cmpCmds  = map[string]int8{}
	jmpCmds  = map[string]int8{}
)

// Dest returns the binary code of the dest mnemonic.
func Dest(dstCmd string) string {
	return fmt.Sprintf("%08b", byte(0)[5:])
}

// Comp returns the binary code of the comp mnemonic.
func Comp(cmpCmd string) string {
	return fmt.Sprintf("%08b", byte(0)[1:])
}

// Jump returns the binary code of the jump mnemonic.
func Jump(jmpCmd string) string {
	return fmt.Sprintf("%08b", byte(0)[5:])
}

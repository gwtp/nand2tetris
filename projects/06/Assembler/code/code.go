// Package code translates Hack assembly language mnemonics into binary codes.
package code

import (
	"fmt"
	"strings"
)

var (
	cmpCmds = map[string]int8{
		"0":   42,
		"1":   63,
		"-1":  58,
		"D":   12,
		"A":   48,
		"M":   48,
		"!D":  13,
		"!A":  49,
		"!M":  49,
		"-D":  15,
		"-A":  51,
		"-M":  51,
		"D+1": 31,
		"A+1": 55,
		"M+1": 55,
		"D-1": 14,
		"A-1": 50,
		"M-1": 50,
		"D+A": 2,
		"D+M": 2,
		"D-A": 19,
		"D-M": 19,
		"A-D": 7,
		"M-D": 7,
		"D&A": 0,
		"D&M": 0,
		"D|A": 21,
		"D|M": 21,
	}
	dstCmds = map[string]int8{
		"M":   1,
		"D":   2,
		"MD":  3,
		"A":   4,
		"AM":  5,
		"AD":  6,
		"AMD": 7,
	}
	jmpCmds = map[string]int8{
		"JGT": 1,
		"JEQ": 2,
		"JGE": 3,
		"JLT": 4,
		"JNE": 5,
		"JLE": 6,
		"JMP": 7,
	}
)

// Comp returns the binary code of the comp mnemonic.
func Comp(cmpCmd string) string {
	var aBit int8
	if strings.Contains(cmpCmd, "M") {
		aBit = 64
	}
	return fmt.Sprintf("%08b", byte(cmpCmds[cmpCmd]+aBit))[1:]
}

// Dest returns the binary code of the dest mnemonic.
func Dest(dstCmd string) string {
	var cmdDecimal int8
	// Defaults to 0 (null) if dstCmd is non-existent.
	if d, ok := dstCmds[dstCmd]; ok {
		cmdDecimal = d
	}

	return fmt.Sprintf("%08b", byte(cmdDecimal))[5:]
}

// Jump returns the binary code of the jump mnemonic.
func Jump(jmpCmd string) string {
	var cmdDecimal int8
	// Defaults to 0 (null) if jmpCmd is non-existent.
	if d, ok := jmpCmds[jmpCmd]; ok {
		cmdDecimal = d
	}

	return fmt.Sprintf("%08b", byte(cmdDecimal))[5:]
}

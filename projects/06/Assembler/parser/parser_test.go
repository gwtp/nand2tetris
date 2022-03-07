package parser

import (
	"bufio"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParserSymbol(t *testing.T) {

	tests := []struct {
		name            string
		input           string
		hasMoreCommands bool
		cType           CommandType
		wantSymbol      string
	}{
		{
			name:            "Label",
			input:           "(LOOP)",
			hasMoreCommands: true,
			cType:           LCommand,
			wantSymbol:      "LOOP",
		},
		{
			name:            "A Command variable",
			input:           "@i",
			hasMoreCommands: true,
			cType:           ACommand,
			wantSymbol:      "i",
		},
		{
			name:            "A Command decimal",
			input:           "@100",
			hasMoreCommands: true,
			cType:           ACommand,
			wantSymbol:      "100",
		},
		{
			name:            "A Command trailing comment",
			input:           "@100 // foobar",
			hasMoreCommands: true,
			cType:           ACommand,
			wantSymbol:      "100",
		},
		{
			name:            "Comment",
			input:           "// @foobar",
			hasMoreCommands: true,
			cType:           UnknownCommand,
			wantSymbol:      "",
		},
		{
			name:            "empty input",
			input:           "",
			hasMoreCommands: false,
			cType:           UnknownCommand,
			wantSymbol:      "",
		},
	}

	for _, tc := range tests {
		p := Parser{scanner: bufio.NewScanner(strings.NewReader(tc.input))}
		got := p.HasMoreCommands()
		if diff := cmp.Diff(tc.hasMoreCommands, got); diff != "" {
			t.Errorf("p.HasMoreCommands(%s) mismatch (-want +got):\n%s", tc.name, diff)
			continue
		}
		if err := p.Error(); err != nil {
			t.Fatalf("p.Error(%s): %v", tc.name, err)
		}

		// No commands to process
		if got == false {
			continue
		}

		// Parse the next line.
		p.Advance()

		// Validate command type.
		if diff := cmp.Diff(tc.cType, p.CommandType()); diff != "" {
			t.Errorf("p.CommandType(%s) mismatch (-want +got):\n%s", tc.name, diff)
		}

		// Validate symbol output.
		if diff := cmp.Diff(tc.wantSymbol, p.Symbol()); diff != "" {
			t.Errorf("p.Symbol(%s) mismatch (-want +got):\n%s", tc.name, diff)
		}
	}
}

func TestParserCCommand(t *testing.T) {

	tests := []struct {
		name            string
		input           string
		hasMoreCommands bool
		cType           CommandType
		wantComp        string
		wantDest        string
		wantJump        string
	}{
		{
			name:            "C Instruction",
			input:           "M=D+M",
			hasMoreCommands: true,
			cType:           CCommand,
			wantComp:        "D+M",
			wantDest:        "M",
			wantJump:        "",
		},
		{
			name:            "C Instruction Goto",
			input:           "D;JGT",
			hasMoreCommands: true,
			cType:           CCommand,
			wantComp:        "D",
			wantDest:        "",
			wantJump:        "JGT",
		},
		{
			name:            "C Instruction Goto trailing comment",
			input:           "0;JMP // foobar",
			hasMoreCommands: true,
			cType:           CCommand,
			wantComp:        "0",
			wantDest:        "",
			wantJump:        "JMP",
		},
		{
			name:            "Comment 1",
			input:           "// M=D+M",
			hasMoreCommands: true,
			cType:           UnknownCommand,
			wantComp:        "",
			wantDest:        "",
			wantJump:        "",
		},
		{
			name:            "Comment 2",
			input:           "// 0;JMP",
			hasMoreCommands: true,
			cType:           UnknownCommand,
			wantComp:        "",
			wantDest:        "",
			wantJump:        "",
		},
		{
			name:            "Empty input",
			input:           "",
			hasMoreCommands: false,
			cType:           UnknownCommand,
			wantComp:        "",
			wantDest:        "",
			wantJump:        "",
		},
	}

	for _, tc := range tests {
		p := Parser{scanner: bufio.NewScanner(strings.NewReader(tc.input))}
		got := p.HasMoreCommands()
		if diff := cmp.Diff(tc.hasMoreCommands, got); diff != "" {
			t.Errorf("p.HasMoreCommands(%s) mismatch (-want +got):\n%s", tc.name, diff)
			continue
		}
		if err := p.Error(); err != nil {
			t.Fatalf("p.Error(%s): %v", tc.name, err)
		}

		// No commands to process
		if got == false {
			continue
		}

		// Parse the next line.
		p.Advance()

		// Validate command type.
		if diff := cmp.Diff(tc.cType, p.CommandType()); diff != "" {
			t.Errorf("p.CommandType(%s) mismatch (-want +got):\n%s", tc.name, diff)
		}

		// Validate comp output.
		if diff := cmp.Diff(tc.wantComp, p.Comp()); diff != "" {
			t.Errorf("p.Comp(%s) mismatch (-want +got):\n%s", tc.name, diff)
		}
		// Validate dest output.
		if diff := cmp.Diff(tc.wantDest, p.Dest()); diff != "" {
			t.Errorf("p.Dest(%s) mismatch (-want +got):\n%s", tc.name, diff)
		}
		// Validate jump output.
		if diff := cmp.Diff(tc.wantJump, p.Jump()); diff != "" {
			t.Errorf("p.Jump(%s) mismatch (-want +got):\n%s", tc.name, diff)
		}
	}
}

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gwtp/nand2tetris/projects/06/Assembler/symbol"
)

func mapUnion(t *testing.T, m1, m2 map[string]int) map[string]int {
	for k, v := range m1 {
		if _, ok := m2[k]; ok {
			t.Fatalf("mapUnion(): duplicate key: %v", k)
		}
		m2[k] = v
	}

	return m2
}

func TestLabels(t *testing.T) {
	prog := `@R0
	D=M
	@R1
	D=D-M
	@OUTPUT_FIRST
	D;JGT
	@R1
	D=M
	@OUTPUT_D
	0;JMP
 (OUTPUT_FIRST)
	@R0             
	D=M
 (OUTPUT_D)
	@R2
	M=D
 (INFINITE_LOOP)
	@INFINITE_LOOP
	0;JMP`

	tests := []struct {
		name    string
		in      string
		want    *symbol.Symbol
		wantErr bool
	}{
		{
			name: "Prog",
			in:   prog,
			want: symbol.New(mapUnion(t, symbol.BuiltIn, map[string]int{"OUTPUT_FIRST": 10, "OUTPUT_D": 12, "INFINITE_LOOP": 14})),
		},
	}

	for _, tc := range tests {
		s, err := labels(strings.NewReader(tc.in))
		if want, got := tc.wantErr, (err != nil); want != got {
			t.Errorf("labels(%s): got error %t, want %t, err %v", tc.name, want, got, err)
		}

		if diff := cmp.Diff(tc.want.All(), s.All()); diff != "" {
			t.Errorf("labels(%s) mismatch (-want +got):\n%s", tc.name, diff)
		}
	}

}

func TestTranslate(t *testing.T) {
	prog := `// Computes R0 = 2 + 3  (R0 refers to RAM[0])
	@2
	D=A
	@3
	D=D+A
	@0
	M=D`
	wantHack := `0000000000000010
1110110000010000
0000000000000011
1110000010010000
0000000000000000
1110001100001000
`

	tests := []struct {
		name    string
		in      string
		symbols *symbol.Symbol
		want    string
		wantErr bool
	}{
		{
			name: "Prog",
			in:   prog,
			want: wantHack,
		},
	}

	for _, tc := range tests {
		buf := new(bytes.Buffer)
		err := translate(strings.NewReader(tc.in), buf, tc.symbols)
		if want, got := tc.wantErr, (err != nil); want != got {
			t.Errorf("translate(%s): got error %t, want %t, err %v", tc.name, want, got, err)
		}

		if diff := cmp.Diff(tc.want, buf.String()); diff != "" {
			t.Errorf("translate(%s) mismatch (-want +got):\n%s", tc.name, diff)
		}
	}
}

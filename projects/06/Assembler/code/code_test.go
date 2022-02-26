package code

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestComp(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "0",
			want:  "0101010",
		},
		{
			input: "1",
			want:  "0111111",
		},
		{
			input: "A",
			want:  "0110000",
		},
		{
			input: "M",
			want:  "1110000",
		},
		{
			input: "D",
			want:  "0001100",
		},
	}

	for _, tc := range tests {
		got := Comp(tc.input)
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("Comp(%s) mismatch (-want +got):\n%s", tc.input, diff)
		}
	}
}

func TestDest(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "", // null input
			want:  "000",
		},
		{
			input: "0",
			want:  "000",
		},
		{
			input: "M",
			want:  "001",
		},
		{
			input: "D",
			want:  "010",
		},
		{
			input: "MD",
			want:  "011",
		},
		{
			input: "A",
			want:  "100",
		},
		{
			input: "AM",
			want:  "101",
		},
		{
			input: "AD",
			want:  "110",
		},
		{
			input: "AMD",
			want:  "111",
		},
	}

	for _, tc := range tests {
		got := Dest(tc.input)
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("Dest(%s) mismatch (-want +got):\n%s", tc.input, diff)
		}
	}
}

func TestJump(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "", // null input
			want:  "000",
		},
		{
			input: "JGT",
			want:  "001",
		},
		{
			input: "JEQ",
			want:  "010",
		},
		{
			input: "JGE",
			want:  "011",
		},
		{
			input: "JLT",
			want:  "100",
		},
		{
			input: "JNE",
			want:  "101",
		},
		{
			input: "JLE",
			want:  "110",
		},
		{
			input: "JMP",
			want:  "111",
		},
	}

	for _, tc := range tests {
		got := Jump(tc.input)
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("Jump(%s) mismatch (-want +got):\n%s", tc.input, diff)
		}
	}
}

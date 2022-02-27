package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
		err := translate(strings.NewReader(tc.in), buf)
		if want, got := tc.wantErr, (err != nil); want != got {
			t.Errorf("translate(%s): got error %t, want %t, err %v", tc.name, want, got, err)
		}

		if diff := cmp.Diff(tc.want, buf.String()); diff != "" {
			t.Errorf("translate(%s) mismatch (-want +got):\n%s", tc.name, diff)
		}
	}
}

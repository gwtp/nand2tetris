package symbol

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSymbol(t *testing.T) {
	s := New(map[string]int{"FOO": 1})

	s.AddEntry("bar", 2)

	if diff := cmp.Diff(true, s.Contains("FOO")); diff != "" {
		t.Errorf("s.Contains() mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(true, s.Contains("bar")); diff != "" {
		t.Errorf("s.Contains() mismatch (-want +got):\n%s", diff)
	}
	// No entry for baz
	if diff := cmp.Diff(false, s.Contains("baz")); diff != "" {
		t.Errorf("s.Contains() mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(1, s.GetAddress("FOO")); diff != "" {
		t.Errorf("s.GetAddress() mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(2, s.GetAddress("bar")); diff != "" {
		t.Errorf("s.GetAddress() mismatch (-want +got):\n%s", diff)
	}
	// No entry for baz
	if diff := cmp.Diff(0, s.GetAddress("baz")); diff != "" {
		t.Errorf("s.GetAddress() mismatch (-want +got):\n%s", diff)
	}
}

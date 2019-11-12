package blacklist

import "testing"

func TestReplace(t *testing.T) {
	text := "Text"
	blacklist := "Text"
	want := ""
	if got := Replace(text, blacklist); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

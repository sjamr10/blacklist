package blacklist

import "testing"

func TestReplace(t *testing.T) {
	t.Run("text is the same as blacklist", func(t *testing.T) {
		text := "Text"
		blacklist := "Text"
		got, _ := Replace(text, blacklist)
		want := ""

		assertStrings(t, got, want)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package blacklist

import "testing"

func TestReplace(t *testing.T) {
	cases := []struct {
		name      string
		text      string
		blacklist string
		expected  string
	}{
		{
			name:      "text is the same as blacklist",
			text:      "Text",
			blacklist: "Text",
			expected:  "",
		},
		{
			name:      "text has allowed chars",
			text:      "More Text",
			blacklist: "Text",
			expected:  "More ",
		},
		{
			name:      "blacklist has more words",
			text:      "More Text",
			blacklist: "More\nText",
			expected:  " ",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _ := Replace(c.text, c.blacklist)

			assertStrings(t, got, c.expected)
		})
	}
}

func assertStrings(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

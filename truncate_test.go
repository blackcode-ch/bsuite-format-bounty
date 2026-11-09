package format

import "testing"

func TestTruncatePositiveLimits(t *testing.T) {
	cases := []struct {
		name string
		input string
		limit int
		want string
	}{
		{"empty", "", 3, ""},
		{"short", "abc", 5, "abc"},
		{"exact", "abc", 3, "abc"},
		{"one", "abc", 1, "…"},
		{"ascii", "abcd", 3, "ab…"},
		{"unicode", "界🙂ab", 3, "界🙂…"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Truncate(tc.input, tc.limit); got != tc.want {
				t.Fatalf("Truncate(%q, %d) = %q, want %q", tc.input, tc.limit, got, tc.want)
			}
		})
	}
}


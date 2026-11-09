package format

import (
	"math"
	"testing"
)

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

func TestTruncateNonPositiveAndRegressionLimits(t *testing.T) {
	cases := []struct {
		name  string
		input string
		limit int
		want  string
	}{
		{"empty zero limit", "", 0, ""},
		{"empty negative limit", "", -1, ""},
		{"empty min int limit", "", math.MinInt, ""},
		{"nonempty zero limit", "sample", 0, ""},
		{"nonempty negative one", "sample", -1, ""},
		{"nonempty negative large", "sample", -100, ""},
		{"nonempty min int limit", "sample", math.MinInt, ""},
		{"multibyte zero limit", "日本語ab", 0, ""},
		{"multibyte negative limit", "日本語ab", -5, ""},
		{"multibyte min int limit", "日本語ab", math.MinInt, ""},
		{"empty positive limit", "", 1, ""},
		{"single rune exact fit", "a", 1, "a"},
		{"ascii limit one", "sample", 1, "…"},
		{"multibyte limit one", "日本語ab", 1, "…"},
		{"multibyte exact fit", "日本語", 3, "日本語"},
		{"multibyte truncated", "日本語ab", 3, "日本…"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Truncate(tc.input, tc.limit); got != tc.want {
				t.Fatalf("Truncate(%q, %d) = %q, want %q", tc.input, tc.limit, got, tc.want)
			}
		})
	}
}
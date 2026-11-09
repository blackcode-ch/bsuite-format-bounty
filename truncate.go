package format

// Truncate limits text by Unicode code points, including a trailing ellipsis.
func Truncate(s string, n int) string {
	if len([]rune(s)) <= n {
		return s
	}
	r := []rune(s)
	return string(r[:n-1]) + "…"
}

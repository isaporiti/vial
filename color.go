package vial

const (
	colorRed   = color("\033[31m")
	colorReset = color("\033[0m")
)

// color is a simple wrapper around ANSI escape sequences
// to format terminal text.
type color string

// wrap returns the input string wrapped in the ANSI color and a reset sequence.
// For example, ColorRed.wrap("error") returns the string with red coloring applied.
func (c color) wrap(s string) string {
	return c.String() + s + colorReset.String()
}

// String returns the raw ANSI escape sequence string.
func (c color) String() string {
	return string(c)
}

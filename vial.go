// A minimalistic test assertion library for lab Gophers
package vial

// Equal reports a test failure if want and got are not equal.
// It logs the difference using a red-colored format if available.
//
// It supports special formatting for strings and runes to make
// diffs easier to read.
func Equal[C comparable](t T, want, got C) {
	if want != got {
		var format string
		switch any(want).(type) {
		case string:
			format = "want %q, got %q"
		case rune:
			format = "want %q, got %q"
		default:
			format = "want %v, got %v"
		}
		t.Errorf(ColorRed.wrap(format), want, got)
	}
}

// True reports a test failure if expr is not true.
// It is intended as a lightweight assertion for boolean expressions.
func True(t T, expr bool) {
	if !expr {
		t.Errorf(ColorRed.wrap("expression is not true"))
	}
}

// NoError reports a test failure if err is not nil.
// To assert that operations succeeded without error.
func NoError(t T, err error) {
	if err != nil {
		t.Errorf(ColorRed.wrap("unexpected error: %v"), err)
	}
}

// T is an interface that partailly matches *testing.T.
type T interface {
	Errorf(format string, args ...any)
}

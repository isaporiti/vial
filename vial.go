// A minimalistic test assertion library for lab Gophers
package vial

// Equal reports a test failure if want and got are not equal.
// It logs the difference using a red-colored format if available.
//
// It supports special formatting for strings and runes to make
// diffs easier to read.
func Equal[C comparable](t test, want, got C) {
	t.Helper()
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
		t.Errorf(format, want, got)
	}
}

// True reports a test failure if expr is not true.
// It is intended as a lightweight assertion for boolean expressions.
func True(t test, expr bool) {
	t.Helper()
	if !expr {
		t.Errorf("expression is not true")
	}
}

// NoError reports a test failure if err is not nil.
// To assert that operations succeeded without error.
func NoError(t test, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

// test is an interface that partially matches *testing.T.
type test interface {
	Errorf(format string, args ...any)
	Helper()
}

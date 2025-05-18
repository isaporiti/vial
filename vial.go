// A minimalistic test assertion library for lab Gophers
package vial

import (
	"fmt"
	"runtime"
)

// Equal reports a test failure if want and got are not equal.
// It logs the difference using a red-colored format if available.
//
// It supports special formatting for strings and runes to make
// diffs easier to read.
func Equal[C comparable](t test, want, got C, opts ...option) {
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
		o := newOptions(opts...)
		t.Errorf(colorRed.wrap("%s: "+format), o.getCaller(), want, got)
	}
}

// True reports a test failure if expr is not true.
// It is intended as a lightweight assertion for boolean expressions.
func True(t test, expr bool, opts ...option) {
	if !expr {
		o := newOptions(opts...)
		t.Errorf(colorRed.wrap("%s: expression is not true"), o.getCaller())
	}
}

// NoError reports a test failure if err is not nil.
// To assert that operations succeeded without error.
func NoError(t test, err error, opts ...option) {
	if err != nil {
		o := newOptions(opts...)
		t.Errorf(colorRed.wrap("%s: unexpected error: %v"), o.getCaller(), err)
	}
}

// test is an interface that partailly matches *testing.test.
type test interface {
	Errorf(format string, args ...any)
}

type callerFunc func() string

func defaultCaller() string {
	_, path, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", path, line)
}

type options struct {
	getCaller callerFunc
}

func newOptions(opts ...option) *options {
	o := &options{
		getCaller: defaultCaller,
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type option func(*options)

func WithCallerFunc(f callerFunc) func(o *options) {
	return func(o *options) {
		o.getCaller = f
	}
}

package test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/isaporiti/vial"
)

func TestEqual(t *testing.T) {
	for _, tc := range []struct {
		desc string
		a, b any
		want string
	}{
		{
			desc: "equal string",
			a:    "foo",
			b:    "foo",
		},
		{
			desc: "unequal string",
			a:    "foo",
			b:    "bar",
			want: "\033[31m" + `a/b/c.go:4: want "foo", got "bar"` + "\033[0m",
		},
		{
			desc: "equal int",
			a:    7,
			b:    7,
		},
		{
			desc: "unequal int",
			a:    7,
			b:    0,
			want: "\033[31m" + "a/b/c.go:4: want 7, got 0" + "\033[0m",
		},
		{
			desc: "equal float32",
			a:    0.53,
			b:    0.53,
		},
		{
			desc: "unequal float32",
			a:    1.53,
			b:    0.53,
			want: "\033[31m" + "a/b/c.go:4: want 1.53, got 0.53" + "\033[0m",
		},
		{
			desc: "equal bool",
			a:    true,
			b:    true,
		},
		{
			desc: "unequal bool",
			a:    true,
			b:    false,
			want: "\033[31m" + "a/b/c.go:4: want true, got false" + "\033[0m",
		},
		{
			desc: "equal rune",
			a:    'z',
			b:    'z',
		},
		{
			desc: "unequal rune",
			a:    'a',
			b:    'c',
			want: "\033[31m" + "a/b/c.go:4: want 'a', got 'c'" + "\033[0m",
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			out := &bytes.Buffer{}
			mt := mockT{out: out}

			vial.Equal(&mt, tc.a, tc.b, vial.WithCallerFunc(mockCallerFunc))

			if tc.want == "" && mt.failed {
				t.Error("want test to pass, but it failed")
			}
			if tc.want != "" && !mt.failed {
				t.Error("want test to fail, but did not")
			}

			got := out.String()
			if tc.want != got {
				t.Errorf("%q != %q", tc.want, got)
			}
		})
	}
}

func TestTrue(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		t.Parallel()
		var mt mockT

		vial.True(&mt, "foo" != "bar", vial.WithCallerFunc(mockCallerFunc))

		if mt.failed {
			t.Error("want mock T not to fail, but did")
		}
	})

	t.Run("not true", func(t *testing.T) {
		t.Parallel()
		out := &bytes.Buffer{}
		mt := mockT{out: out}

		vial.True(&mt, "foo" == "bar", vial.WithCallerFunc(mockCallerFunc))

		if !mt.failed {
			t.Error("want test to fail, but did not")
		}

		want := "\033[31m" + "a/b/c.go:4: expression is not true" + "\033[0m"
		got := out.String()
		if want != got {
			t.Errorf("%q != %q", want, got)
		}
	})
}

func TestNoError(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		t.Parallel()
		var mt mockT

		vial.NoError(&mt, nil, vial.WithCallerFunc(mockCallerFunc))

		if mt.failed {
			t.Error("want mock T not to fail, but did")
		}
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		out := &bytes.Buffer{}
		mt := mockT{out: out}
		err := errors.New("uh-oh")

		vial.NoError(&mt, err, vial.WithCallerFunc(mockCallerFunc))

		if !mt.failed {
			t.Error("want mock T to fail, but didn't")
		}

		want := "\033[31m" +
			"a/b/c.go:4: " +
			"unexpected error: " +
			err.Error() +
			"\033[0m"
		got := out.String()
		if want != got {
			t.Errorf("%q != %q", want, got)
		}
	})
}

type mockT struct {
	failed bool
	out    io.Writer
}

func (t *mockT) Errorf(format string, args ...any) {
	t.failed = true
	fmt.Fprintf(t.out, format, args...)
}

func mockCallerFunc() string {
	return "a/b/c.go:4"
}

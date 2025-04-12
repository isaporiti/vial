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
			want: vial.ColorRed.String() + `want "foo", got "bar"` + vial.ColorReset.String(),
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
			want: vial.ColorRed.String() + "want 7, got 0" + vial.ColorReset.String(),
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
			want: vial.ColorRed.String() + "want 1.53, got 0.53" + vial.ColorReset.String(),
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
			want: vial.ColorRed.String() + "want true, got false" + vial.ColorReset.String(),
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
			want: vial.ColorRed.String() + "want 'a', got 'c'" + vial.ColorReset.String(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			out := &bytes.Buffer{}
			mt := mockT{out: out}

			vial.Equal(&mt, tc.a, tc.b)

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

		vial.True(&mt, "foo" != "bar")

		if mt.failed {
			t.Error("want mock T not to fail, but did")
		}
	})

	t.Run("not true", func(t *testing.T) {
		t.Parallel()
		out := &bytes.Buffer{}
		mt := mockT{out: out}

		vial.True(&mt, "foo" == "bar")

		if !mt.failed {
			t.Error("want test to fail, but did not")
		}

		want := vial.ColorRed.String() + "expression is not true" + vial.ColorReset.String()
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

		vial.NoError(&mt, nil)

		if mt.failed {
			t.Error("want mock T not to fail, but did")
		}
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		out := &bytes.Buffer{}
		mt := mockT{out: out}
		err := errors.New("uh-oh")

		vial.NoError(&mt, err)

		if !mt.failed {
			t.Error("want mock T to fail, but didn't")
		}

		want := vial.ColorRed.String() +
			"unexpected error: " +
			err.Error() +
			vial.ColorReset.String()
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

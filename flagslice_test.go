package flagslice_test

/*
 * flagslice_test.go
 * Test for flagslice
 * By J. Stuart McMurray
 * Created 20171009
 * Last Modified 20171009
 */

import (
	"flag"
	"testing"

	"github.com/magisterquis/flagslice"
)

func TestString(t *testing.T) {
	var (
		p    *[]string
		in   = s("-f", "foo", "-f=bar")
		want = s("foo", "bar")
	)

	/* Test string with nil default */
	p = testParse(nil, in)
	if !cs(*p, want) {
		t.Errorf("Default nil %v: Got %v, Want %v", in, *p, want)
	}

	/* With empty default */
	p = testParse(s(), in)
	if !cs(*p, want) {
		t.Errorf("No default %v: Got %v, Want %v", in, *p, want)
	}

	/* Test string with default and arguments */
	p = testParse(s("tridge", "baaz"), in)
	if !cs(*p, want) {
		t.Errorf("With default %v: Got %v, Want %v", in, *p, want)
	}

	/* With nil default and no arguments */
	p = testParse(nil, s())
	if !cs(*p, s()) {
		t.Errorf("Default nil and no arguments: Got %v", *p)
	}
}

/* parseFlag tests the given arguments to String in a new flag.FlagSet.  It
panics if argv doesn't parse.  "f" is used as the flag name. */
func testParse(def, argv []string) *[]string {
	fs := flag.NewFlagSet("", flag.PanicOnError)
	s := flagslice.StringFS(fs, "f", def, "ignored")
	fs.Parse(argv)
	return s
}

/* s returns its arguments as a string slice */
func s(a ...string) []string { return a }

/* cs returns false if the two slices are different */
func cs(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

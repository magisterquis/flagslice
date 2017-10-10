/*
Package flagslice provides what should have been flag.StringSlice

Examples:

Declare a command-line argument -s which may repeated multiple times
	ss := flagslice.String("s", nil, "A `repeatable string` flag")
	for i,s := range *ss {
		fmt.Printf("%v: %v", i, s)
	}

As above, but passing in a pointer and a default
	var ss *string
	flagslice.StringVar(
		ss,
		"s",
		[]string{"foo", "bar},
		"A `repeatable string` with a default",
	)
	for i,s := range *ss {
		fmt.Printf("%v: %v", i, s)
	}

Both of the above examples could be called with
	./foo -s bar -s tridge -s baaz
Which would have the same effect as
	ss := &[]string{"bar", "tridge", "baaz"}
*/
package flagslice

/*
 * flagslice.go
 * Provide what should have been flag.StringSlice
 * By J. Stuart McMurray
 * Created 20171009
 * Last Modified 20171009
 */

import (
	"flag"
	"fmt"
	"sync"
)

/* stringSlice and its supporting functions wrap a string in a flag.Value */
type stringSlice struct {
	s *[]string /* Underlying array */
	d bool      /* True if we're replacing the default */
	l *sync.Mutex
}

func (s *stringSlice) String() string {
	if nil == s.s || 0 == len(*s.s) {
		return ""
	}
	return fmt.Sprintf("%v", *s.s)
}

func (s *stringSlice) Set(v string) error {
	s.l.Lock()
	defer s.l.Unlock()
	/* Clear the default array if it's the first use */
	if !s.d {
		*s.s = []string{}
		s.d = true
	}
	*s.s = append(*s.s, v)
	return nil
}

// String is analogous to flag.String, except it returns a slice of strings
// which will be filled after flag.Parse returns.  The slice passed in as the
// default value will not be modified.
func String(name string, value []string, usage string) *[]string {
	return StringFS(flag.CommandLine, name, value, usage)
}

// StringFS is like String, but operates on the specified flag.FlagSet
func StringFS(f *flag.FlagSet, name string, value []string, usage string) *[]string {
	ss := &[]string{}
	StringVarFS(f, ss, name, value, usage)
	return ss
}

// StringVar is analogous to flag.StringVar, except the first argument points
// to a slice which will be updated after flag.Parse returns.
func StringVar(p *[]string, name string, value []string, usage string) {
	StringVarFS(flag.CommandLine, p, name, value, usage)
}

// StringVarFS is like StringVar but operates on the specified flag.FlagSet
func StringVarFS(f *flag.FlagSet, p *[]string, name string, value []string, usage string) {
	s := &stringSlice{s: p, l: &sync.Mutex{}}
	if nil != value {
		*s.s = append(*s.s, value...)
	}
	f.Var(s, name, usage)
}

flagslice
=========
Flagslice provides what should have been `flag.StringSlice`.  It is likely that
in the future other types will be suppotred as well.

[![GoDoc](https://godoc.org/github.com/magisterquis/flagslice?status.svg)](https://godoc.org/github.com/magisterquis/flagslice)

Usage
-----
Use much like `flag.String`:
```go
files := flagslice.String("f", nil, "A `repeatable filename`")
pets := flagslice.String("n", []string{"cats", "dogs"}, "Fuzzy `pets`")
flag.Parse()
/* Do something with *files and *names */
```

A program using the above would be invoked somethnig like
```sh
./foo -f shadow -f master.passwd -n bunnies -n moose
```

<h1 align="center">
    <img alt="gopher illustration" src="https://spirited.io/wp-content/uploads/elementor/thumbs/1-prc44rwae7cvpfcnuotwqkc46fiz99oyv9553ip1tc.jpg">
</h1>

<p align="center">
	<a href="https://pkg.go.dev/github.com/negrel/assert">
		<img alt="PkgGoDev" src="https://pkg.go.dev/badge/github.com/negrel/assert">
	</a>
	<a href="https://goreportcard.com/report/github.com/negrel/assert">
		<img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/negrel/assert">
	</a>
</p>

# `assert` - Zero cost debug assertions.

This package provides zero cost debug assertions to your Go programs. It is based
on the excellent [`github.com/stretchr/testify/assert`](https://github.com/stretchr/testify)
package and provide the same API (minus `t testing.T` parameter).

## Why?

This is a complete rewrite of [`debuggo`](https://github.com/negrel/debuggo) that aims
to be up to date and more maintainable.

It aims to provide the same API as [`github.com/stretchr/testify/assert`](https://github.com/stretchr/testify).
* Prints friendly, easy to read failure descriptions
* Allows for very readable code
* Optionally annotate each assertion with a message
* No performance impact on production build (see benchmarks)

Debug assertions main use is to assert invariant that can't be encoded in the
type systems. For example, a method that should be called only when mutex is
locked:

```go
package mypkg

import (
	"sync"
	"github.com/negrel/assert"
)

type myType struct {
	mu sync.Mutex
	// other fields...
}

// doWork perform internal work. Caller must hold mutex while calling this
// function.
func (mt *myType) doWork(k string) {
	assert.Locked(&mt.mu) // panic if assertions are enabled and mutex isn't locked

	// Do work...
}
```

## How does it works?

[Read my blog post](https://www.negrel.dev/blog/zero-cost-debug-assertions-in-go/)
about to understand how `assert` works and why it is designed that way.

## Getting started

Here is our example program:

```go
package main

import (
	"github.com/negrel/assert"
)

func main() {
	assert.True(false)
	println("Hello world!")
}
```

A simple `go run .` will simply print `Hello world!` as all `assert` functions
are removed by the compiler.

Now, if we compile and run it with assertions enabled `go run -tags assert .`,
it will output something like:

```
panic:
        Error Trace:    /home/anegrel/code/go/assert/example/main.go:8
                                                /usr/share/go/src/runtime/proc.go:267
                                                /usr/share/go/src/runtime/asm_amd64.s:1650
        Error:          Should be true


goroutine 1 [running]:
github.com/negrel/assert.Fail({0x568dc2, 0xe}, {0x0, 0x0, 0x0})
        /home/anegrel/code/go/assert/assertions.go:349 +0x168
github.com/negrel/assert.True(...)
        /home/anegrel/code/go/assert/assertions.go:754
main.main()
        /home/anegrel/code/go/assert/example/main.go:8 +0x27
exit status 2
```

Note that most `go` subcommands (build, run, test, ...) supports `-tags` flag.
You may want to set `GOFLAGS` environment variable to `-tags assert` make it
permanent and avoid specifying it on each command.

## Benchmarks

As we've seen previously, assertions are hidden behind a compilation flag. If
the flag is absent, all assertions functions will be empty/noop function that
the compiler will optimize.

**WITH** `-tags assert`:

```
goos: linux
goarch: amd64
pkg: github.com/negrel/assert/tests
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkSliceIndexWithoutBoundCheckAssertions
BenchmarkSliceIndexWithoutBoundCheckAssertions-8        728439501                1.407 ns/op
BenchmarkSliceIndexWithBoundCheckAssertions
BenchmarkSliceIndexWithBoundCheckAssertions-8           27423670                40.80 ns/op
PASS
ok      github.com/negrel/assert/tests  3.338s
```

**WITHOUT** `-tags assert`:

```
goos: linux
goarch: amd64
pkg: github.com/negrel/assert/tests
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkSliceIndexWithoutBoundCheckAssertions
BenchmarkSliceIndexWithoutBoundCheckAssertions-8        772181695                1.399 ns/op
BenchmarkSliceIndexWithBoundCheckAssertions
BenchmarkSliceIndexWithBoundCheckAssertions-8           802181890                1.412 ns/op
PASS
ok      github.com/negrel/assert/tests  2.531s
```

However, keep in mind that `assert` may slightly increase binary size (~100 KiB)
as it imports `net/http` and `reflect`.

## Contributing

If you want to contribute to `assert` to add a feature or improve the code contact
me at [negrel.dev@protonmail.com](mailto:negrel.dev@protonmail.com), open an
[issue](https://github.com/negrel/assert/issues) or make a
[pull request](https://github.com/negrel/assert/pulls).

## :stars: Show your support

Please give a :star: if this project helped you!

[![buy me a coffee](.github/images/bmc-button.png)](https://www.buymeacoffee.com/negrel)

## :scroll: License

MIT © [Alexandre Negrel](https://www.negrel.dev/)

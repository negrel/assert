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

## How does it works?

[Read my blog post](https://www.negrel.dev/blog/zero-cost-debug-assertions-in-go/)
about to understand how `assert` works and why it is designed that way.

## Getting started

Here is our example program:

```go
package main

import "github.com/negrel/assert"

func safeIndex(slice []string, index int) string {
	// Ensure index is not out of bounds.
	assert.GreaterOrEqual(index, 0, "negative index not allowed")
	assert.Lessf(index, len(slice), "index out of bounds (slice: %v)", slice)

	return slice[index]
}

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	println(safeIndex(days, 8))
}
```

A simple `go run .` will produce the following error:
```
panic: runtime error: index out of range [8] with length 7

goroutine 1 [running]:
main.safeIndex(...)
        /home/anegrel/code/go/assert/example/main.go:10
main.main()
        /home/anegrel/code/go/assert/example/main.go:15 +0xd6
exit status 2
```

Now, if you enable assertions with a compile time flags, `go run -tags assert .`,
you will get something similar to:

```
panic:
        Error Trace:    /home/anegrel/code/go/assert/example/main.go:8
                                                /home/anegrel/code/go/assert/example/main.go:15
                                                /nix/store/dwmb0qcai52d0zkgpm6f5ifx2a8yvsdg-go-1.21.3/share/go/src/runtime/proc.go:267
                                                /nix/store/dwmb0qcai52d0zkgpm6f5ifx2a8yvsdg-go-1.21.3/share/go/src/runtime/asm_amd64.s:1650
        Error:          "8" is not less than "7"
        Messages:       index out of bounds (slice: [Monday Tuesday Wednesday Thursday Friday Saturday Sunday])


goroutine 1 [running]:
github.com/negrel/assert.Fail({0xc0001e2168, 0x18}, {0xc0001be240, 0x2, 0x2})
        /home/anegrel/code/go/assert/assertions.go:331 +0x168
github.com/negrel/assert.compareTwoValues({0x53eb00, 0x695cc0}, {0x53eb00, 0x695cb8}, {0xc0001f5e88, 0x1, 0x40e285?}, {0x57307c, 0x1a}, {0xc0001be240, ...})
        /home/anegrel/code/go/assert/assertion_compare.go:425 +0x2af
github.com/negrel/assert.Less(...)
        /home/anegrel/code/go/assert/assertion_compare.go:380
github.com/negrel/assert.Lessf(...)
        /home/anegrel/code/go/assert/assertion_format.go:355
main.safeIndex({0xc0001bdea0?, 0x7, 0x7}, 0x8)
        /home/anegrel/code/go/assert/example/main.go:8 +0x234
main.main()
        /home/anegrel/code/go/assert/example/main.go:15 +0xb4
exit status 2
```

That's the same output as `testify/assert` output except that we have a stacktrace
because this is a panic.

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

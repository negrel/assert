// Package assert provides a set of comprehensive debug assertions.
//
// # Example Usage
//
// The following is a complete example using assert in a function:
//
//	import (
//	  "github.com/negrel/assert"
//	)
//
//	func Divide(a, b int) int {
//	  assert.NotEqual(b, 0, "Can't divide by 0.")
//	  return a / b
//	}
//
// # Debug Assertions
//
// Debug assertions are programming statements that help developers catch
// logical errors during development by verifying assumptions about the
// program's state. They're typically used to check conditions that should
// always be true during normal program execution. If the condition is false,
// the assertions is either wrong or there is a programming error. In
// both case, program panics if it was compiled with the `assert` tags
// (`go build -tags assert ./path/to/my/package`)
//
// Every assertion function also takes an optional string message as the final argument,
// allowing custom error messages to be appended to the message the assertion method outputs.
package assert

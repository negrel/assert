//go:build !assert

package assert

import (
	"reflect"
	"time"

	"github.com/davecgh/go-spew/spew"
)

//go:generate sh -c "cd ../_codegen && go build && cd - && ../_codegen/_codegen -output-package=assert -template=assertion_format.go.tmpl"

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Errorf(format string, args ...interface{})
}

// ComparisonAssertionFunc is a common function prototype when comparing two values.  Can be useful
// for table driven tests.
type ComparisonAssertionFunc func(interface{}, interface{}, ...interface{}) bool

// ValueAssertionFunc is a common function prototype when validating a single value.  Can be useful
// for table driven tests.
type ValueAssertionFunc func(interface{}, ...interface{}) bool

// BoolAssertionFunc is a common function prototype when validating a bool value.  Can be useful
// for table driven tests.
type BoolAssertionFunc func(bool, ...interface{}) bool

// ErrorAssertionFunc is a common function prototype when validating an error value.  Can be useful
// for table driven tests.
type ErrorAssertionFunc func(error, ...interface{}) bool

// Comparison is a custom function that returns true on success and false on failure
type Comparison func() (success bool)

/*
	Helper functions
*/

// ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func ObjectsAreEqual(expected, actual interface{}) {}

// copyExportedFields iterates downward through nested data structures and creates a copy
// that only contains the exported struct fields.
func copyExportedFields(expected interface{}) {}

// ObjectsExportedFieldsAreEqual determines if the exported (public) fields of two objects are
// considered equal. This comparison of only exported fields is applied recursively to nested data
// structures.
//
// This function does no assertion of any kind.
//
// Deprecated: Use [EqualExportedValues] instead.
func ObjectsExportedFieldsAreEqual(expected, actual interface{}) {}

// ObjectsAreEqualValues gets whether two objects are equal, or if their
// values are equal.
func ObjectsAreEqualValues(expected, actual interface{}) {}

// isNumericType returns true if the type is one of:
// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
// float32, float64, complex64, complex128
func isNumericType(t reflect.Type) {}

/* CallerInfo is necessary because the assert functions use the testing object
internally, causing it to print the file:line of the assert method, rather than where
the problem actually occurred in calling code.*/

// CallerInfo returns an array of strings containing the file and line number
// of each stack frame leading from the current test to the assert call that
// failed.
func CallerInfo() {}

// Stolen from the `go test` tool.
// isTest tells whether name looks like a test (or benchmark, according to prefix).
// It is a Test (say) if there is a character after Test that is not a lower-case letter.
// We don't want TesticularCancer.
func isTest(name, prefix string) {}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) {}

// Aligns the provided message so that all lines after the first line start at the same location as the first line.
// Assumes that the first line starts at the correct location (after carriage return, tab, label, spacer and tab).
// The longestLabelLen parameter specifies the length of the longest label in the output (required because this is the
// basis on which the alignment occurs).
func indentMessageLines(message string, longestLabelLen int) {}

type failNower interface {
	FailNow()
}

// FailNow fails test
func FailNow(failureMessage string, msgAndArgs ...interface{}) {}

// Fail reports a failure through
func Fail(failureMessage string, msgAndArgs ...interface{}) {}

type labeledContent struct {
	label   string
	content string
}

// labeledOutput returns a string consisting of the provided labeledContent. Each labeled output is appended in the following manner:
//
//	\t{{label}}:{{align_spaces}}\t{{content}}\n
//
// The initial carriage return is required to undo/erase any padding added by testing.T.Errorf. The "\t{{label}}:" is for the label.
// If a label is shorter than the longest label provided, padding spaces are added to make all the labels match in length. Once this
// alignment is achieved, "\t{{content}}\n" is added for the output.
//
// If the content of the labeledOutput contains line breaks, the subsequent lines are aligned so that they start at the same location as the first line.
func labeledOutput(content ...labeledContent) {}

// Implements asserts that an object is implemented by the specified interface.
//
//	assert.Implements((*MyInterface)(nil), new(MyObject))
func Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {}

// NotImplements asserts that an object does not implement the specified interface.
//
//	assert.NotImplements((*MyInterface)(nil), new(MyObject))
func NotImplements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {}

// IsType asserts that the specified objects are of the same type.
func IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {}

// Equal asserts that two objects are equal.
//
//	assert.Equal(123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equal(expected, actual interface{}, msgAndArgs ...interface{}) {}

// validateEqualArgs checks whether provided arguments can be safely used in the
// Equal/NotEqual functions.
func validateEqualArgs(expected, actual interface{}) {}

// Same asserts that two pointers reference the same object.
//
//	assert.Same(ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func Same(expected, actual interface{}, msgAndArgs ...interface{}) {}

// NotSame asserts that two pointers do not reference the same object.
//
//	assert.NotSame(ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func NotSame(expected, actual interface{}, msgAndArgs ...interface{}) {}

// samePointers compares two generic interface objects and returns whether
// they point to the same object
func samePointers(first, second interface{}) {}

// formatUnequalValues takes two values of arbitrary types and returns string
// representations appropriate to be presented to the user.
//
// If the values are not of like type, the returned strings will be prefixed
// with the type name, and the value will be enclosed in parentheses similar
// to a type conversion in the Go grammar.
func formatUnequalValues(expected, actual interface{}) {}

// truncatingFormat formats the data and truncates it if it's too long.
//
// This helps keep formatted error messages lines from exceeding the
// bufio.MaxScanTokenSize max line length that the go testing framework imposes.
func truncatingFormat(data interface{}) {}

// EqualValues asserts that two objects are equal or convertible to the same types
// and equal.
//
//	assert.EqualValues(uint32(123), int32(123))
func EqualValues(expected, actual interface{}, msgAndArgs ...interface{}) {}

// EqualExportedValues asserts that the types of two objects are equal and their public
// fields are also equal. This is useful for comparing structs that have private fields
// that could potentially differ.
//
//	 type S struct {
//		Exported     	int
//		notExported   	int
//	 }
//	 assert.EqualExportedValues(S{1, 2}, S{1, 3}) => true
//	 assert.EqualExportedValues(S{1, 2}, S{2, 3}) => false
func EqualExportedValues(expected, actual interface{}, msgAndArgs ...interface{}) {}

// Exactly asserts that two objects are equal in value and type.
//
//	assert.Exactly(int32(123), int64(123))
func Exactly(expected, actual interface{}, msgAndArgs ...interface{}) {}

// NotNil asserts that the specified object is not nil.
//
//	assert.NotNil(err)
func NotNil(object interface{}, msgAndArgs ...interface{}) {}

// isNil checks if a specified object is nil or not, without Failing.
func isNil(object interface{}) {}

// Nil asserts that the specified object is nil.
//
//	assert.Nil(err)
func Nil(object interface{}, msgAndArgs ...interface{}) {}

// isEmpty gets whether the specified object is considered empty or not.
func isEmpty(object interface{}) {}

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//	assert.Empty(obj)
func Empty(object interface{}, msgAndArgs ...interface{}) {}

// NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//	if assert.NotEmpty(obj) {
//	  assert.Equal("two", obj[1])
//	}
func NotEmpty(object interface{}, msgAndArgs ...interface{}) {}

// getLen tries to get the length of an object.
// It returns (0, false) if impossible.
func getLen(x interface{}) {}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//	assert.Len(mySlice, 3)
func Len(object interface{}, length int, msgAndArgs ...interface{}) {}

// True asserts that the specified value is true.
//
//	assert.True(myBool)
func True(value bool, msgAndArgs ...interface{}) {}

// False asserts that the specified value is false.
//
//	assert.False(myBool)
func False(value bool, msgAndArgs ...interface{}) {}

// NotEqual asserts that the specified values are NOT equal.
//
//	assert.NotEqual(obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func NotEqual(expected, actual interface{}, msgAndArgs ...interface{}) {}

// NotEqualValues asserts that two objects are not equal even when converted to the same type
//
//	assert.NotEqualValues(obj1, obj2)
func NotEqualValues(expected, actual interface{}, msgAndArgs ...interface{}) {}

// containsElement try loop over the list check if the list includes the element.
// return (false, false) if impossible.
// return (true, false) if element was not found.
// return (true, true) if element was found.
func containsElement(list interface{}, element interface{}) {}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//	assert.Contains("Hello World", "World")
//	assert.Contains(["Hello", "World"], "World")
//	assert.Contains({"Hello": "World"}, "Hello")
func Contains(s, contains interface{}, msgAndArgs ...interface{}) {}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//	assert.NotContains("Hello World", "Earth")
//	assert.NotContains(["Hello", "World"], "Earth")
//	assert.NotContains({"Hello": "World"}, "Earth")
func NotContains(s, contains interface{}, msgAndArgs ...interface{}) {}

// Subset asserts that the specified list(array, slice...) or map contains all
// elements given in the specified subset list(array, slice...) or map.
//
//	assert.Subset([1, 2, 3], [1, 2])
//	assert.Subset({"x": 1, "y": 2}, {"x": 1})
func Subset(list, subset interface{}, msgAndArgs ...interface{}) {}

// NotSubset asserts that the specified list(array, slice...) or map does NOT
// contain all elements given in the specified subset list(array, slice...) or
// map.
//
//	assert.NotSubset([1, 3, 4], [1, 2])
//	assert.NotSubset({"x": 1, "y": 2}, {"z": 3})
func NotSubset(list, subset interface{}, msgAndArgs ...interface{}) {}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatch([1, 3, 2, 3], [1, 3, 3, 2])
func ElementsMatch(listA, listB interface{}, msgAndArgs ...interface{}) {}

// isList checks that the provided value is array or slice.
func isList(list interface{}, msgAndArgs ...interface{}) {}

// diffLists diffs two arrays/slices and returns slices of elements that are only in A and only in B.
// If some element is present multiple times, each instance is counted separately (e.g. if something is 2x in A and
// 5x in B, it will be 0x in extraA and 3x in extraB). The order of items in both lists is ignored.
func diffLists(listA, listB interface{}) {}

func formatListDiff(listA, listB interface{}, extraA, extraB []interface{}) {}

// Condition uses a Comparison to assert a complex condition.
func Condition(comp Comparison, msgAndArgs ...interface{}) {}

// PanicTestFunc defines a func that should be passed to the assert.Panics and assert.NotPanics
// methods, and represents a simple func that takes no arguments, and returns nothing.
type PanicTestFunc func()

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
func didPanic(f PanicTestFunc) {}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//	assert.Panics(func(){ GoCrazy() })
func Panics(f PanicTestFunc, msgAndArgs ...interface{}) {}

// PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//	assert.PanicsWithValue("crazy error", func(){ GoCrazy() })
func PanicsWithValue(expected interface{}, f PanicTestFunc, msgAndArgs ...interface{}) {}

// PanicsWithError asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
//	assert.PanicsWithError("crazy error", func(){ GoCrazy() })
func PanicsWithError(errString string, f PanicTestFunc, msgAndArgs ...interface{}) {}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//	assert.NotPanics(func(){ RemainCalm() })
func NotPanics(f PanicTestFunc, msgAndArgs ...interface{}) {}

// WithinDuration asserts that the two times are within duration delta of each other.
//
//	assert.WithinDuration(time.Now(), time.Now(), 10*time.Second)
func WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {}

// WithinRange asserts that a time is within a time range (inclusive).
//
//	assert.WithinRange(time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second))
func WithinRange(actual, start, end time.Time, msgAndArgs ...interface{}) {}

func toFloat(x interface{}) {}

// InDelta asserts that the two numerals are within delta of each other.
//
//	assert.InDelta(math.Pi, 22/7.0, 0.01)
func InDelta(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {}

// InDeltaSlice is the same as InDelta, except it compares two slices.
func InDeltaSlice(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {}

// InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
func InDeltaMapValues(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {}

func calcRelativeError(expected, actual interface{}) {}

// InEpsilon asserts that expected and actual have a relative error less than epsilon
func InEpsilon(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
func InEpsilonSlice(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {}

/*
	Errors
*/

// NoError asserts that a function returned no error (i.e. `nil`).
//
//	  actualObj, err := SomeFunction()
//	  if assert.NoError(err) {
//		   assert.Equal(expectedObj, actualObj)
//	  }
func NoError(err error, msgAndArgs ...interface{}) {}

// Error asserts that a function returned an error (i.e. not `nil`).
//
//	  actualObj, err := SomeFunction()
//	  if assert.Error(err) {
//		   assert.Equal(expectedError, err)
//	  }
func Error(err error, msgAndArgs ...interface{}) {}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//	actualObj, err := SomeFunction()
//	assert.EqualError(err,  expectedErrorString)
func EqualError(theError error, errString string, msgAndArgs ...interface{}) {}

// ErrorContains asserts that a function returned an error (i.e. not `nil`)
// and that the error contains the specified substring.
//
//	actualObj, err := SomeFunction()
//	assert.ErrorContains(err,  expectedErrorSubString)
func ErrorContains(theError error, contains string, msgAndArgs ...interface{}) {}

// matchRegexp return true if a specified regexp matches a string.
func matchRegexp(rx interface{}, str interface{}) {}

// Regexp asserts that a specified regexp matches a string.
//
//	assert.Regexp(regexp.MustCompile("start"), "it's starting")
//	assert.Regexp("start...$", "it's not starting")
func Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {}

// NotRegexp asserts that a specified regexp does not match a string.
//
//	assert.NotRegexp(regexp.MustCompile("starts"), "it's starting")
//	assert.NotRegexp("^start", "it's not starting")
func NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {}

// Zero asserts that i is the zero value for its type.
func Zero(i interface{}, msgAndArgs ...interface{}) {}

// NotZero asserts that i is not the zero value for its type.
func NotZero(i interface{}, msgAndArgs ...interface{}) {}

// FileExists checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
func FileExists(path string, msgAndArgs ...interface{}) {}

// NoFileExists checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
func NoFileExists(path string, msgAndArgs ...interface{}) {}

// DirExists checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
func DirExists(path string, msgAndArgs ...interface{}) {}

// NoDirExists checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
func NoDirExists(path string, msgAndArgs ...interface{}) {}

// JSONEq asserts that two JSON strings are equivalent.
//
//	assert.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
func JSONEq(expected string, actual string, msgAndArgs ...interface{}) {}

// YAMLEq asserts that two YAML strings are equivalent.
func YAMLEq(expected string, actual string, msgAndArgs ...interface{}) {}

func typeAndKind(v interface{}) {}

// diff returns a diff of both values as long as both are of the same type and
// are a struct, map, slice, array or string. Otherwise it returns an empty string.
func diff(expected interface{}, actual interface{}) {}

func isFunction(arg interface{}) {}

var spewConfig = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
	DisableMethods:          true,
	MaxDepth:                10,
}

var spewConfigStringerEnabled = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
	MaxDepth:                10,
}

type tHelper interface {
	Helper()
}

// Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//	assert.Eventually(func() bool { return true; }, time.Second, 10*time.Millisecond)
func Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
}

// CollectT implements the TestingT interface and collects all errors.
type CollectT struct {
	errors []error
}

// Errorf collects the error.
func (c *CollectT) Errorf(format string, args ...interface{}) {}

// FailNow panics.
func (*CollectT) FailNow() {}

// Deprecated: That was a method for internal usage that should not have been published. Now just panics.
func (*CollectT) Reset() {}

// Deprecated: That was a method for internal usage that should not have been published. Now just panics.
func (*CollectT) Copy(TestingT) {}

// EventuallyWithT asserts that given condition will be met in waitFor time,
// periodically checking target function each tick. In contrast to Eventually,
// it supplies a CollectT to the condition function, so that the condition
// function can use the CollectT to call other assertions.
// The condition is considered "met" if no errors are raised in a tick.
// The supplied CollectT collects all errors from one tick (if there are any).
// If the condition is not met before waitFor, the collected errors of
// the last tick are copied to t.
//
//	externalValue := false
//	go func() {}

// Never asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//	assert.Never(func() bool { return false; }, time.Second, 10*time.Millisecond)
func Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
}

// ErrorIs asserts that at least one of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func ErrorIs(err, target error, msgAndArgs ...interface{}) {}

// NotErrorIs asserts that at none of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func NotErrorIs(err, target error, msgAndArgs ...interface{}) {}

// ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
// This is a wrapper for errors.As.
func ErrorAs(err error, target interface{}, msgAndArgs ...interface{}) {}

func buildErrorChainString(err error) {}

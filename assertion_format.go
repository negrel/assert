//go:build assert

// Code generated with github.com/stretchr/testify/_codegen; DO NOT EDIT.

package assert

import (
	http "net/http"
	url "net/url"
	time "time"
)

// Conditionf uses a Comparison to assert a complex condition.
func Conditionf(comp Comparison, msg string, args ...interface{}) bool {
	return Condition(comp, append([]interface{}{msg}, args...)...)
}

// Containsf asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//	assert.Containsf("Hello World", "World", "error message %s", "formatted")
//	assert.Containsf(["Hello", "World"], "World", "error message %s", "formatted")
//	assert.Containsf({"Hello": "World"}, "Hello", "error message %s", "formatted")
func Containsf(s interface{}, contains interface{}, msg string, args ...interface{}) bool {
	return Contains(s, contains, append([]interface{}{msg}, args...)...)
}

// DirExistsf checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
func DirExistsf(path string, msg string, args ...interface{}) bool {
	return DirExists(path, append([]interface{}{msg}, args...)...)
}

// ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatchf([1, 3, 2, 3], [1, 3, 3, 2], "error message %s", "formatted")
func ElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{}) bool {
	return ElementsMatch(listA, listB, append([]interface{}{msg}, args...)...)
}

// Emptyf asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//	assert.Emptyf(obj, "error message %s", "formatted")
func Emptyf(object interface{}, msg string, args ...interface{}) bool {
	return Empty(object, append([]interface{}{msg}, args...)...)
}

// Equalf asserts that two objects are equal.
//
//	assert.Equalf(123, 123, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equalf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return Equal(expected, actual, append([]interface{}{msg}, args...)...)
}

// EqualErrorf asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//	actualObj, err := SomeFunction()
//	assert.EqualErrorf(err,  expectedErrorString, "error message %s", "formatted")
func EqualErrorf(theError error, errString string, msg string, args ...interface{}) bool {
	return EqualError(theError, errString, append([]interface{}{msg}, args...)...)
}

// EqualExportedValuesf asserts that the types of two objects are equal and their public
// fields are also equal. This is useful for comparing structs that have private fields
// that could potentially differ.
//
//	 type S struct {
//		Exported     	int
//		notExported   	int
//	 }
//	 assert.EqualExportedValuesf(S{1, 2}, S{1, 3}, "error message %s", "formatted") => true
//	 assert.EqualExportedValuesf(S{1, 2}, S{2, 3}, "error message %s", "formatted") => false
func EqualExportedValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return EqualExportedValues(expected, actual, append([]interface{}{msg}, args...)...)
}

// EqualValuesf asserts that two objects are equal or convertible to the same types
// and equal.
//
//	assert.EqualValuesf(uint32(123), int32(123), "error message %s", "formatted")
func EqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return EqualValues(expected, actual, append([]interface{}{msg}, args...)...)
}

// Errorf asserts that a function returned an error (i.e. not `nil`).
//
//	  actualObj, err := SomeFunction()
//	  if assert.Errorf(err, "error message %s", "formatted") {
//		   assert.Equal(expectedErrorf, err)
//	  }
func Errorf(err error, msg string, args ...interface{}) bool {
	return Error(err, append([]interface{}{msg}, args...)...)
}

// ErrorAsf asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
// This is a wrapper for errors.As.
func ErrorAsf(err error, target interface{}, msg string, args ...interface{}) bool {
	return ErrorAs(err, target, append([]interface{}{msg}, args...)...)
}

// ErrorContainsf asserts that a function returned an error (i.e. not `nil`)
// and that the error contains the specified substring.
//
//	actualObj, err := SomeFunction()
//	assert.ErrorContainsf(err,  expectedErrorSubString, "error message %s", "formatted")
func ErrorContainsf(theError error, contains string, msg string, args ...interface{}) bool {
	return ErrorContains(theError, contains, append([]interface{}{msg}, args...)...)
}

// ErrorIsf asserts that at least one of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func ErrorIsf(err error, target error, msg string, args ...interface{}) bool {
	return ErrorIs(err, target, append([]interface{}{msg}, args...)...)
}

// Eventuallyf asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//	assert.Eventuallyf(func() bool { return true; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
func Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool {
	return Eventually(condition, waitFor, tick, append([]interface{}{msg}, args...)...)
}

// EventuallyWithTf asserts that given condition will be met in waitFor time,
// periodically checking target function each tick. In contrast to Eventually,
// it supplies a CollectT to the condition function, so that the condition
// function can use the CollectT to call other assertions.
// The condition is considered "met" if no errors are raised in a tick.
// The supplied CollectT collects all errors from one tick (if there are any).
// If the condition is not met before waitFor, the collected errors of
// the last tick are copied to t.
//
//	externalValue := false
//	go func() {
//		time.Sleep(8*time.Second)
//		externalValue = true
//	}()
//	assert.EventuallyWithTf(func(c *assert.CollectT, "error message %s", "formatted") {
//		// add assertions as needed; any assertion failure will fail the current tick
//		assert.True(c, externalValue, "expected 'externalValue' to be true")
//	}, 1*time.Second, 10*time.Second, "external state has not changed to 'true'; still false")
func EventuallyWithTf(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool {
	return EventuallyWithT(condition, waitFor, tick, append([]interface{}{msg}, args...)...)
}

// Exactlyf asserts that two objects are equal in value and type.
//
//	assert.Exactlyf(int32(123), int64(123), "error message %s", "formatted")
func Exactlyf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return Exactly(expected, actual, append([]interface{}{msg}, args...)...)
}

// Failf reports a failure through
func Failf(failureMessage string, msg string, args ...interface{}) bool {
	return Fail(failureMessage, append([]interface{}{msg}, args...)...)
}

// FailNowf fails test
func FailNowf(failureMessage string, msg string, args ...interface{}) bool {
	return FailNow(failureMessage, append([]interface{}{msg}, args...)...)
}

// Falsef asserts that the specified value is false.
//
//	assert.Falsef(myBool, "error message %s", "formatted")
func Falsef(value bool, msg string, args ...interface{}) bool {
	return False(value, append([]interface{}{msg}, args...)...)
}

// FileExistsf checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
func FileExistsf(path string, msg string, args ...interface{}) bool {
	return FileExists(path, append([]interface{}{msg}, args...)...)
}

// Greaterf asserts that the first element is greater than the second
//
//	assert.Greaterf(2, 1, "error message %s", "formatted")
//	assert.Greaterf(float64(2), float64(1), "error message %s", "formatted")
//	assert.Greaterf("b", "a", "error message %s", "formatted")
func Greaterf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {
	return Greater(e1, e2, append([]interface{}{msg}, args...)...)
}

// GreaterOrEqualf asserts that the first element is greater than or equal to the second
//
//	assert.GreaterOrEqualf(2, 1, "error message %s", "formatted")
//	assert.GreaterOrEqualf(2, 2, "error message %s", "formatted")
//	assert.GreaterOrEqualf("b", "a", "error message %s", "formatted")
//	assert.GreaterOrEqualf("b", "b", "error message %s", "formatted")
func GreaterOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {
	return GreaterOrEqual(e1, e2, append([]interface{}{msg}, args...)...)
}

// HTTPBodyContainsf asserts that a specified handler returns a
// body that contains a string.
//
//	assert.HTTPBodyContainsf(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) bool {
	return HTTPBodyContains(handler, method, url, values, str, append([]interface{}{msg}, args...)...)
}

// HTTPBodyNotContainsf asserts that a specified handler returns a
// body that does not contain a string.
//
//	assert.HTTPBodyNotContainsf(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) bool {
	return HTTPBodyNotContains(handler, method, url, values, str, append([]interface{}{msg}, args...)...)
}

// HTTPErrorf asserts that a specified handler returns an error status code.
//
//	assert.HTTPErrorf(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool {
	return HTTPError(handler, method, url, values, append([]interface{}{msg}, args...)...)
}

// HTTPRedirectf asserts that a specified handler returns a redirect status code.
//
//	assert.HTTPRedirectf(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool {
	return HTTPRedirect(handler, method, url, values, append([]interface{}{msg}, args...)...)
}

// HTTPStatusCodef asserts that a specified handler returns a specified status code.
//
//	assert.HTTPStatusCodef(myHandler, "GET", "/notImplemented", nil, 501, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{}) bool {
	return HTTPStatusCode(handler, method, url, values, statuscode, append([]interface{}{msg}, args...)...)
}

// HTTPSuccessf asserts that a specified handler returns a success status code.
//
//	assert.HTTPSuccessf(myHandler, "POST", "http://www.google.com", nil, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool {
	return HTTPSuccess(handler, method, url, values, append([]interface{}{msg}, args...)...)
}

// Implementsf asserts that an object is implemented by the specified interface.
//
//	assert.Implementsf((*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
func Implementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{}) bool {
	return Implements(interfaceObject, object, append([]interface{}{msg}, args...)...)
}

// InDeltaf asserts that the two numerals are within delta of each other.
//
//	assert.InDeltaf(math.Pi, 22/7.0, 0.01, "error message %s", "formatted")
func InDeltaf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool {
	return InDelta(expected, actual, delta, append([]interface{}{msg}, args...)...)
}

// InDeltaMapValuesf is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
func InDeltaMapValuesf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool {
	return InDeltaMapValues(expected, actual, delta, append([]interface{}{msg}, args...)...)
}

// InDeltaSlicef is the same as InDelta, except it compares two slices.
func InDeltaSlicef(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool {
	return InDeltaSlice(expected, actual, delta, append([]interface{}{msg}, args...)...)
}

// InEpsilonf asserts that expected and actual have a relative error less than epsilon
func InEpsilonf(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool {
	return InEpsilon(expected, actual, epsilon, append([]interface{}{msg}, args...)...)
}

// InEpsilonSlicef is the same as InEpsilon, except it compares each value from two slices.
func InEpsilonSlicef(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool {
	return InEpsilonSlice(expected, actual, epsilon, append([]interface{}{msg}, args...)...)
}

// IsDecreasingf asserts that the collection is decreasing
//
//	assert.IsDecreasingf([]int{2, 1, 0}, "error message %s", "formatted")
//	assert.IsDecreasingf([]float{2, 1}, "error message %s", "formatted")
//	assert.IsDecreasingf([]string{"b", "a"}, "error message %s", "formatted")
func IsDecreasingf(object interface{}, msg string, args ...interface{}) bool {
	return IsDecreasing(object, append([]interface{}{msg}, args...)...)
}

// IsIncreasingf asserts that the collection is increasing
//
//	assert.IsIncreasingf([]int{1, 2, 3}, "error message %s", "formatted")
//	assert.IsIncreasingf([]float{1, 2}, "error message %s", "formatted")
//	assert.IsIncreasingf([]string{"a", "b"}, "error message %s", "formatted")
func IsIncreasingf(object interface{}, msg string, args ...interface{}) bool {
	return IsIncreasing(object, append([]interface{}{msg}, args...)...)
}

// IsNonDecreasingf asserts that the collection is not decreasing
//
//	assert.IsNonDecreasingf([]int{1, 1, 2}, "error message %s", "formatted")
//	assert.IsNonDecreasingf([]float{1, 2}, "error message %s", "formatted")
//	assert.IsNonDecreasingf([]string{"a", "b"}, "error message %s", "formatted")
func IsNonDecreasingf(object interface{}, msg string, args ...interface{}) bool {
	return IsNonDecreasing(object, append([]interface{}{msg}, args...)...)
}

// IsNonIncreasingf asserts that the collection is not increasing
//
//	assert.IsNonIncreasingf([]int{2, 1, 1}, "error message %s", "formatted")
//	assert.IsNonIncreasingf([]float{2, 1}, "error message %s", "formatted")
//	assert.IsNonIncreasingf([]string{"b", "a"}, "error message %s", "formatted")
func IsNonIncreasingf(object interface{}, msg string, args ...interface{}) bool {
	return IsNonIncreasing(object, append([]interface{}{msg}, args...)...)
}

// IsTypef asserts that the specified objects are of the same type.
func IsTypef(expectedType interface{}, object interface{}, msg string, args ...interface{}) bool {
	return IsType(expectedType, object, append([]interface{}{msg}, args...)...)
}

// JSONEqf asserts that two JSON strings are equivalent.
//
//	assert.JSONEqf(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "error message %s", "formatted")
func JSONEqf(expected string, actual string, msg string, args ...interface{}) bool {
	return JSONEq(expected, actual, append([]interface{}{msg}, args...)...)
}

// Lenf asserts that the specified object has specific length.
// Lenf also fails if the object has a type that len() not accept.
//
//	assert.Lenf(mySlice, 3, "error message %s", "formatted")
func Lenf(object interface{}, length int, msg string, args ...interface{}) bool {
	return Len(object, length, append([]interface{}{msg}, args...)...)
}

// Lessf asserts that the first element is less than the second
//
//	assert.Lessf(1, 2, "error message %s", "formatted")
//	assert.Lessf(float64(1), float64(2), "error message %s", "formatted")
//	assert.Lessf("a", "b", "error message %s", "formatted")
func Lessf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {
	return Less(e1, e2, append([]interface{}{msg}, args...)...)
}

// LessOrEqualf asserts that the first element is less than or equal to the second
//
//	assert.LessOrEqualf(1, 2, "error message %s", "formatted")
//	assert.LessOrEqualf(2, 2, "error message %s", "formatted")
//	assert.LessOrEqualf("a", "b", "error message %s", "formatted")
//	assert.LessOrEqualf("b", "b", "error message %s", "formatted")
func LessOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {
	return LessOrEqual(e1, e2, append([]interface{}{msg}, args...)...)
}

// Negativef asserts that the specified element is negative
//
//	assert.Negativef(-1, "error message %s", "formatted")
//	assert.Negativef(-1.23, "error message %s", "formatted")
func Negativef(e interface{}, msg string, args ...interface{}) bool {
	return Negative(e, append([]interface{}{msg}, args...)...)
}

// Neverf asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//	assert.Neverf(func() bool { return false; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
func Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool {
	return Never(condition, waitFor, tick, append([]interface{}{msg}, args...)...)
}

// Nilf asserts that the specified object is nil.
//
//	assert.Nilf(err, "error message %s", "formatted")
func Nilf(object interface{}, msg string, args ...interface{}) bool {
	return Nil(object, append([]interface{}{msg}, args...)...)
}

// NoDirExistsf checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
func NoDirExistsf(path string, msg string, args ...interface{}) bool {
	return NoDirExists(path, append([]interface{}{msg}, args...)...)
}

// NoErrorf asserts that a function returned no error (i.e. `nil`).
//
//	  actualObj, err := SomeFunction()
//	  if assert.NoErrorf(err, "error message %s", "formatted") {
//		   assert.Equal(expectedObj, actualObj)
//	  }
func NoErrorf(err error, msg string, args ...interface{}) bool {
	return NoError(err, append([]interface{}{msg}, args...)...)
}

// NoFileExistsf checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
func NoFileExistsf(path string, msg string, args ...interface{}) bool {
	return NoFileExists(path, append([]interface{}{msg}, args...)...)
}

// NotContainsf asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//	assert.NotContainsf("Hello World", "Earth", "error message %s", "formatted")
//	assert.NotContainsf(["Hello", "World"], "Earth", "error message %s", "formatted")
//	assert.NotContainsf({"Hello": "World"}, "Earth", "error message %s", "formatted")
func NotContainsf(s interface{}, contains interface{}, msg string, args ...interface{}) bool {
	return NotContains(s, contains, append([]interface{}{msg}, args...)...)
}

// NotEmptyf asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//	if assert.NotEmptyf(obj, "error message %s", "formatted") {
//	  assert.Equal("two", obj[1])
//	}
func NotEmptyf(object interface{}, msg string, args ...interface{}) bool {
	return NotEmpty(object, append([]interface{}{msg}, args...)...)
}

// NotEqualf asserts that the specified values are NOT equal.
//
//	assert.NotEqualf(obj1, obj2, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func NotEqualf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return NotEqual(expected, actual, append([]interface{}{msg}, args...)...)
}

// NotEqualValuesf asserts that two objects are not equal even when converted to the same type
//
//	assert.NotEqualValuesf(obj1, obj2, "error message %s", "formatted")
func NotEqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return NotEqualValues(expected, actual, append([]interface{}{msg}, args...)...)
}

// NotErrorIsf asserts that at none of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func NotErrorIsf(err error, target error, msg string, args ...interface{}) bool {
	return NotErrorIs(err, target, append([]interface{}{msg}, args...)...)
}

// NotNilf asserts that the specified object is not nil.
//
//	assert.NotNilf(err, "error message %s", "formatted")
func NotNilf(object interface{}, msg string, args ...interface{}) bool {
	return NotNil(object, append([]interface{}{msg}, args...)...)
}

// NotPanicsf asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//	assert.NotPanicsf(func(){ RemainCalm() }, "error message %s", "formatted")
func NotPanicsf(f PanicTestFunc, msg string, args ...interface{}) bool {
	return NotPanics(f, append([]interface{}{msg}, args...)...)
}

// NotRegexpf asserts that a specified regexp does not match a string.
//
//	assert.NotRegexpf(regexp.MustCompile("starts"), "it's starting", "error message %s", "formatted")
//	assert.NotRegexpf("^start", "it's not starting", "error message %s", "formatted")
func NotRegexpf(rx interface{}, str interface{}, msg string, args ...interface{}) bool {
	return NotRegexp(rx, str, append([]interface{}{msg}, args...)...)
}

// NotSamef asserts that two pointers do not reference the same object.
//
//	assert.NotSamef(ptr1, ptr2, "error message %s", "formatted")
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func NotSamef(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return NotSame(expected, actual, append([]interface{}{msg}, args...)...)
}

// NotSubsetf asserts that the specified list(array, slice...) or map does NOT
// contain all elements given in the specified subset list(array, slice...) or
// map.
//
//	assert.NotSubsetf([1, 3, 4], [1, 2], "error message %s", "formatted")
//	assert.NotSubsetf({"x": 1, "y": 2}, {"z": 3}, "error message %s", "formatted")
func NotSubsetf(list interface{}, subset interface{}, msg string, args ...interface{}) bool {
	return NotSubset(list, subset, append([]interface{}{msg}, args...)...)
}

// NotZerof asserts that i is not the zero value for its type.
func NotZerof(i interface{}, msg string, args ...interface{}) bool {
	return NotZero(i, append([]interface{}{msg}, args...)...)
}

// Panicsf asserts that the code inside the specified PanicTestFunc panics.
//
//	assert.Panicsf(func(){ GoCrazy() }, "error message %s", "formatted")
func Panicsf(f PanicTestFunc, msg string, args ...interface{}) bool {
	return Panics(f, append([]interface{}{msg}, args...)...)
}

// PanicsWithErrorf asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
//	assert.PanicsWithErrorf("crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
func PanicsWithErrorf(errString string, f PanicTestFunc, msg string, args ...interface{}) bool {
	return PanicsWithError(errString, f, append([]interface{}{msg}, args...)...)
}

// PanicsWithValuef asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//	assert.PanicsWithValuef("crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
func PanicsWithValuef(expected interface{}, f PanicTestFunc, msg string, args ...interface{}) bool {
	return PanicsWithValue(expected, f, append([]interface{}{msg}, args...)...)
}

// Positivef asserts that the specified element is positive
//
//	assert.Positivef(1, "error message %s", "formatted")
//	assert.Positivef(1.23, "error message %s", "formatted")
func Positivef(e interface{}, msg string, args ...interface{}) bool {
	return Positive(e, append([]interface{}{msg}, args...)...)
}

// Regexpf asserts that a specified regexp matches a string.
//
//	assert.Regexpf(regexp.MustCompile("start"), "it's starting", "error message %s", "formatted")
//	assert.Regexpf("start...$", "it's not starting", "error message %s", "formatted")
func Regexpf(rx interface{}, str interface{}, msg string, args ...interface{}) bool {
	return Regexp(rx, str, append([]interface{}{msg}, args...)...)
}

// Samef asserts that two pointers reference the same object.
//
//	assert.Samef(ptr1, ptr2, "error message %s", "formatted")
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func Samef(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {
	return Same(expected, actual, append([]interface{}{msg}, args...)...)
}

// Subsetf asserts that the specified list(array, slice...) or map contains all
// elements given in the specified subset list(array, slice...) or map.
//
//	assert.Subsetf([1, 2, 3], [1, 2], "error message %s", "formatted")
//	assert.Subsetf({"x": 1, "y": 2}, {"x": 1}, "error message %s", "formatted")
func Subsetf(list interface{}, subset interface{}, msg string, args ...interface{}) bool {
	return Subset(list, subset, append([]interface{}{msg}, args...)...)
}

// Truef asserts that the specified value is true.
//
//	assert.Truef(myBool, "error message %s", "formatted")
func Truef(value bool, msg string, args ...interface{}) bool {
	return True(value, append([]interface{}{msg}, args...)...)
}

// WithinDurationf asserts that the two times are within duration delta of each other.
//
//	assert.WithinDurationf(time.Now(), time.Now(), 10*time.Second, "error message %s", "formatted")
func WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) bool {
	return WithinDuration(expected, actual, delta, append([]interface{}{msg}, args...)...)
}

// WithinRangef asserts that a time is within a time range (inclusive).
//
//	assert.WithinRangef(time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second), "error message %s", "formatted")
func WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ...interface{}) bool {
	return WithinRange(actual, start, end, append([]interface{}{msg}, args...)...)
}

// YAMLEqf asserts that two YAML strings are equivalent.
func YAMLEqf(expected string, actual string, msg string, args ...interface{}) bool {
	return YAMLEq(expected, actual, append([]interface{}{msg}, args...)...)
}

// Zerof asserts that i is the zero value for its type.
func Zerof(i interface{}, msg string, args ...interface{}) bool {
	return Zero(i, append([]interface{}{msg}, args...)...)
}

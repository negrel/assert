//go:build !assert

package assert

import (
	"reflect"
	"time"
)

type CompareType int

const (
	compareLess CompareType = iota - 1
	compareEqual
	compareGreater
)

var (
	intType   = reflect.TypeOf(int(1))
	int8Type  = reflect.TypeOf(int8(1))
	int16Type = reflect.TypeOf(int16(1))
	int32Type = reflect.TypeOf(int32(1))
	int64Type = reflect.TypeOf(int64(1))

	uintType   = reflect.TypeOf(uint(1))
	uint8Type  = reflect.TypeOf(uint8(1))
	uint16Type = reflect.TypeOf(uint16(1))
	uint32Type = reflect.TypeOf(uint32(1))
	uint64Type = reflect.TypeOf(uint64(1))

	float32Type = reflect.TypeOf(float32(1))
	float64Type = reflect.TypeOf(float64(1))

	stringType = reflect.TypeOf("")

	timeType  = reflect.TypeOf(time.Time{})
	bytesType = reflect.TypeOf([]byte{})
)

func compare(obj1, obj2 interface{}, kind reflect.Kind) {}

// Greater asserts that the first element is greater than the second
//
//	assert.Greater(2, 1)
//	assert.Greater(float64(2), float64(1))
//	assert.Greater("b", "a")
func Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {}

// GreaterOrEqual asserts that the first element is greater than or equal to the second
//
//	assert.GreaterOrEqual(2, 1)
//	assert.GreaterOrEqual(2, 2)
//	assert.GreaterOrEqual("b", "a")
//	assert.GreaterOrEqual("b", "b")
func GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {}

// Less asserts that the first element is less than the second
//
//	assert.Less(1, 2)
//	assert.Less(float64(1), float64(2))
//	assert.Less("a", "b")
func Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {}

// LessOrEqual asserts that the first element is less than or equal to the second
//
//	assert.LessOrEqual(1, 2)
//	assert.LessOrEqual(2, 2)
//	assert.LessOrEqual("a", "b")
//	assert.LessOrEqual("b", "b")
func LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {}

// Positive asserts that the specified element is positive
//
//	assert.Positive(1)
//	assert.Positive(1.23)
func Positive(e interface{}, msgAndArgs ...interface{}) {}

// Negative asserts that the specified element is negative
//
//	assert.Negative(-1)
//	assert.Negative(-1.23)
func Negative(e interface{}, msgAndArgs ...interface{}) {}

func compareTwoValues(e1 interface{}, e2 interface{}, allowedComparesResults []CompareType, failMessage string, msgAndArgs ...interface{}) {
}

func containsValue(values []CompareType, value CompareType) {}

//go:build !assert

package assert

// isOrdered checks that collection contains orderable elements.
func isOrdered(object interface{}, allowedComparesResults []CompareType, failMessage string, msgAndArgs ...interface{}) {
}

// IsIncreasing asserts that the collection is increasing
//
//	assert.IsIncreasing([]int{1, 2, 3})
//	assert.IsIncreasing([]float{1, 2})
//	assert.IsIncreasing([]string{"a", "b"})
func IsIncreasing(object interface{}, msgAndArgs ...interface{}) {}

// IsNonIncreasing asserts that the collection is not increasing
//
//	assert.IsNonIncreasing([]int{2, 1, 1})
//	assert.IsNonIncreasing([]float{2, 1})
//	assert.IsNonIncreasing([]string{"b", "a"})
func IsNonIncreasing(object interface{}, msgAndArgs ...interface{}) {}

// IsDecreasing asserts that the collection is decreasing
//
//	assert.IsDecreasing([]int{2, 1, 0})
//	assert.IsDecreasing([]float{2, 1})
//	assert.IsDecreasing([]string{"b", "a"})
func IsDecreasing(object interface{}, msgAndArgs ...interface{}) {}

// IsNonDecreasing asserts that the collection is not decreasing
//
//	assert.IsNonDecreasing([]int{1, 1, 2})
//	assert.IsNonDecreasing([]float{1, 2})
//	assert.IsNonDecreasing([]string{"a", "b"})
func IsNonDecreasing(object interface{}, msgAndArgs ...interface{}) {}

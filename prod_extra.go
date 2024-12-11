//go:build !assert

package assert

import "sync"

// A TryLocker represents an object that can attempt to acquire a lock and report
// whether it succeeded.
//
// [sync.Mutex] and [sync.RWMutex] implements this interface.
type TryLocker interface {
	sync.Locker
	TryLock() bool
}

// Locked asserts that the [TryLocker] is already locked.
//
//	assert.IsDecreasing([]int{2, 1, 0})
//	assert.IsDecreasing([]float{2, 1})
//	assert.IsDecreasing([]string{"b", "a"})
func Locked(locker TryLocker, msgAndArgs ...any) {}

// Lockedf asserts that the [TryLocker] is already locked.
//
//	assert.IsDecreasing([]int{2, 1, 0})
//	assert.IsDecreasing([]float{2, 1})
//	assert.IsDecreasing([]string{"b", "a"})
func Lockedf(locker TryLocker, msg string, args ...any) {}

// Unlocked asserts that the [TryLocker] is unlocked.
//
//	assert.IsDecreasing([]int{2, 1, 0})
//	assert.IsDecreasing([]float{2, 1})
//	assert.IsDecreasing([]string{"b", "a"})
func Unlocked(locker TryLocker, msgAndArgs ...any) {}

// UnLockedf asserts that the [TryLocker] is unlocked.
//
//	assert.IsDecreasing([]int{2, 1, 0})
//	assert.IsDecreasing([]float{2, 1})
//	assert.IsDecreasing([]string{"b", "a"})
func Unlockedf(locker TryLocker, msg string, args ...any) {}

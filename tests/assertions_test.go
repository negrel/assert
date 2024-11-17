package assert

import (
	"errors"
	"io"
	"regexp"
	"sync"
	"testing"

	"github.com/negrel/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSliceIndexWithoutBoundCheckAssertions(b *testing.B) {
	get := func(slice []string, index int) string {
		return slice[index]
	}
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 0; i < b.N; i++ {
		_ = get(days, i%len(days))
	}
}

func BenchmarkSliceIndexWithBoundCheckAssertions(b *testing.B) {
	get := func(slice []string, index int) string {
		assert.GreaterOrEqual(index, 0)
		assert.Less(index, len(slice))
		return slice[index]
	}
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 0; i < b.N; i++ {
		_ = get(days, i%len(days))
	}
}

func TestAssertCondition(t *testing.T) {
	t.Run("ReturnsTrueOk", func(t *testing.T) {
		assert.Condition(func() bool {
			return true
		})
	})

	t.Run("ReturnsFalsePanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Condition(func() bool {
				return false
			})
		})
	})
}

func TestAssertContains(t *testing.T) {
	t.Run("Substring", func(t *testing.T) {
		t.Run("PresentOk", func(t *testing.T) {
			assert.Contains("Hello World", "World")
		})
		t.Run("AbsentPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.Contains("Hello World", "bye")
			})
		})
	})

	t.Run("SliceElement", func(t *testing.T) {
		t.Run("PresentOk", func(t *testing.T) {
			assert.Contains([]string{"Hello", "World"}, "World")
		})
		t.Run("AbsentPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.Contains([]string{"Hello", "World"}, "bye")
			})
		})
	})

	t.Run("MapKey", func(t *testing.T) {
		t.Run("PresentOk", func(t *testing.T) {
			assert.Contains(map[string]string{"Hello": "World"}, "Hello")
		})
		t.Run("AbsentPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.Contains(map[string]string{"Hello": "World"}, "bye")
			})
		})
	})
}

func TestAssertDirExists(t *testing.T) {
	t.Run("ExistOk", func(t *testing.T) {
		assert.DirExists(".")
	})

	t.Run("DoesNotExistPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.DirExists("./nonexistentdir")
		})
	})
}

func TestAssertElementsMatch(t *testing.T) {
	t.Run("Array", func(t *testing.T) {
		t.Run("MatchOk", func(t *testing.T) {
			assert.ElementsMatch([3]int{1, 2, 3}, [3]int{3, 2, 1})
		})
		t.Run("NoMatchPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.ElementsMatch([3]int{4, 5, 6}, [3]int{3, 2, 1})
			})
		})
	})

	t.Run("Slice", func(t *testing.T) {
		t.Run("MatchOk", func(t *testing.T) {
			assert.ElementsMatch([]int{1, 2, 3}, []int{3, 2, 1})
		})
		t.Run("NoMatchPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.ElementsMatch([]int{4, 5, 6}, []int{3, 2, 1})
			})
		})
	})
}

func TestAssertEmpty(t *testing.T) {
	t.Run("Slice", func(t *testing.T) {
		t.Run("EmptyOk", func(t *testing.T) {
			assert.Empty([]string{})
		})
		t.Run("NotEmptyPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.Empty([]string{""})
			})
		})
	})

	t.Run("Ptr", func(t *testing.T) {
		t.Run("EmptyOk", func(t *testing.T) {
			assert.Empty(nil)
		})
		t.Run("NotEmptyPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.Empty(t)
			})
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run("EmptyOk", func(t *testing.T) {
			assert.Empty("")
		})
		t.Run("NotEmptyPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.Empty("Hello world")
			})
		})
	})

	t.Run("Bool", func(t *testing.T) {
		t.Run("EmptyOk", func(t *testing.T) {
			assert.Empty(false)
		})
		t.Run("NotEmptyPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.Empty(true)
			})
		})
	})
}

func TestAssertEqual(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		t.Run("EqualOk", func(t *testing.T) {
			assert.Equal(123, 123)
		})
		t.Run("NotEqualPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.Equal(123, 456)
			})
		})
	})

	t.Run("Bool", func(t *testing.T) {
		t.Run("EqualOk", func(t *testing.T) {
			assert.Equal(false, false)
		})
		t.Run("NotEqualPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.Equal(true, false)
			})
		})
	})
}

func TestAssertEqualError(t *testing.T) {
	t.Run("EqualOk", func(t *testing.T) {
		assert.EqualError(errors.New("foo"), "foo")
	})
	t.Run("NotEqualPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.EqualError(errors.New("foo"), "bar")
		})
	})
}

func TestAssertEqualExportedValue(t *testing.T) {
	type S struct {
		private int
		Public  bool
	}

	t.Run("EqualOk", func(t *testing.T) {
		expected := S{-1, true}
		actual := S{-3, true}
		assert.EqualExportedValues(expected, actual)
	})
	t.Run("NotEqualPanics", func(t *testing.T) {
		require.Panics(t, func() {
			expected := S{-1, true}
			actual := S{-1, false}
			assert.EqualExportedValues(expected, actual)
		})
	})
}

func TestAssertEqualValues(t *testing.T) {
	t.Run("EqualOk", func(t *testing.T) {
		assert.EqualValues(uint32(123), uint64(123))
	})
	t.Run("NotEqualPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.EqualValues(uint32(123), uint32(456))
		})
	})
}

func TestAssertError(t *testing.T) {
	t.Run("ErrorOk", func(t *testing.T) {
		assert.Error(errors.New("foo"))
	})
	t.Run("NoErrorPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Error(nil)
		})
	})
}

func TestAssertErrorAs(t *testing.T) {
	t.Run("MatchOk", func(t *testing.T) {
		assert.ErrorAs(errors.Join(errors.New("foo"), io.EOF), &io.EOF)
	})
	t.Run("NoMatchOk", func(t *testing.T) {
		// ErrorAs never panics
		// require.Panics(t, func() {
		assert.ErrorAs(errors.Join(errors.New("foo"), errors.New("bar")), &io.EOF)
		// })
	})
}

func TestAssertErrorContains(t *testing.T) {
	t.Run("ContainsOk", func(t *testing.T) {
		assert.ErrorContains(errors.New("foobarqux"), "qux")
	})
	t.Run("DoesNotContainPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.ErrorContains(errors.New("foobarqux"), "baz")
		})
	})
}

func TestAssertErrorIs(t *testing.T) {
	t.Run("MatchOk", func(t *testing.T) {
		assert.ErrorIs(errors.Join(errors.New("foo"), io.EOF), io.EOF)
	})
	t.Run("DoesNotMatchPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.ErrorIs(errors.Join(errors.New("foo"), errors.New("bar")), io.EOF)
		})
	})
}

func TestAssertExactly(t *testing.T) {
	t.Run("EqualOk", func(t *testing.T) {
		assert.Exactly(uint32(123), uint32(123))
	})
	t.Run("NotEqualPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Exactly(uint32(123), uint64(123))
		})
	})
}

func TestAssertFalse(t *testing.T) {
	t.Run("TrueValuePanic", func(t *testing.T) {
		require.Panics(t, func() {
			assert.False(true)
		})
	})

	t.Run("FalseValueOk", func(t *testing.T) {
		assert.False(false)
	})
}

func TestAssertFileExists(t *testing.T) {
	t.Run("ExistsOk", func(t *testing.T) {
		assert.FileExists("assertions_test.go")
	})

	t.Run("DoesNotExistPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.FileExists("nonexistentfile")
		})
	})
}

func TestAssertGreater(t *testing.T) {
	t.Run("GreaterOk", func(t *testing.T) {
		assert.Greater(2, 1)
	})

	t.Run("EqualPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Greater(2, 2)
		})
	})

	t.Run("LessPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Greater(1, 2)
		})
	})
}

func TestAssertGreaterOrEqual(t *testing.T) {
	t.Run("GreaterOk", func(t *testing.T) {
		assert.GreaterOrEqual(2, 1)
	})

	t.Run("EqualOk", func(t *testing.T) {
		assert.GreaterOrEqual(2, 2)
	})

	t.Run("LessPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.GreaterOrEqual(1, 2)
		})
	})
}

func TestAssertImplements(t *testing.T) {
	t.Run("ImplementsOk", func(t *testing.T) {
		assert.Implements((*require.TestingT)(nil), t)
	})

	t.Run("DoesNotImplementPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Implements((interface{ abc() })(nil), t)
		})
	})
}

func TestAssertInDelta(t *testing.T) {
	t.Run("WithinOk", func(t *testing.T) {
		assert.InDelta(1, 3, 2)
	})

	t.Run("NotWithinPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.InDelta(1, 3, 0)
		})
	})
}

func TestAssertIsDecreasing(t *testing.T) {
	t.Run("DecreasingOk", func(t *testing.T) {
		assert.IsDecreasing([]int{3, 2, 1})
	})

	t.Run("NotDecreasingPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.IsDecreasing([]int{3, 2, 3})
		})
	})
}

func TestAssertIsIncreasing(t *testing.T) {
	t.Run("IncreasingOk", func(t *testing.T) {
		assert.IsIncreasing([]int{1, 2, 3})
	})

	t.Run("NotIncreasingPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.IsIncreasing([]int{1, 2, 1})
		})
	})
}

func TestAssertIsNonDecreasing(t *testing.T) {
	t.Run("NotDecreasingOk", func(t *testing.T) {
		assert.IsNonDecreasing([]int{1, 2, 3})
	})

	t.Run("DecreasingPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.IsNonDecreasing([]int{1, 2, 1})
		})
	})
}

func TestAssertIsNonIncreasing(t *testing.T) {
	t.Run("NonIncreasingOk", func(t *testing.T) {
		assert.IsNonIncreasing([]int{3, 2, 1})
	})

	t.Run("DecreasingPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.IsNonIncreasing([]int{3, 2, 3})
		})
	})
}

func TestAssertIsType(t *testing.T) {
	t.Run("MatchOk", func(t *testing.T) {
		assert.IsType(int32(0), int32(0))
	})

	t.Run("NotMatchPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.IsType(uint32(0), int32(0))
		})
	})
}

func TestAssertJSONEq(t *testing.T) {
	t.Run("EqualOk", func(t *testing.T) {
		assert.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	})

	t.Run("NotMatchPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.JSONEq(`{"bye": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
		})
	})
}

func TestAssertLen(t *testing.T) {
	t.Run("MatchOk", func(t *testing.T) {
		assert.Len([]int{1, 2, 3}, 3)
	})

	t.Run("NotMatchPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Len([]int{1, 2, 3}, 0)
		})
	})
}

func TestAssertLess(t *testing.T) {
	t.Run("LessOk", func(t *testing.T) {
		assert.Less(1, 2)
	})

	t.Run("EqualPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Less(2, 2)
		})
	})

	t.Run("GreaterPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Less(2, 1)
		})
	})
}

func TestAssertLessOrEqual(t *testing.T) {
	t.Run("LessOk", func(t *testing.T) {
		assert.LessOrEqual(1, 2)
	})

	t.Run("EqualOk", func(t *testing.T) {
		assert.LessOrEqual(2, 2)
	})

	t.Run("GreaterPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.LessOrEqual(2, 1)
		})
	})
}

func TestAssertNegative(t *testing.T) {
	t.Run("NegativeOk", func(t *testing.T) {
		assert.Negative(-1)
	})

	t.Run("PositivePanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Negative(2)
		})
	})
}

func TestAssertNil(t *testing.T) {
	t.Run("NilOk", func(t *testing.T) {
		assert.Nil(nil)
	})

	t.Run("NotNilPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Nil(t)
		})
	})
}

func TestAssertNoDirExists(t *testing.T) {
	t.Run("ExistPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.NoDirExists(".")
		})
	})

	t.Run("DoesNotExistOk", func(t *testing.T) {
		assert.NoDirExists("./nonexistentdir")
	})
}

func TestAssertNoError(t *testing.T) {
	t.Run("NoErrorOk", func(t *testing.T) {
		require.Panics(t, func() {
			assert.NoError(errors.New("foo"))
		})
	})

	t.Run("ErrorPanics", func(t *testing.T) {
		assert.NoError(nil)
	})
}

func TestAssertNoFileExists(t *testing.T) {
	t.Run("ExistsPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.NoFileExists("assertions_test.go")
		})
	})

	t.Run("DoesNotExistOk", func(t *testing.T) {
		assert.NoFileExists("nonexistentfile")
	})
}

func TestAssertNotContains(t *testing.T) {
	t.Run("Substring", func(t *testing.T) {
		t.Run("PresentPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.NotContains("Hello World", "World")
			})
		})
		t.Run("AbsentOk", func(t *testing.T) {
			assert.NotContains("Hello World", "bye")
		})
	})

	t.Run("SliceElement", func(t *testing.T) {
		t.Run("PresentPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.NotContains([]string{"Hello", "World"}, "World")
			})
		})
		t.Run("AbsentOk", func(t *testing.T) {
			assert.NotContains([]string{"Hello", "World"}, "bye")
		})
	})

	t.Run("MapKey", func(t *testing.T) {
		t.Run("PresentPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.NotContains(map[string]string{"Hello": "World"}, "Hello")
			})
		})
		t.Run("AbsentOk", func(t *testing.T) {
			assert.NotContains(map[string]string{"Hello": "World"}, "bye")
		})
	})
}

func TestAssertNotEmpty(t *testing.T) {
	t.Run("Slice", func(t *testing.T) {
		t.Run("EmptyPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.NotEmpty([]string{})
			})
		})
		t.Run("NotEmptyOk", func(t *testing.T) {
			assert.NotEmpty([]string{""})
		})
	})

	t.Run("Ptr", func(t *testing.T) {
		t.Run("EmptyPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.NotEmpty(nil)
			})
		})
		t.Run("NotEmptyOk", func(t *testing.T) {
			assert.NotEmpty(t)
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run("EmptyPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.NotEmpty("")
			})
		})
		t.Run("NotEmptyOk", func(t *testing.T) {
			assert.NotEmpty("Hello world")
		})
	})

	t.Run("Bool", func(t *testing.T) {
		t.Run("EmptyPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.NotEmpty(false)
			})
		})
		t.Run("NotEmptyOk", func(t *testing.T) {
			assert.NotEmpty(true)
		})
	})
}

func TestAssertNotEqual(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		t.Run("EqualPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.NotEqual(123, 123)
			})
		})
		t.Run("NotEqualOk", func(t *testing.T) {
			assert.NotEqual(123, 456)
		})
	})

	t.Run("Bool", func(t *testing.T) {
		t.Run("EqualPanics", func(t *testing.T) {
			require.Panics(t, func() {
				assert.NotEqual(false, false)
			})
		})
		t.Run("NotEqualOk", func(t *testing.T) {
			assert.NotEqual(true, false)
		})
	})
}

func TestAssertNotEqualValues(t *testing.T) {
	t.Run("EqualPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.NotEqualValues(uint32(123), uint64(123))
		})
	})
	t.Run("NotEqualOk", func(t *testing.T) {
		assert.NotEqualValues(uint32(123), uint32(456))
	})
}

func TestAssertNotErrorIs(t *testing.T) {
	t.Run("MatchPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.NotErrorIs(errors.Join(errors.New("foo"), io.EOF), io.EOF)
		})
	})
	t.Run("DoesNotMatchOk", func(t *testing.T) {
		assert.NotErrorIs(errors.Join(errors.New("foo"), errors.New("bar")), io.EOF)
	})
}

func TestAssertNotNil(t *testing.T) {
	t.Run("NilPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.NotNil(nil)
		})
	})

	t.Run("NotNilOk", func(t *testing.T) {
		assert.NotNil(t)
	})
}

func TestAssertNotSame(t *testing.T) {
	t.Run("SamePanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.NotSame(t, t)
		})
	})

	t.Run("NotSameOk", func(t *testing.T) {
		assert.NotSame(t, nil)
	})
}

func TestAssertNotSubset(t *testing.T) {
	t.Run("SubsetPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.NotSubset([]int{1, 3, 4}, []int{1, 3})
		})
	})

	t.Run("NotSubsetOk", func(t *testing.T) {
		assert.NotSubset([]int{1, 3, 4}, []int{1, 2})
	})
}

func TestAssertNotZero(t *testing.T) {
	t.Run("ZeroPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.NotZero(false)
		})
	})

	t.Run("NotZeroOk", func(t *testing.T) {
		assert.NotZero(true)
	})
}

func TestAssertPositive(t *testing.T) {
	t.Run("PositiveOk", func(t *testing.T) {
		assert.Positive(1)
	})

	t.Run("NegativePanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Positive(-2)
		})
	})
}

func TestAssertRegexp(t *testing.T) {
	t.Run("MatchOk", func(t *testing.T) {
		assert.Regexp(regexp.MustCompile(`\d+`), "123")
	})

	t.Run("NoMatchPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Regexp(regexp.MustCompile(`\d+`), "abc")
		})
	})
}

func TestAssertSame(t *testing.T) {
	t.Run("SameOk", func(t *testing.T) {
		assert.Same(t, t)
	})

	t.Run("NotSamePanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Same(t, nil)
		})
	})
}

func TestAssertSubset(t *testing.T) {
	t.Run("SubsetOk", func(t *testing.T) {
		assert.Subset([]int{1, 3, 4}, []int{1, 3})
	})

	t.Run("NotSubsetPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Subset([]int{1, 3, 4}, []int{1, 2})
		})
	})
}

func TestAssertTrue(t *testing.T) {
	t.Run("FalseValuePanic", func(t *testing.T) {
		require.Panics(t, func() {
			assert.True(false)
		})
	})

	t.Run("TrueValueOk", func(t *testing.T) {
		assert.True(true)
	})
}

func TestAssertYAMLEq(t *testing.T) {
	t.Run("EqualOk", func(t *testing.T) {
		assert.YAMLEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	})

	t.Run("NotMatchPanics", func(t *testing.T) {
		require.Panics(t, func() {
			assert.YAMLEq(`{"bye": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
		})
	})
}

func TestAssertZero(t *testing.T) {
	t.Run("ZeroOk", func(t *testing.T) {
		assert.Zero(false)
	})

	t.Run("NotZeroOk", func(t *testing.T) {
		require.Panics(t, func() {
			assert.Zero(true)
		})
	})
}

func TestAssertLocked(t *testing.T) {
	t.Run("Locked", func(t *testing.T) {
		var mu sync.Mutex
		mu.Lock()
		assert.Locked(&mu)
	})

	t.Run("Unlocked", func(t *testing.T) {
		var mu sync.Mutex
		require.Panics(t, func() {
			assert.Locked(&mu)
		})
	})
}

func TestAssertUnlocked(t *testing.T) {
	t.Run("Unlocked", func(t *testing.T) {
		var mu sync.Mutex
		assert.Unlocked(&mu)
	})

	t.Run("Locked", func(t *testing.T) {
		var mu sync.Mutex
		mu.Lock()
		require.Panics(t, func() {
			assert.Unlocked(&mu)
		})
	})
}

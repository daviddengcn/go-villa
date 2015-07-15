package villa

import (
	"fmt"
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestIntSlice(t *testing.T) {
	defer __(o_(t))

	var s IntSlice
	for i := 0; i < 1000; i++ {
		s.Add(i)
	} // for i

	assert.Equal(t, "len", len(s), 1000)
	//fmt.Println(s)
	s.Clear()
	assert.Equal(t, "len", len(s), 0)

	s = IntSlice{}
	s.Add(1)
	s.Insert(0, 2)
	s.Insert(1, 3)
	t.Logf("%v", s)
	assert.Equal(t, "len", len(s), 3)
	assert.StringEqual(t, "s", s, "[2 3 1]")
}

func TestIntSliceRemove(t *testing.T) {
	defer __(o_(t))

	var s IntSlice
	s.Add(1, 2, 3, 4, 5, 6, 7)
	assert.Equal(t, "len", len(s), 7)
	assert.StringEqual(t, "s", s, "[1 2 3 4 5 6 7]")

	s.Fill(2, 5, 9)
	assert.StringEqual(t, "s", s, "[1 2 9 9 9 6 7]")

	s.RemoveRange(2, 5)
	assert.Equal(t, "len", len(s), 4)
	assert.StringEqual(t, "s", s, "[1 2 6 7]")

	s.Remove(2)
	assert.Equal(t, "len", len(s), 3)
	assert.StringEqual(t, "s", s, "[1 2 7]")
}

func TestIntSliceEquals(t *testing.T) {
	s := IntSlice([]int{1, 2, 3, 4})

	assert.Equal(t, "s.Equals(nil)", s.Equals(nil), false)
	assert.Equal(t, "s.Equals([1, 2, 3, 4])", s.Equals([]int{1, 2, 3, 4}), true)
	assert.Equal(t, "s.Equals([1, 2, 5, 4])", s.Equals([]int{1, 2, 5, 4}), false)
	assert.Equal(t, "s.Equals([1, 2, 3, 4, 5])", s.Equals([]int{1, 2, 3, 4, 5}), false)

	assert.Equal(t, "nil.Equals([]int{})", IntSlice(nil).Equals(s[:0]), true)
	assert.Equal(t, "nil.Equals([]int{1, 2})", IntSlice(nil).Equals([]int{1, 2}), false)
}

func ExampleIntSlice_direct() {
	var s IntSlice
	s.Add(10, 20, 30)
	fmt.Println(s)
	s.Insert(1, 40, 50)
	fmt.Println(s)
	s.Swap(1, len(s)-1)
	fmt.Println(s)
	s.RemoveRange(1, 3)
	fmt.Println(s)
	s.Fill(0, len(s), 55)
	fmt.Println(s)
	s.Clear()
	fmt.Println(s)
	// Output:
	// [10 20 30]
	// [10 40 50 20 30]
	// [10 30 50 20 40]
	// [10 20 40]
	// [55 55 55]
	// []
}
func ExampleIntSlice_typecnv() {
	var s []int
	s = append(s, 10, 20, 30)
	fmt.Println(s)
	(*IntSlice)(&s).Insert(1, 40, 50)
	fmt.Println(s)
	IntSlice(s).Swap(1, len(s)-1)
	fmt.Println(s)
	(*IntSlice)(&s).RemoveRange(1, 3)
	fmt.Println(s)
	IntSlice(s).Fill(0, len(s), 55)
	fmt.Println(s)
	s = s[:0]
	fmt.Println(s)
	// Output:
	// [10 20 30]
	// [10 40 50 20 30]
	// [10 30 50 20 40]
	// [10 20 40]
	// [55 55 55]
	// []
}

func BenchmarkIntSliceInsert(b *testing.B) {
	b.StopTimer()
	s := make(IntSlice, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s.Insert(1, i)
	}
}

func BenchmarkIntSliceInsertByAppend(b *testing.B) {
	b.StopTimer()
	s := make([]int, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s = append(s[:1], append([]int{i}, s[1:]...)...)
	}
}

func BenchmarkIntSliceInsertByCopy(b *testing.B) {
	b.StopTimer()
	s := make([]int, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s = append(s, 0)
		copy(s[2:], s[1:])
		s[1] = i
	}
}

package villa

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/golangplus/testing/assert"
)

func o_(t *testing.T) (string, *testing.T) {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	if p := strings.LastIndexAny(name, `./\`); p >= 0 {
		name = name[p+1:]
	} // if
	t.Logf("== BEGIN %s ==", name)
	return name, t
}

func __(name string, t *testing.T) {
	t.Logf("== END %s ==", name)
}

var intInterfaceCmpFunc = CmpFunc(func(e1, e2 interface{}) int {
	if e1.(int) < e2.(int) {
		return -1
	} else if e1.(int) > e2.(int) {
		return 1
	}
	return 0
})

func TestSlice(t *testing.T) {
	defer __(o_(t))

	var s Slice
	for i := 0; i < 1000; i++ {
		s.Add(i)
	}

	assert.Equal(t, "len(s)", len(s), 1000)
	s.Clear()
	assert.Equal(t, "len(s)", len(s), 0)

	s = Slice{}
	s.Add(4, 1)
	s.Insert(1, 2, 3)
	assert.Equal(t, "len(s)", len(s), 4)
	assert.StringEqual(t, "s", s, "[4 2 3 1]")
}

func ExampleSlice_direct() {
	type A struct {
		B, C int
	}

	var s Slice
	s.Add(10, "20", 30)
	fmt.Println(s)
	s.InsertSlice(len(s), []A{A{50, 60}, A{70, 80}})
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
	// [10 20 30 {50 60} {70 80}]
	// [10 40 50 20 30 {50 60} {70 80}]
	// [10 {70 80} 50 20 30 {50 60} 40]
	// [10 20 30 {50 60} 40]
	// [55 55 55 55 55]
	// []
}

func ExampleSlice_typecnv() {
	type A struct {
		B, C int
	}

	var s []interface{}
	s = append(s, 10, "20", 30)
	fmt.Println(s)
	(*Slice)(&s).InsertSlice(len(s), []A{A{50, 60}, A{70, 80}})
	fmt.Println(s)
	(*Slice)(&s).Insert(1, 40, 50)
	fmt.Println(s)
	Slice(s).Swap(1, len(s)-1)
	fmt.Println(s)
	(*Slice)(&s).RemoveRange(1, 3)
	fmt.Println(s)
	Slice(s).Fill(0, len(s), 55)
	fmt.Println(s)
	s = s[:0]
	fmt.Println(s)
	// Output:
	// [10 20 30]
	// [10 20 30 {50 60} {70 80}]
	// [10 40 50 20 30 {50 60} {70 80}]
	// [10 {70 80} 50 20 30 {50 60} 40]
	// [10 20 30 {50 60} 40]
	// [55 55 55 55 55]
	// []
}

func TestSliceRemove(t *testing.T) {
	defer __(o_(t))

	var s Slice
	s.Add(1, 2, 3, 4, 5, 6, 7)
	assert.Equal(t, "len(s)", len(s), 7)
	assert.StringEqual(t, "s", s, "[1 2 3 4 5 6 7]")

	s.RemoveRange(2, 5)
	assert.Equal(t, "len(s)", len(s), 4)
	assert.StringEqual(t, "s", s, "[1 2 6 7]")

	s.Remove(2)
	assert.Equal(t, "len(s)", len(s), 3)
	assert.StringEqual(t, "s", s, "[1 2 7]")
}

func TestSliceEquals(t *testing.T) {
	s := Slice{1, 2, 3, 4}

	assert.Equal(t, "s.Equals(nil)", s.Equals(nil), false)
	assert.Equal(t, "s.Equals([1, 2, 3, 4])", s.Equals([]interface{}{1, 2, 3, 4}), true)
	assert.Equal(t, "s.Equals([1, 2, 5, 4])", s.Equals([]interface{}{1, 2, 5, 4}), false)
	assert.Equal(t, "s.Equals([1, 2, 3, 4, 5])", s.Equals([]interface{}{1, 2, 3, 4, 5}), false)

	assert.Equal(t, "nil.Equals([]int{})", Slice(nil).Equals(s[:0]), true)
	assert.Equal(t, "nil.Equals([]int{1, 2})", Slice(nil).Equals([]interface{}{1, 2}), false)
}

func BenchmarkSliceInsert(b *testing.B) {
	b.StopTimer()
	s := make(Slice, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s.Insert(1, i)
	}
}

func BenchmarkSliceInsertByAppend(b *testing.B) {
	b.StopTimer()
	s := make([]int, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s = append(s[:1], append([]int{i}, s[1:]...)...)
	}
}

func BenchmarkSliceInsertByCopy(b *testing.B) {
	b.StopTimer()
	s := make([]int, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s = append(s, 0)
		copy(s[2:], s[1:])
		s[1] = i
	}
}

func TestN(t *testing.T) {
	var s Slice
	for i := range N[:10] {
		s.Add(i)
	}
	t.Logf("%v", s)
	assert.Equal(t, "Equals", s.Equals([]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), true)
}

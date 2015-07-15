package villa

import (
	"math/cmplx"
	"testing"

	"github.com/golangplus/testing/assert"
)

var cmplexAbsCmpFunc = ComplexCmpFunc(func(a, b complex128) int {
	absA, absB := cmplx.Abs(a), cmplx.Abs(b)
	if absA < absB {
		return -1
	} else if absA > absB {
		return 1
	} // else if

	// The following implementation defines a total order comparator which returns 0 iff. a == b.
	// A simpler alternative could simply return 0 when there modulus' equals.

	if real(a) < real(b) {
		return -1
	} else if real(a) > real(b) {
		return 1
	} // else if

	if imag(a) < imag(b) {
		return -1
	} else if imag(a) > imag(b) {
		return 1
	} // else if

	return 0
})

func TestComplexSlice(t *testing.T) {
	defer __(o_(t))

	var s ComplexSlice
	for i := 0; i < 1000; i++ {
		s.Add(complex(float64(i), float64(-i)))
	} // for i

	assert.Equal(t, "len", len(s), 1000)
	//fmt.Println(s)
	s.Clear()
	assert.Equal(t, "len", len(s), 0)

	s = ComplexSlice{}
	s.Add(-4, -1)
	s.Insert(1, -2, -3)
	t.Logf("%v", s)
	assert.Equal(t, "len", len(s), 4)
	assert.StringEqual(t, "s", s, "[(-4+0i) (-2+0i) (-3+0i) (-1+0i)]")
}

func TestComplexSliceRemove(t *testing.T) {
	defer __(o_(t))

	var s ComplexSlice
	s.Add(-1, -2, -3, -4, -5, -6, -7)
	assert.Equal(t, "len", len(s), 7)
	assert.StringEqual(t, "s", s, "[(-1+0i) (-2+0i) (-3+0i) (-4+0i) (-5+0i) (-6+0i) (-7+0i)]")

	s.Fill(2, 5, -9-8i)
	assert.StringEqual(t, "s", s, "[(-1+0i) (-2+0i) (-9-8i) (-9-8i) (-9-8i) (-6+0i) (-7+0i)]")

	s.RemoveRange(2, 5)
	assert.Equal(t, "len", len(s), 4)
	assert.StringEqual(t, "s", s, "[(-1+0i) (-2+0i) (-6+0i) (-7+0i)]")

	s.Remove(2)
	assert.Equal(t, "len", len(s), 3)
	assert.StringEqual(t, "s", s, "[(-1+0i) (-2+0i) (-7+0i)]")
}

func TestComplexSliceEquals(t *testing.T) {
	s := ComplexSlice([]complex128{1, 2, 3, 4})

	assert.Equal(t, "s.Equals(nil)", s.Equals(nil, 100), false)
	assert.Equal(t, "s.Equals([1, 2, 3, 4])", s.Equals([]complex128{1, 2, 3, 4}, 1e-5), true)
	assert.Equal(t, "s.Equals([1, 2, 5, 4])", s.Equals([]complex128{1, 2, 5, 4}, 1e-5), false)
	assert.Equal(t, "s.Equals([1, 2, 5, 4])", s.Equals([]complex128{1, 2, 5, 4}, 10), true)
	assert.Equal(t, "s.Equals([1, 2, 3, 4, 5])", s.Equals([]complex128{1, 2, 3, 4, 5}, 1e-5), false)

	assert.Equal(t, "nil.Equals([]float64{})", ComplexSlice(nil).Equals(s[:0], 100), true)
	assert.Equal(t, "nil.Equals([]float64{1, 2})", ComplexSlice(nil).Equals([]complex128{1, 2}, 1e-5), false)
}

func BenchmarkComplexSliceInsert(b *testing.B) {
	b.StopTimer()
	s := make(ComplexSlice, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s.Insert(1, complex(float64(i), float64(-i)))
		//s.RemoveRange(1, 1+len(a)*2)
	} // for i
}

func BenchmarkComplexSliceInsertByAppend(b *testing.B) {
	b.StopTimer()
	s := make([]complex128, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s = append(s[:1], append([]complex128{complex(float64(i), float64(-i))}, s[1:]...)...)
	} // for i
}

func BenchmarkComplexSliceInsertByCopy(b *testing.B) {
	b.StopTimer()
	s := make([]complex128, 100000, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s = append(s, 0)
		copy(s[2:], s[1:])
		s[1] = complex(float64(i), float64(-i))
	} // for i
}

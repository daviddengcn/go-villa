package villa

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestFloatSlice(t *testing.T) {
	defer __(o_(t))

	var s FloatSlice
	for i := 0; i < 1000; i++ {
		s.Add(float64(i))
	} // for i

	assert.Equal(t, "len", len(s), 1000)
	//fmt.Println(s)
	s.Clear()
	assert.Equal(t, "len", len(s), 0)

	s = FloatSlice{}
	s.Add(1)
	s.Insert(0, 2)
	s.Insert(1, 3)
	t.Logf("%v", s)
	assert.Equal(t, "len", len(s), 3)
	assert.StringEqual(t, "s", s, "[2 3 1]")
}

func TestFloatSliceRemove(t *testing.T) {
	defer __(o_(t))

	var s FloatSlice
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

func TestFloatSliceEquals(t *testing.T) {
	s := FloatSlice([]float64{1, 2, 3, 4})

	assert.Equal(t, "s.Equals(nil)", s.Equals(nil, 100), false)
	assert.Equal(t, "s.Equals([1, 2, 3, 4])", s.Equals([]float64{1, 2, 3, 4}, 1e-5), true)
	assert.Equal(t, "s.Equals([1, 2, 5, 4])", s.Equals([]float64{1, 2, 5, 4}, 1e-5), false)
	assert.Equal(t, "s.Equals([1, 2, 5, 4])", s.Equals([]float64{1, 2, 5, 4}, 10), true)
	assert.Equal(t, "s.Equals([1, 2, 3, 4, 5])", s.Equals([]float64{1, 2, 3, 4, 5}, 1e-5), false)

	assert.Equal(t, "nil.Equals([]float64{})", FloatSlice(nil).Equals(s[:0], 100), true)
	assert.Equal(t, "nil.Equals([]float64{1, 2})", FloatSlice(nil).Equals([]float64{1, 2}, 1e-5), false)
}

func BenchmarkFloatSliceInsert(b *testing.B) {
	b.StopTimer()
	s := make(FloatSlice, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s.Insert(1, float64(i))
	} // for i
}

func BenchmarkFloatSliceInsertByAppend(b *testing.B) {
	b.StopTimer()
	s := make([]float64, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s = append(s[:1], append([]float64{float64(i)}, s[1:]...)...)
	} // for i
}

func BenchmarkFloatSliceInsertByCopy(b *testing.B) {
	b.StopTimer()
	s := make([]float64, 100000)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s = append(s, 0)
		copy(s[2:], s[1:])
		s[1] = float64(i)
	} // for i
}

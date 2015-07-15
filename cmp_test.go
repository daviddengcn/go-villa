package villa

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestMerge(t *testing.T) {
	defer __(o_(t))

	var a, b Slice
	for i := 0; i < 100; i++ {
		a.Add(rand.Int())
	} // for i
	for i := 0; i < 200; i++ {
		b.Add(rand.Int())
	} // for i

	var cc Slice
	cc.Add(a...)
	cc.Add(b...)
	intInterfaceCmpFunc.Sort(cc)

	intInterfaceCmpFunc.Sort(a)
	intInterfaceCmpFunc.Sort(b)

	c := intInterfaceCmpFunc.Merge(a, b)
	assert.Equal(t, "len(c)", len(c), len(cc))
	assert.StringEqual(t, "c", c, cc)
}

func TestMergeInt(t *testing.T) {
	defer __(o_(t))

	var a, b IntSlice
	for i := 0; i < 100; i++ {
		a.Add(rand.Int())
	} // for i
	for i := 0; i < 200; i++ {
		b.Add(rand.Int())
	} // for i

	var cc IntSlice
	cc.Add(a...)
	cc.Add(b...)
	IntValueCompare.Sort(cc)

	IntValueCompare.Sort(a)
	IntValueCompare.Sort(b)

	c := IntValueCompare.Merge(a, b)
	assert.Equal(t, "len(c)", len(c), len(cc))
	assert.StringEqual(t, "c", c, cc)
}

func TestMergeString(t *testing.T) {
	defer __(o_(t))

	var a, b StringSlice
	for i := 0; i < 100; i++ {
		a.Add(fmt.Sprint(rand.Int()))
	} // for i
	for i := 0; i < 200; i++ {
		b.Add(fmt.Sprint(rand.Int()))
	} // for i

	var cc StringSlice
	cc.InsertSlice(0, a)
	cc.InsertSlice(0, b)
	StrValueCompare.Sort(cc)

	StrValueCompare.Sort(a)
	StrValueCompare.Sort(b)

	c := StrValueCompare.Merge(a, b)
	assert.Equal(t, "len(c)", len(c), len(cc))
	assert.StringEqual(t, "c", c, cc)
}

func TestMergeFloat(t *testing.T) {
	defer __(o_(t))

	var a, b FloatSlice
	for i := 0; i < 100; i++ {
		a.Add(rand.Float64())
	} // for i
	for i := 0; i < 200; i++ {
		b.Add(rand.Float64())
	} // for i

	var cc FloatSlice
	cc.Add(a...)
	cc.Add(b...)
	FloatValueCompare.Sort(cc)

	FloatValueCompare.Sort(a)
	FloatValueCompare.Sort(b)

	c := FloatValueCompare.Merge(a, b)
	assert.Equal(t, "len(c)", len(c), len(cc))
	assert.StringEqual(t, "c", c, cc)
}

func TestMergeComplex(t *testing.T) {
	defer __(o_(t))

	var a, b ComplexSlice
	for i := 0; i < 100; i++ {
		a.Add(complex(rand.Float64(), rand.Float64()))
	} // for i
	for i := 0; i < 200; i++ {
		b.Add(complex(rand.Float64(), rand.Float64()))
	} // for i

	var cc ComplexSlice
	cc.Add(a...)
	cc.Add(b...)
	cmplexAbsCmpFunc.Sort(cc)

	cmplexAbsCmpFunc.Sort(a)
	cmplexAbsCmpFunc.Sort(b)

	c := cmplexAbsCmpFunc.Merge(a, b)
	assert.Equal(t, "len(c)", len(c), len(cc))
	assert.StringEqual(t, "c", c, cc)
}

func TestSortBinarySearch(t *testing.T) {
	defer __(o_(t))

	s := make(Slice, 0, 100)
	for i := 0; i < 100; i++ {
		s.Add(rand.Int())
	} // for i

	//fmt.Println(s)

	intInterfaceCmpFunc.Sort(s)

	//fmt.Println(s)
	for i := 1; i < len(s); i++ {
		if s[i-1].(int) > s[i].(int) {
			t.Errorf("s[%d](%v) is supposed to be less or equal than s[%d](%v)", i-1, s[i-1], i, s[i])
		} //  if
	} //  if

	for i := range s {
		p, found := intInterfaceCmpFunc.BinarySearch(s, s[i])
		assert.Equal(t, fmt.Sprintf("%d found", i), found, true)
		if found {
			assert.Equal(t, fmt.Sprintf("%d found element", i), s[p], s[i])
		} // if
	} // for i

	for i := range s {
		e := rand.Int()
		p, found := intInterfaceCmpFunc.BinarySearch(s, e)
		if found {
			assert.Equal(t, fmt.Sprintf("%d found element", i), s[p], e)
		} else {
			beforeOk := p == 0 || s[p-1].(int) <= e
			afterOk := p == len(s) || s[p].(int) >= e

			if !beforeOk || !afterOk {
				t.Errorf("Wrong position %d for %v", p, e)
			}
		}
	}
}

func TestStrSortBinarySearch(t *testing.T) {
	defer __(o_(t))

	var s StringSlice
	for i := 0; i < 100; i++ {
		s.Add(fmt.Sprint(rand.Int()))
	} // for i

	//fmt.Println(s)

	StrValueCompare.Sort(s)

	//fmt.Println(s)
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			t.Errorf("s[%d](%v) is supposed to be less or equal than s[%d](%v)", i-1, s[i-1], i, s[i])
		} //  if
	} //  if

	for i := range s {
		p, found := StrValueCompare.BinarySearch(s, s[i])
		assert.Equal(t, fmt.Sprintf("%d found", i), found, true)
		if found {
			assert.Equal(t, fmt.Sprintf("%d found element", i), s[p], s[i])
		}
	}

	for i := range s {
		e := fmt.Sprint(rand.Int())
		p, found := StrValueCompare.BinarySearch(s, e)
		if found {
			assert.Equal(t, fmt.Sprintf("%d found element", i), s[p], e)
		} else {
			beforeOk := p == 0 || s[p-1] <= e
			afterOk := p == len(s) || s[p] >= e

			if !beforeOk || !afterOk {
				t.Errorf("Wrong position %d for %v", p, e)
			}
		}
	}
}

func TestIntSortBinarySearch(t *testing.T) {
	defer __(o_(t))

	var s IntSlice
	for i := 0; i < 100; i++ {
		s.Add(rand.Int())
	} // for i

	//fmt.Println(s)

	IntValueCompare.Sort(s)

	//fmt.Println(s)
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			t.Errorf("s[%d](%v) is supposed to be less or equal than s[%d](%v)", i-1, s[i-1], i, s[i])
		} //  if
	} //  if

	for i := range s {
		p, found := IntValueCompare.BinarySearch(s, s[i])
		assert.Equal(t, fmt.Sprintf("%d found", i), found, true)
		if found {
			assert.Equal(t, fmt.Sprintf("%d found element", i), s[p], s[i])
		}
	}

	for i := range s {
		e := rand.Int()
		p, found := IntValueCompare.BinarySearch(s, e)
		if found {
			assert.Equal(t, fmt.Sprintf("%d found element", i), s[p], e)
		} else {
			beforeOk := p == 0 || s[p-1] <= e
			afterOk := p == len(s) || s[p] >= e

			if !beforeOk || !afterOk {
				t.Errorf("Wrong position %d for %v", p, e)
			}
		}
	}
}

func TestFloatSortBinarySearch(t *testing.T) {
	defer __(o_(t))

	var s FloatSlice
	for i := 0; i < 100; i++ {
		s.Add(rand.Float64())
	} // for i

	//fmt.Println(s)

	FloatValueCompare.Sort(s)

	//fmt.Println(s)
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			t.Errorf("s[%d](%v) is supposed to be less or equal than s[%d](%v)", i-1, s[i-1], i, s[i])
		} //  if
	} //  if

	for i := range s {
		p, found := FloatValueCompare.BinarySearch(s, s[i])
		assert.Equal(t, fmt.Sprintf("%d found", i), found, true)
		if found {
			assert.Equal(t, fmt.Sprintf("%d found element", i), s[p], s[i])
		}
	}

	for i := range s {
		e := rand.Float64()
		p, found := FloatValueCompare.BinarySearch(s, e)
		if found {
			assert.Equal(t, fmt.Sprintf("%d found element", i), s[p], e)
		} else {
			beforeOk := p == 0 || s[p-1] <= e
			afterOk := p == len(s) || s[p] >= e

			if !beforeOk || !afterOk {
				t.Errorf("Wrong position %d for %v", p, e)
			}
		}
	}
}

func TestComplexSortBinarySearch(t *testing.T) {
	defer __(o_(t))

	var s ComplexSlice
	for i := 0; i < 100; i++ {
		s.Add(complex(rand.Float64(), rand.Float64()))
	} // for i

	//fmt.Println(s)

	cmplexAbsCmpFunc.Sort(s)

	//fmt.Println(s)
	for i := 1; i < len(s); i++ {
		if cmplexAbsCmpFunc(s[i-1], s[i]) > 0 {
			t.Errorf("s[%d](%v) is supposed to be less or equal than s[%d](%v)", i-1, s[i-1], i, s[i])
		} //  if
	} //  if

	for i := range s {
		p, found := cmplexAbsCmpFunc.BinarySearch(s, s[i])
		assert.Equal(t, fmt.Sprintf("%d found", i), found, true)
		if found {
			assert.Equal(t, fmt.Sprintf("%d found element", i), s[p], s[i])
		}
	}

	for i := range s {
		e := complex(rand.Float64(), rand.Float64())
		p, found := cmplexAbsCmpFunc.BinarySearch(s, e)
		if found {
			assert.Equal(t, fmt.Sprintf("%d found element", i), s[p], e)
		} else {
			beforeOk := p == 0 || cmplexAbsCmpFunc(s[p-1], e) <= 0
			afterOk := p == len(s) || cmplexAbsCmpFunc(s[p], e) >= 0

			if !beforeOk || !afterOk {
				t.Errorf("Wrong position %d for %v", p, e)
			}
		}
	}
}

func TestStrDiff(t *testing.T) {
	defer __(o_(t))

	s1 := StringSlice{"a", "b", "d", "f"}
	s2 := StringSlice{"b", "c", "d", "g"}

	d1, d2 := StrValueCompare.DiffSlicePair(s1, s2)
	assert.StringEqual(t, "d1", d1, "[a f]")
	assert.StringEqual(t, "d2", d2, "[c g]")
}

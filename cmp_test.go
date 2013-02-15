package villa

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMerge(t *testing.T) {
	defer __(o_())

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
	AssertEquals(t, "len(c)", len(c), len(cc))
	AssertStringEquals(t, "c", c, cc)
}

func TestMergeInt(t *testing.T) {
	defer __(o_())

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
	AssertEquals(t, "len(c)", len(c), len(cc))
	AssertStringEquals(t, "c", c, cc)
}

func TestMergeFloat(t *testing.T) {
	defer __(o_())

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
	AssertEquals(t, "len(c)", len(c), len(cc))
	AssertStringEquals(t, "c", c, cc)
}

func TestMergeComplex(t *testing.T) {
	defer __(o_())

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
	AssertEquals(t, "len(c)", len(c), len(cc))
	AssertStringEquals(t, "c", c, cc)
}

func TestSortBinarySearch(t *testing.T) {
	defer __(o_())

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
		AssertEquals(t, fmt.Sprintf("%d found", i), found, true)
		if found {
			AssertEquals(t, fmt.Sprintf("%d found element", i), s[p], s[i])
		} // if
	} // for i

	for i := range s {
		e := rand.Int()
		p, found := intInterfaceCmpFunc.BinarySearch(s, e)
		if found {
			AssertEquals(t, fmt.Sprintf("found element", i), s[p], e)
		} else {
			beforeOk := p == 0 || s[p-1].(int) <= e
			afterOk := p == len(s) || s[p].(int) >= e

			if !beforeOk || !afterOk {
				t.Errorf("Wrong position %d for %v", p, e)
			} // if
		} // else
	} // for i
}

func TestIntSortBinarySearch(t *testing.T) {
	defer __(o_())

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
		AssertEquals(t, fmt.Sprintf("%d found", i), found, true)
		if found {
			AssertEquals(t, fmt.Sprintf("%d found element", i), s[p], s[i])
		} // if
	} // for i

	for i := range s {
		e := rand.Int()
		p, found := IntValueCompare.BinarySearch(s, e)
		if found {
			AssertEquals(t, fmt.Sprintf("found element", i), s[p], e)
		} else {
			beforeOk := p == 0 || s[p-1] <= e
			afterOk := p == len(s) || s[p] >= e

			if !beforeOk || !afterOk {
				t.Errorf("Wrong position %d for %v", p, e)
			} // if
		} // else
	} // for i
}

func TestFloatSortBinarySearch(t *testing.T) {
	defer __(o_())

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
		AssertEquals(t, fmt.Sprintf("%d found", i), found, true)
		if found {
			AssertEquals(t, fmt.Sprintf("%d found element", i), s[p], s[i])
		} // if
	} // for i

	for i := range s {
		e := rand.Float64()
		p, found := FloatValueCompare.BinarySearch(s, e)
		if found {
			AssertEquals(t, fmt.Sprintf("found element", i), s[p], e)
		} else {
			beforeOk := p == 0 || s[p-1] <= e
			afterOk := p == len(s) || s[p] >= e

			if !beforeOk || !afterOk {
				t.Errorf("Wrong position %d for %v", p, e)
			} // if
		} // else
	} // for i
}

func TestComplexSortBinarySearch(t *testing.T) {
	defer __(o_())

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
		AssertEquals(t, fmt.Sprintf("%d found", i), found, true)
		if found {
			AssertEquals(t, fmt.Sprintf("%d found element", i), s[p], s[i])
		} // if
	} // for i

	for i := range s {
		e := complex(rand.Float64(), rand.Float64())
		p, found := cmplexAbsCmpFunc.BinarySearch(s, e)
		if found {
			AssertEquals(t, fmt.Sprintf("found element", i), s[p], e)
		} else {
			beforeOk := p == 0 || cmplexAbsCmpFunc(s[p-1], e) <= 0
			afterOk := p == len(s) || cmplexAbsCmpFunc(s[p], e) >= 0

			if !beforeOk || !afterOk {
				t.Errorf("Wrong position %d for %v", p, e)
			} // if
		} // else
	} // for i
}

package villa

import(
    "testing"
    "sort"
    "math/rand"
    "math/cmplx"
    "fmt"
)

func cmplexAbsCmpFunc(a, b complex128) int {
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
}

func TestComplexSlice(t *testing.T) {
    fmt.Println("== Begin TestComplexSlice...");
    defer fmt.Println("== End TestComplexSlice.");
    
    var s ComplexSlice
    for i := 0; i < 1000; i ++ {
        s.Add(complex(float64(i), float64(-i)))
    } // for i
    
    AssertEquals(t, "len", len(s), 1000)
    //fmt.Println(s)
    s.Clear()
    AssertEquals(t, "len", len(s), 0)
    
    s = ComplexSlice{}
    s.Add(-4, -1)
    s.Insert(1, -2, -3)
    fmt.Println(s)
    AssertEquals(t, "len", len(s), 4)
    AssertStringEquals(t, "s", s, "[(-4+0i) (-2+0i) (-3+0i) (-1+0i)]")
    
    sort.Sort(s.NewSortList(cmplexAbsCmpFunc))
    AssertStringEquals(t, "s", s, "[(-1+0i) (-2+0i) (-3+0i) (-4+0i)]")
}

func TestComplexSliceRemove(t *testing.T) {
    fmt.Println("== Begin TestComplexSliceRemove...");
    defer fmt.Println("== End TestComplexSliceRemove.");
    var s ComplexSlice
    s.Add(-1, -2, -3, -4, -5, -6, -7)
    AssertEquals(t, "len", len(s), 7)
    AssertStringEquals(t, "s", s, "[(-1+0i) (-2+0i) (-3+0i) (-4+0i) (-5+0i) (-6+0i) (-7+0i)]")
    
    s.RemoveRange(2, 5)
    AssertEquals(t, "len", len(s), 4)
    AssertStringEquals(t, "s", s, "[(-1+0i) (-2+0i) (-6+0i) (-7+0i)]")
    
    s.Remove(2)
    AssertEquals(t, "len", len(s), 3)
    AssertStringEquals(t, "s", s, "[(-1+0i) (-2+0i) (-7+0i)]")
}

func TestComplexSliceSort(t *testing.T) {
    var s ComplexSlice
    for i := 0; i < 100; i ++ {
        s.Add(complex(rand.Float64(), rand.Float64()))
    } // for i
    
    //fmt.Println(s)
    
    adp := s.NewSortList(cmplexAbsCmpFunc)
    sort.Sort(adp)
    
    //fmt.Println(s)
    for i := 1; i < len(s); i ++ {
        if cmplexAbsCmpFunc(s[i - 1], s[i]) > 0 {
            t.Errorf("s[%d](%v) is supposed to be less or equal than s[%d](%v)", i - 1, s[i - 1], i, s[i])
        } //  if
    } //  if
    
    for i := range(s) {
        p, found := adp.BinarySearch(s[i])
        AssertEquals(t, fmt.Sprintf("%d found", i), found, true)
        if found {
            AssertEquals(t, fmt.Sprintf("%d found element", i), s[p], s[i])
        } // if
    } // for i
    
    for i := range(s) {
        e := complex(rand.Float64(), rand.Float64())
        p, found := adp.BinarySearch(e)
        if found {
            AssertEquals(t, fmt.Sprintf("found element", i), s[p], e)
        } else {
            beforeOk := p == 0 || cmplexAbsCmpFunc(s[p - 1], e) <= 0;
            afterOk := p == len(s) || cmplexAbsCmpFunc(s[p], e) >= 0;
            
            if !beforeOk || !afterOk {
                t.Errorf("Wrong position %d for %v", p, e)
            } // if
        } // else
    } // for i
}

func BenchmarkComplexSliceInsert(b *testing.B) {
    b.StopTimer()
    s := make(ComplexSlice, 100000, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s.Insert(1, complex(float64(i), float64(-i)))
    } // for i
}

func BenchmarkComplexSliceInsertByAppend(b *testing.B) {
    b.StopTimer()
    s := make([]complex128, 100000, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s = append(s[:1], append([]complex128{complex(float64(i), float64(-i))}, s[1:]...)...)
    } // for i
}

func BenchmarkComplexSliceInsertByCopy(b *testing.B) {
    b.StopTimer()
    s := make([]complex128, 100000, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s = append(s, 0)
        copy(s[2:], s[1:])
        s[1] = complex(float64(i), float64(-i))
    } // for i
}
  
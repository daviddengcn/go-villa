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

func TestComplexArrayList(t *testing.T) {
    fmt.Println("== Begin TestComplexArrayList...");
    defer fmt.Println("== End TestComplexArrayList.");
    
    lst := NewComplexArrayList()
    for i := 0; i < 1000; i ++ {
        lst.Add(complex(float64(i), float64(-i)))
    } // for i
    
    AssertEquals(t, "lst.Len()", lst.Len(), 1000)
    //fmt.Println(lst)
    lst.Clear()
    AssertEquals(t, "lst.Len()", lst.Len(), 0)
    
    lst = NewComplexArrayList()
    lst.Add(complex(1, -1))
    lst.Insert(0, complex(2, -2))
    lst.Insert(1, complex(3, -3))
    fmt.Println(lst)
    AssertEquals(t, "lst.Len()", lst.Len(), 3)
    // FIXME: when go compiler fixes its bug of []complex64/128.String(), uncommet the following line
    // AssertStringEquals(t, "lst", "[(2-2i) (3-3i) (1-1i)]", lst)
    
    sort.Sort(lst.NewCmpAdapter(cmplexAbsCmpFunc))
    // FIXME: when go compiler fixes its bug of []complex64/128.String(), uncommet the following line
    // AssertStringEquals(t, "lst", lst, "[(1-i) (2-2i) (3-3i)]")
}

func TestComplexArrayListRemove(t *testing.T) {
    fmt.Println("== Begin TestComplexArrayListRemove...");
    defer fmt.Println("== End TestComplexArrayListRemove.");
    
    lst := NewComplexArrayList()
    lst.AddSlice([]complex128{1, 2, 3, 4, 5, 6, 7})
    AssertEquals(t, "lst.Len()", lst.Len(), 7)
    // FIXME: when go compiler fixes its bug of []complex64/128.String(), uncommet the following line
    // AssertStringEquals(t, "lst", lst, "[1 2 3 4 5 6 7]")
    
    lst.RemoveRange(2, 5)
    AssertEquals(t, "lst.Len()", lst.Len(), 4)
    // FIXME: when go compiler fixes its bug of []complex64/128.String(), uncommet the following line
    // AssertStringEquals(t, "lst", lst, "[1 2 6 7]")
    
    lst.Remove(2)
    AssertEquals(t, "lst.Len()", lst.Len(), 3)
    // FIXME: when go compiler fixes its bug of []complex64/128.String(), uncommet the following line
    // AssertStringEquals(t, "lst", lst, "[1 2 7]")
}

func TestComplexArrayListSort(t *testing.T) {
    lst := NewComplexArrayList()
    for i := 0; i < 100; i ++ {
        lst.Add(complex(rand.Float64(), rand.Float64()))
    } // for i
    
    //fmt.Println(lst)
    
    sort.Sort(lst.NewCmpAdapter(cmplexAbsCmpFunc))
    
    //fmt.Println(lst)
    for i := 1; i < lst.Len(); i ++ {
        if cmplx.Abs(lst.Get(i - 1)) > cmplx.Abs(lst.Get(i)) {
            t.Errorf("lst[%d](%v) is supposed to be less or equal than lst[%d](%v)", i - 1, lst.Get(i - 1), i, lst.Get(i))
        } //  if
    } //  if
}

func BenchmarkComplexArrayListInsert(b *testing.B) {
    b.StopTimer()
    lst := NewComplexArrayList()
    for i := 0; i < 100000; i ++ {
        lst.Push(complex(float64(i), float64(-i)))
    } // for i
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst.Insert(1, complex(float64(i), float64(-i)))
    } // for i
}

func BenchmarkComplexSliceInsertByAppend(b *testing.B) {
    b.StopTimer()
    lst := []complex128{}
    for i := 0; i < 100000; i ++ {
        lst = append(lst, complex(float64(i), float64(-i)))
    } // for i
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst = append(lst[:1], append([]complex128{complex(float64(i), float64(-i))}, lst[1:]...)...)
    } // for i
}

func BenchmarkComplexSliceInsertByCopy(b *testing.B) {
    b.StopTimer()
    lst := []complex128{}
    for i := 0; i < 100000; i ++ {
        lst = append(lst, complex(float64(i), float64(-i)))
    } // for i
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst = append(lst, 0)
        copy(lst[2:], lst[1:])
        lst[1] = complex(float64(i), float64(-i))
    } // for i
}

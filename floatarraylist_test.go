package villa

import(
    "testing"
    "sort"
    "math/rand"
    "fmt"
)

func TestFloatArrayList(t *testing.T) {
    fmt.Println("== Begin TestFloatArrayList...");
    defer fmt.Println("== End TestFloatArrayList.");
    
    lst := NewFloatArrayList()
    for i := 0; i < 1000; i ++ {
        lst.Add(0.5*float64(i))
    } // for i
    
    AssertEquals(t, "lst.Len()", lst.Len(), 1000)
    //fmt.Println(lst)
    lst.Clear()
    AssertEquals(t, "lst.Len()", lst.Len(), 0)
    
    lst = NewFloatArrayList()
    lst.Add(1)
    lst.Insert(0, 2)
    lst.Insert(1, 3)
    fmt.Println(lst)
    AssertEquals(t, "lst.Len()", lst.Len(), 3)
    AssertStringEquals(t, "lst", lst, "[2 3 1]")
    
    sort.Sort(lst.NewCmpAdapter(FloatValueCompare))
    AssertStringEquals(t, "lst", lst, "[1 2 3]")
}

func TestFloatArrayListCap(t *testing.T) {
    lst := NewFloatArrayListCap(10, 20)
    AssertEquals(t, "lst.Len()", lst.Len(), 10)
    AssertEquals(t, "cap(lst.Data())", cap(lst.Data()), 20)
}

func TestFloatArrayListRemove(t *testing.T) {
    fmt.Println("== Begin TestFloatArrayListRemove...");
    defer fmt.Println("== End TestFloatArrayListRemove.");
    
    lst := NewFloatArrayList()
    lst.AddSlice([]float64{1, 2, 3, 4, 5, 6, 7})
    AssertEquals(t, "lst.Len()", lst.Len(), 7)
    AssertStringEquals(t, "lst", lst, "[1 2 3 4 5 6 7]")
    
    lst.RemoveRange(2, 5)
    AssertEquals(t, "lst.Len()", lst.Len(), 4)
    AssertStringEquals(t, "lst", lst, "[1 2 6 7]")
    
    lst.Remove(2)
    AssertEquals(t, "lst.Len()", lst.Len(), 3)
    AssertStringEquals(t, "lst", lst, "[1 2 7]")
}

func TestFloatArrayListSort(t *testing.T) {
    lst := NewFloatArrayList()
    for i := 0; i < 100; i ++ {
        lst.Add(rand.Float64())
    } // for i
    
    //fmt.Println(lst)
    
    sort.Sort(lst.NewCmpAdapter(FloatValueCompare))
    
    //fmt.Println(lst)
    for i := 1; i < lst.Len(); i ++ {
        if lst.Get(i - 1) > lst.Get(i) {
            t.Errorf("lst[%d](%v) is supposed to be less or equal than lst[%d](%v)", i - 1, lst.Get(i - 1), i, lst.Get(i))
        } //  if
    } //  if
}

func BenchmarkFloatArrayListInsert(b *testing.B) {
    b.StopTimer()
    lst := NewFloatArrayList()
    for i := 0; i < 100000; i ++ {
        lst.Push(float64(i))
    } // for i
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst.Insert(1, float64(i))
    } // for i
}

func BenchmarkFloatSliceInsertByAppend(b *testing.B) {
    b.StopTimer()
    lst := []float64{}
    for i := 0; i < 100000; i ++ {
        lst = append(lst, float64(i))
    } // for i
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst = append(lst[:1], append([]float64{float64(i)}, lst[1:]...)...)
    } // for i
}

func BenchmarkFloatSliceInsertByCopy(b *testing.B) {
    b.StopTimer()
    lst := []float64{}
    for i := 0; i < 100000; i ++ {
        lst = append(lst, float64(i))
    } // for i
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst = append(lst, 0)
        copy(lst[2:], lst[1:])
        lst[1] = float64(i)
    } // for i
}
 
package villa

import(
    "testing"
    "sort"
    "math/rand"
    "fmt"
)

func TestIntArrayList(t *testing.T) {
    fmt.Println("== Begin TestIntArrayList...");
    defer fmt.Println("== End TestIntArrayList.");
    
    lst := NewIntArrayList()
    for i := 0; i < 1000; i ++ {
        lst.Add(i)
    } // for i
    
    AssertEquals(t, "lst.Len()", lst.Len(), 1000)
    //fmt.Println(lst)
    lst.Clear()
    AssertEquals(t, "lst.Len()", lst.Len(), 0)
    
    lst = NewIntArrayList()
    lst.Add(1)
    lst.Insert(0, 2)
    lst.Insert(1, 3)
    fmt.Println(lst)
    AssertEquals(t, "lst.Len()", lst.Len(), 3)
    AssertStringEquals(t, "lst", lst, "[2 3 1]")
    
    sort.Sort(lst.NewCmpAdapter(IntValueCompare))
    AssertStringEquals(t, "lst", lst, "[1 2 3]")
}

func TestIntArrayListCap(t *testing.T) {
    lst := NewIntArrayListCap(10, 20)
    AssertEquals(t, "lst.Len()", lst.Len(), 10)
    AssertEquals(t, "cap(lst.Data())", cap(lst.Data()), 20)
}

func TestIntArrayListRemove(t *testing.T) {
    fmt.Println("== Begin TestIntArrayListRemove...");
    defer fmt.Println("== End TestIntArrayListRemove.");
    lst := NewIntArrayList()
    lst.AddSlice([]int{1, 2, 3, 4, 5, 6, 7})
    AssertEquals(t, "lst.Len()", lst.Len(), 7)
    AssertStringEquals(t, "lst", lst, "[1 2 3 4 5 6 7]")
    
    lst.RemoveRange(2, 5)
    AssertEquals(t, "lst.Len()", lst.Len(), 4)
    AssertStringEquals(t, "lst", lst, "[1 2 6 7]")
    
    lst.Remove(2)
    AssertEquals(t, "lst.Len()", lst.Len(), 3)
    AssertStringEquals(t, "lst", lst, "[1 2 7]")
}

func TestIntArrayListSort(t *testing.T) {
    lst := NewIntArrayList()
    for i := 0; i < 100; i ++ {
        lst.Add(rand.Int())
    } // for i
    
    //fmt.Println(lst)
    
    adp := lst.NewCmpAdapter(IntValueCompare)
    sort.Sort(adp)
    
    //fmt.Println(lst)
    for i := 1; i < lst.Len(); i ++ {
        if lst.Get(i - 1) > lst.Get(i) {
            t.Errorf("lst[%d](%v) is supposed to be less or equal than lst[%d](%v)", i - 1, lst.Get(i - 1), i, lst.Get(i))
        } //  if
    } //  if
    
    for i := 0; i < lst.Len(); i ++ {
        p, found := adp.BinarySearch(lst.Get(i))
        AssertEquals(t, fmt.Sprintf("%d found", i), found, true)
        if found {
            AssertEquals(t, fmt.Sprintf("%d found element", i), lst.Get(p), lst.Get(i))
        } // if
    } // for i
    
    for i := 0; i < lst.Len(); i ++ {
        e := rand.Int()
        p, found := adp.BinarySearch(e)
        if found {
            AssertEquals(t, fmt.Sprintf("found element", i), lst.Get(p), e)
        } else {
            beforeOk := p == 0 || lst.Get(p - 1) <= e;
            afterOk := p == lst.Len() || lst.Get(p) >= e;
            
            if !beforeOk || !afterOk {
                t.Errorf("Wrong position %d for %v", p, e)
            } // if
        } // else
    } // for i
}

func BenchmarkIntArrayListInsert(b *testing.B) {
    b.StopTimer()
    lst := NewIntArrayList()
    for i := 0; i < 100000; i ++ {
        lst.Push(i)
    } // for i
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst.Insert(1, i)
    } // for i
}

func BenchmarkIntSliceInsertByAppend(b *testing.B) {
    b.StopTimer()
    lst := []int{}
    for i := 0; i < 100000; i ++ {
        lst = append(lst, i)
    } // for i
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst = append(lst[:1], append([]int{i}, lst[1:]...)...)
    } // for i
}

func BenchmarkIntSliceInsertByCopy(b *testing.B) {
    b.StopTimer()
    lst := []int{}
    for i := 0; i < 100000; i ++ {
        lst = append(lst, i)
    } // for i
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst = append(lst, 0)
        copy(lst[2:], lst[1:])
        lst[1] = i
    } // for i
}
 
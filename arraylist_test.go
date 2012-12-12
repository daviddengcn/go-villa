package villa

import(
    "testing"
    "sort"
    "math/rand"
    "fmt"
)

func AssertEquals(t *testing.T, name string, a, b interface{}) {
    if a != b {
        t.Errorf("%s is expected to be %v, but got %v", name, a, b)
    } // if
}

func AssertStringEquals(t *testing.T, name string, a, b interface{}) {
    if fmt.Sprintf("%v", a) != fmt.Sprintf("%v", b) {
        t.Errorf("%s is expected to be %v, but got %v", name, a, b)
    } // if
}

func intInterfaceCmpFunc(e1, e2 interface{}) int {
    if e1.(int) < e2.(int) {
        return -1
    } else if e1.(int) > e2.(int) {
        return 1
    } // else if
    return 0
}

func TestArrayList(t *testing.T) {
    fmt.Println("== Begin TestArrayList...");
    defer fmt.Println("== End TestArrayList.");
    
    lst := NewArrayList()
    for i := 0; i < 1000; i ++ {
        lst.Add(i)
    } // for i
    
    AssertEquals(t, "lst.Len()", lst.Len(), 1000)
    //fmt.Println(lst)
    lst.Clear()
    AssertEquals(t, "lst.Len()", lst.Len(), 0)
    
    lst = NewArrayList()
    lst.Add(1)
    lst.Insert(0, 2)
    lst.Insert(1, 3)
    AssertEquals(t, "lst.Len()", lst.Len(), 3)
    AssertStringEquals(t, "lst", lst, "[2 3 1]")
    
    sort.Sort(lst.NewCmpAdapter(intInterfaceCmpFunc))
    AssertStringEquals(t, "lst", lst, "[1 2 3]")
}

func TestArrayListRemove(t *testing.T) {
    lst := NewArrayList()
    lst.AddSlice([]interface{}{1, 2, 3, 4, 5, 6, 7})
    AssertEquals(t, "lst.Len()", lst.Len(), 7)
    AssertStringEquals(t, "lst", lst, "[1 2 3 4 5 6 7]")
    
    lst.RemoveRange(2, 5)
    AssertEquals(t, "lst.Len()", lst.Len(), 4)
    AssertStringEquals(t, "lst", lst, "[1 2 6 7]")
    
    lst.Remove(2)
    AssertEquals(t, "lst.Len()", lst.Len(), 3)
    AssertStringEquals(t, "lst", lst, "[1 2 7]")
}

func TestArrayListSort(t *testing.T) {
    lst := NewArrayList()
    for i := 0; i < 100; i ++ {
        lst.Add(int(rand.Int31()))
    } // for i
    
    //fmt.Println(lst)
    
    sort.Sort(lst.NewCmpAdapter(intInterfaceCmpFunc))
    
    //fmt.Println(lst)
    for i := 1; i < lst.Len(); i ++ {
        if lst.Get(i - 1).(int) > lst.Get(i).(int) {
            t.Errorf("lst[%d](%v) is supposed to be less or equal than lst[%d](%v)", i - 1, lst.Get(i - 1), i, lst.Get(i))
        } //  if
    } //  if
}

func BenchmarkArrayListInsert(b *testing.B) {
    b.StopTimer()
    lst := NewArrayList()
    for i := 0; i < 100000; i ++ {
        lst.Push(i)
    } // for i
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst.Insert(1, i)
    } // for i
}

func BenchmarkSliceInsertByAppend(b *testing.B) {
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

func BenchmarkSliceInsertByCopy(b *testing.B) {
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

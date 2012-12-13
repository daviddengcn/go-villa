package villa

import(
    "testing"
    "sort"
    "math/rand"
    "fmt"
)

func AssertEquals(t *testing.T, name string, act, exp interface{}) {
    if act != exp {
        t.Errorf("%s is expected to be %v, but got %v", name, exp, act)
    } // if
}

func AssertStringEquals(t *testing.T, name string, act, exp interface{}) {
    if fmt.Sprintf("%v", act) != fmt.Sprintf("%v", exp) {
        t.Errorf("%s is expected to be %v, but got %v", name, exp, act)
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
    
    var lst ArrayList
    for i := 0; i < 1000; i ++ {
        lst.Add(i)
    } // for i
    
    AssertEquals(t, "len(lst)", len(lst), 1000)
    //fmt.Println(lst)
    lst.Clear()
    AssertEquals(t, "len(lst)", len(lst), 0)
    
    lst = ArrayList{}
    lst.Add(4, 1)
    lst.Insert(1, 2, 3)
    AssertEquals(t, "len(lst)", len(lst), 4)
    AssertStringEquals(t, "lst", lst, "[4 2 3 1]")
    
    sort.Sort(lst.NewSortAdapter(intInterfaceCmpFunc))
    AssertStringEquals(t, "lst", lst, "[1 2 3 4]")
}

func TestArrayListRemove(t *testing.T) {
    var lst ArrayList
    lst.Add(1, 2, 3, 4, 5, 6, 7)
    AssertEquals(t, "len(lst)", len(lst), 7)
    AssertStringEquals(t, "lst", lst, "[1 2 3 4 5 6 7]")
    
    lst.RemoveRange(2, 5)
    AssertEquals(t, "len(lst)", len(lst), 4)
    AssertStringEquals(t, "lst", lst, "[1 2 6 7]")
    
    lst.Remove(2)
    AssertEquals(t, "len(lst)", len(lst), 3)
    AssertStringEquals(t, "lst", lst, "[1 2 7]")
}

func TestArrayListSort(t *testing.T) {
    lst := make(ArrayList, 0, 100)
    for i := 0; i < 100; i ++ {
        lst.Add(rand.Int())
    } // for i
    
    //fmt.Println(lst)
    
    adp := lst.NewSortAdapter(intInterfaceCmpFunc)
    sort.Sort(adp)
    
    //fmt.Println(lst)
    for i := 1; i < len(lst); i ++ {
        if lst[i - 1].(int) > lst[i].(int) {
            t.Errorf("lst[%d](%v) is supposed to be less or equal than lst[%d](%v)", i - 1, lst[i - 1], i, lst[i])
        } //  if
    } //  if
    
    for i := range(lst) {
        p, found := adp.BinarySearch(lst[i])
        AssertEquals(t, fmt.Sprintf("%d found", i), found, true)
        if found {
            AssertEquals(t, fmt.Sprintf("%d found element", i), lst[p], lst[i])
        } // if
    } // for i
    
    for i := range(lst) {
        e := rand.Int()
        p, found := adp.BinarySearch(e)
        if found {
            AssertEquals(t, fmt.Sprintf("found element", i), lst[p], e)
        } else {
            beforeOk := p == 0 || lst[p - 1].(int) <= e;
            afterOk := p == len(lst) || lst[p].(int) >= e;
            
            if !beforeOk || !afterOk {
                t.Errorf("Wrong position %d for %v", p, e)
            } // if
        } // else
    } // for i
}

func BenchmarkArrayListInsert(b *testing.B) {
    b.StopTimer()
    lst := make(ArrayList, 100000, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst.Insert(1, i)
    } // for i
}

func BenchmarkSliceInsertByAppend(b *testing.B) {
    b.StopTimer()
    lst := make([]int, 100000, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst = append(lst[:1], append([]int{i}, lst[1:]...)...)
    } // for i
}

func BenchmarkSliceInsertByCopy(b *testing.B) {
    b.StopTimer()
    lst := make([]int, 100000, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        lst = append(lst, 0)
        copy(lst[2:], lst[1:])
        lst[1] = i
    } // for i
}

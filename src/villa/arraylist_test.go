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


func TestArrayList(t *testing.T) {
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
    
    sort.Sort(lst.NewLessAdapter(func(e1, e2 interface{}) bool {
        return e1.(int) < e2.(int)
    }))
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
        lst.Add(rand.Int31())
    } // for i
    
    //fmt.Println(lst)
    
    sort.Sort(lst.NewLessAdapter(
        func (a, b interface{}) bool {
            return a.(int32) < b.(int32)
        }))
    
    //fmt.Println(lst)
    for i := 1; i < lst.Len(); i ++ {
        if lst.Get(i - 1).(int32) > lst.Get(i).(int32) {
            t.Errorf("lst[%d](%v) is supposed to be less or equal than lst[%d](%v)", i - 1, lst.Get(i - 1), i, lst.Get(i))
        } //  if
    } //  if
}

func BenchmarkArrayList(b *testing.B) {
    lst := NewArrayList()
    for i := 0; i < b.N; i ++ {
        for i := 0; i < 10000; i ++ {
            lst.Insert(0, i)
        } // for i
        
        for lst.Len() > 0 {
            lst.Remove(0)
        } // for
    } // for i
}

package villa

import(
    "testing"
    "fmt"
    "runtime"
    "strings"
)

func o_() string {
    pc, _, _, _ := runtime.Caller(1)
    name := runtime.FuncForPC(pc).Name()
    if p := strings.LastIndexAny(name, `./\`); p >= 0 {
        name = name[p+1:]
    } // if
    fmt.Println("== BEGIN", name, "===")
    return name
}

func __(name string) {
    fmt.Println("== END", name, "===")
}

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

var intInterfaceCmpFunc = CmpFunc(func(e1, e2 interface{}) int {
    if e1.(int) < e2.(int) {
        return -1
    } else if e1.(int) > e2.(int) {
        return 1
    } // else if
    return 0
})

func TestSlice(t *testing.T) {
    defer __(o_())
    
    var s Slice
    for i := 0; i < 1000; i ++ {
        s.Add(i)
    } // for i
    
    AssertEquals(t, "len(s)", len(s), 1000)
    //fmt.Println(s)
    s.Clear()
    AssertEquals(t, "len(s)", len(s), 0)
    
    s = Slice{}
    s.Add(4, 1)
    s.Insert(1, 2, 3)
    AssertEquals(t, "len(s)", len(s), 4)
    AssertStringEquals(t, "s", s, "[4 2 3 1]")
}

func ExampleSlice() {
    var s Slice
    s.Add(1, 2, 3, 4, 5)
    fmt.Println(len(s), s)
    
    s.Fill(1, 4, 10)
    fmt.Println(s)
    
    s.AddSlice([]int{20, 21, 22})
    fmt.Println(s)
    s.AddSlice([]string{"23", "24"})
    fmt.Println(s)

    type A struct {
        X int
        Y int
    }
    
    type B []A
    b := B{A{10, 20}, A{30, 40}}
    s.AddSlice(b)
    fmt.Println(s)
    
    type I interface {
        Func()
    }
    type SI []I
    si := make(SI, 2)
    s.AddSlice(si)
    fmt.Println(s)
    // Output: 
    // 5 [1 2 3 4 5]
    // [1 10 10 10 5]
    // [1 10 10 10 5 20 21 22]
    // [1 10 10 10 5 20 21 22 23 24]
    // [1 10 10 10 5 20 21 22 23 24 {10 20} {30 40}]
    // [1 10 10 10 5 20 21 22 23 24 {10 20} {30 40} <nil> <nil>]
}

func TestSliceRemove(t *testing.T) {
    defer __(o_())
    
    var s Slice
    s.Add(1, 2, 3, 4, 5, 6, 7)
    AssertEquals(t, "len(s)", len(s), 7)
    AssertStringEquals(t, "s", s, "[1 2 3 4 5 6 7]")
    
    s.RemoveRange(2, 5)
    AssertEquals(t, "len(s)", len(s), 4)
    AssertStringEquals(t, "s", s, "[1 2 6 7]")
    
    s.Remove(2)
    AssertEquals(t, "len(s)", len(s), 3)
    AssertStringEquals(t, "s", s, "[1 2 7]")
}

func BenchmarkSliceInsert(b *testing.B) {
    b.StopTimer()
    s := make(Slice, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s.Insert(1, i)
    } // for i
}

func BenchmarkSliceInsertByAppend(b *testing.B) {
    b.StopTimer()
    s := make([]int, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s = append(s[:1], append([]int{i}, s[1:]...)...)
    } // for i
}

func BenchmarkSliceInsertByCopy(b *testing.B) {
    b.StopTimer()
    s := make([]int, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s = append(s, 0)
        copy(s[2:], s[1:])
        s[1] = i
    } // for i
}

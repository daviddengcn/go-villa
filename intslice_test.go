package villa

import(
    "testing"
    "fmt"
)

func TestIntSlice(t *testing.T) {
    defer __(o_())
    
    var s IntSlice
    for i := 0; i < 1000; i ++ {
        s.Add(i)
    } // for i
    
    AssertEquals(t, "len", len(s), 1000)
    //fmt.Println(s)
    s.Clear()
    AssertEquals(t, "len", len(s), 0)
    
    s = IntSlice{}
    s.Add(1)
    s.Insert(0, 2)
    s.Insert(1, 3)
    fmt.Println(s)
    AssertEquals(t, "len", len(s), 3)
    AssertStringEquals(t, "s", s, "[2 3 1]")
}

func TestIntSliceRemove(t *testing.T) {
    defer __(o_())

    var s IntSlice
    s.Add(1, 2, 3, 4, 5, 6, 7)
    AssertEquals(t, "len", len(s), 7)
    AssertStringEquals(t, "s", s, "[1 2 3 4 5 6 7]")
    
    s.Fill(2, 5, 9)
    AssertStringEquals(t, "s", s, "[1 2 9 9 9 6 7]")
    
    s.RemoveRange(2, 5)
    AssertEquals(t, "len", len(s), 4)
    AssertStringEquals(t, "s", s, "[1 2 6 7]")
    
    s.Remove(2)
    AssertEquals(t, "len", len(s), 3)
    AssertStringEquals(t, "s", s, "[1 2 7]")
}

func BenchmarkIntSliceInsert(b *testing.B) {
    b.StopTimer()
    s := make(IntSlice, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s.Insert(1, i)
    } // for i
}

func BenchmarkIntSliceInsertByAppend(b *testing.B) {
    b.StopTimer()
    s := make([]int, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s = append(s[:1], append([]int{i}, s[1:]...)...)
    } // for i
}

func BenchmarkIntSliceInsertByCopy(b *testing.B) {
    b.StopTimer()
    s := make([]int, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s = append(s, 0)
        copy(s[2:], s[1:])
        s[1] = i
    } // for i
}
 
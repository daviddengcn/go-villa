package villa

import(
    "testing"
    "fmt"
)

func TestFloatSlice(t *testing.T) {
    fmt.Println("== Begin TestFloatSlice...");
    defer fmt.Println("== End TestFloatSlice.");
    
    var s FloatSlice
    for i := 0; i < 1000; i ++ {
        s.Add(float64(i))
    } // for i
    
    AssertEquals(t, "len", len(s), 1000)
    //fmt.Println(s)
    s.Clear()
    AssertEquals(t, "len", len(s), 0)
    
    s = FloatSlice{}
    s.Add(1)
    s.Insert(0, 2)
    s.Insert(1, 3)
    fmt.Println(s)
    AssertEquals(t, "len", len(s), 3)
    AssertStringEquals(t, "s", s, "[2 3 1]")
}

func TestFloatSliceRemove(t *testing.T) {
    fmt.Println("== Begin TestFloatSliceRemove...");
    defer fmt.Println("== End TestFloatSliceRemove.");
    var s FloatSlice
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

func BenchmarkFloatSliceInsert(b *testing.B) {
    b.StopTimer()
    s := make(FloatSlice, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s.Insert(1, float64(i))
    } // for i
}

func BenchmarkFloatSliceInsertByAppend(b *testing.B) {
    b.StopTimer()
    s := make([]float64, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s = append(s[:1], append([]float64{float64(i)}, s[1:]...)...)
    } // for i
}

func BenchmarkFloatSliceInsertByCopy(b *testing.B) {
    b.StopTimer()
    s := make([]float64, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s = append(s, 0)
        copy(s[2:], s[1:])
        s[1] = float64(i)
    } // for i
}
  
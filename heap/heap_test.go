package heap

import(
    "testing"
    "fmt"
    "runtime"
    "strings"
    "sort"
    "math/rand"
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

type IntHeap []int
func (h *IntHeap) Pop() int {
    PopToLast(sort.IntSlice(*h))
    res := (*h)[len(*h) - 1]
    *h = (*h)[:len(*h) - 1]
    
    return res
}

func (h *IntHeap) Push(x int) {
    *h = append(*h, x)
    PushLast(sort.IntSlice(*h))
}

type Data struct {
    Value string
    Priority int
}

func TestIntHeap(t *testing.T) {
    defer __(o_())
    
    var h IntHeap
    
    for i := 0; i < 1000; i ++ {
        h.Push(rand.Int())
    } // for i
    
    AssertEquals(t, "len(h)", len(h), 1000)

    peek := h[0]
    last := h.Pop()
    AssertEquals(t, "h.Peek()", peek, last)
//    fmt.Println(h)
    for i := 1; i < 1000; i ++ {
        cur := h.Pop()
        if cur < last {
            t.Errorf("%d should be larger than %d", cur, last)
        } // if
        last = cur
    } // for i
    fmt.Println(h)
}


type DataHeap []Data

func (h DataHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h DataHeap) Len() int {
    return len(h)
}

func (h DataHeap) Less(i, j int) bool {
    return h[i].Priority < h[j].Priority
}

func (h *DataHeap) Pop() Data {
    PopToLast(h)
    res := (*h)[len(*h) - 1]
    *h = (*h)[:len(*h) - 1]
    
    return res
}

func (h *DataHeap) Push(x Data) {
    *h = append(*h, x)
    PushLast(h)
}

func TestDataHeap(t *testing.T) {
    defer __(o_())
    
    var h DataHeap
    
    for i := 0; i < 1000; i ++ {
        h.Push(Data{"A", rand.Int()})
    } // for i
    
    AssertEquals(t, "len(h)", len(h), 1000)

    peek := h[0]
    last := h.Pop()
    AssertEquals(t, "h.Peek()", peek, last)
//    fmt.Println(h)
    for i := 1; i < 1000; i ++ {
        cur := h.Pop()
        if cur.Priority < last.Priority {
            t.Errorf("%d should be larger than %d", cur, last)
        } // if
        last = cur
    } // for i
    fmt.Println(h)
}

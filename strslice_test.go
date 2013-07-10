package villa

import (
	"fmt"
	"testing"

//    "strings"
)

func TestStringSlice(t *testing.T) {
	defer __(o_(t))

	var s StringSlice
	for i := 0; i < 1000; i++ {
		s.Add(string('A' + i))
	}

	AssertEquals(t, "len(s)", len(s), 1000)
	s.Clear()
	AssertEquals(t, "len(s)", len(s), 0)

	s = StringSlice{}
	s.Add("E", "B")
	s.Insert(1, "C", "D")
	AssertEquals(t, "len(s)", len(s), 4)
	AssertStringEquals(t, "s", s, "[E C D B]")
}

func ExampleStringSlice_direct() {
	type A struct {
		B string
		C int
	}

	var s StringSlice
	s.Add(10, "B", 30)
	fmt.Println(s)
	s.InsertSlice(len(s), []A{A{"E", 60}, A{"G", 80}})
	fmt.Println(s)
	s.Insert(1, "D", "E")
	fmt.Println(s)
	s.Swap(1, len(s)-1)
	fmt.Println(s)
	s.RemoveRange(1, 3)
	fmt.Println(s)
	s.Fill(0, len(s), "EE")
	fmt.Println(s)
	s.Clear()
	fmt.Println(s)
	// Output:
	// [10 B 30]
	// [10 B 30 {E 60} {G 80}]
	// [10 D E B 30 {E 60} {G 80}]
	// [10 {G 80} E B 30 {E 60} D]
	// [10 B 30 {E 60} D]
	// [EE EE EE EE EE]
	// []
}

func ExampleStringSlice_typecnv() {
	type A struct {
		B string
		C int
	}

	var s []string
	(*StringSlice)(&s).Add(10, "B", 30)
	fmt.Println(s)
	(*StringSlice)(&s).InsertSlice(len(s), []A{A{"E", 60}, A{"G", 80}})
	fmt.Println(s)
	(*StringSlice)(&s).Insert(1, "D", "E")
	fmt.Println(s)
	StringSlice(s).Swap(1, len(s)-1)
	fmt.Println(s)
	(*StringSlice)(&s).RemoveRange(1, 3)
	fmt.Println(s)
	StringSlice(s).Fill(0, len(s), "EE")
	fmt.Println(s)
	(*StringSlice)(&s).Clear()
	fmt.Println(s)
	// Output:
	// [10 B 30]
	// [10 B 30 {E 60} {G 80}]
	// [10 D E B 30 {E 60} {G 80}]
	// [10 {G 80} E B 30 {E 60} D]
	// [10 B 30 {E 60} D]
	// [EE EE EE EE EE]
	// []
}

func TestStringSliceRemove(t *testing.T) {
	defer __(o_(t))

	var s StringSlice
	s.Add("A", "B", "C", "D", "E", "F", "G")
	AssertEquals(t, "len(s)", len(s), 7)
	AssertStringEquals(t, "s", s, "[A B C D E F G]")

	s.RemoveRange(2, 5)
	AssertEquals(t, "len(s)", len(s), 4)
	AssertStringEquals(t, "s", s, "[A B F G]")

	s.Remove(2)
	AssertEquals(t, "len(s)", len(s), 3)
	AssertStringEquals(t, "s", s, "[A B G]")
}

func TestStringSliceEquals(t *testing.T) {
	s := StringSlice([]string{"1", "2", "3", "4"})

	AssertEquals(t, "s.Equals(nil)", s.Equals(nil), false)
	AssertEquals(t, "s.Equals([1, 2, 3, 4])", s.Equals([]string{"1", "2", "3", "4"}), true)
	AssertEquals(t, "s.Equals([1, 2, 5, 4])", s.Equals([]string{"1", "2", "5", "4"}), false)
	AssertEquals(t, "s.Equals([1, 2, 3, 4, 5])", s.Equals([]string{"1", "2", "3", "4", "5"}), false)

	AssertEquals(t, "nil.Equals([]int{})", StringSlice(nil).Equals(s[:0]), true)
	AssertEquals(t, "nil.Equals([]int{1, 2})", StringSlice(nil).Equals([]string{"1", "2"}), false)
}

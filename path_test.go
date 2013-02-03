package villa

import(
    "testing"
    "path/filepath"
)

func TestPath(t *testing.T) {
    p := Path("/")
    AssertStringEquals(t, "p.Join(abc)", p.Join("abc"), filepath.Join(string(p), "abc"))
	
	p = "abc.efg"
    AssertStringEquals(t, "p.Exe()", p.Ext(), filepath.Ext(string(p)))
}

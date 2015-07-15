package villa

import (
	"path/filepath"
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestPath(t *testing.T) {
	p := Path("/")
	assert.StringEqual(t, "p.Join(abc)", p.Join("abc"), filepath.Join(string(p), "abc"))

	p = "abc.efg"
	assert.StringEqual(t, "p.Exe()", p.Ext(), filepath.Ext(string(p)))
}

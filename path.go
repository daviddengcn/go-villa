package villa

import (
	"os"
	"path/filepath"
	"io/ioutil"
)

// Path is a wrapper for a path in the OS
type Path string

func (p Path) Join(elem ...interface{}) Path {
	els := make(StringSlice, 0, len(elem) + 1)
	els.Add(p).Add(elem...)
	return Path(filepath.Join(els...))
}

func (p Path) Mkdir(perm os.FileMode) error {
	return os.Mkdir(string(p), perm)
}

func (p Path) MkdirAll(perm os.FileMode) error {
	return os.MkdirAll(string(p), perm)
}

func (p Path) Ext() string {
	return filepath.Ext(string(p))
}

func (p Path) S() string {
	return string(p)
}

func (p Path) ReadDir() (fi []os.FileInfo, err error) {
	return ioutil.ReadDir(string(p))
}

func (p Path) Stat() (fi os.FileInfo, err error) {
	return os.Stat(string(p))
}

func (p Path) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(string(p))
}

func (p Path) WriteFile(data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(string(p), data, perm)
}

func (p Path) Symlink(dst Path) error {
	return os.Symlink(string(p), string(dst))
}

func (p Path) TempDir(prefix string) (name Path, err error) {
	nm, err := ioutil.TempDir(string(p), prefix)
	return Path(nm), err
}

func (p Path) Abs() (pth Path, err error) {
	pt, err := filepath.Abs(string(p))
	return Path(pt), err
}
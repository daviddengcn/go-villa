package villa

import (
	"os"
	"path/filepath"
)

// Path is a wrapper for a path in the OS
type Path string

func (p Path) Join(elem ...string) Path {
	return Path(filepath.Join(append([]string{string(p)}, elem...)...))
}

func (p Path) Mkdir(perm os.FileMode) error {
	return os.Mkdir(string(p), perm)
}

func (p Path) MkdirAll(perm os.FileMode) error {
	return os.MkdirAll(string(p), perm)
}


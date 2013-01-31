package villa

import "path/filepath"

// Path is a wrapper for a path in the OS
type Path string

func (p Path) Join(elem ...string) Path {
    return Path(filepath.Join(append([]string{string(p)}, elem...)...))
}
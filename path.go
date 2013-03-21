package villa

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// Path is a wrapper for a path in the OS.
// Some commonly used functions are wrapped as methods of Path, 
// and results, if any, are converted back to Path
type Path string

// Join connect elems to the tail of path
func (p Path) Join(elem ...interface{}) Path {
	els := make(StringSlice, 0, len(elem)+1)
	els.Add(p).Add(elem...)
	return Path(filepath.Join(els...))
}

// Exists checks whether the file exists
func (p Path) Exists() bool {
	_, err := p.Stat()
	return err == nil
}

// S converts Path back to string. This is sometimes more concise than string(p)
func (p Path) S() string {
	return string(p)
}

/*
	wrappers of filepath package
*/

// Abs is a wrapper to filepath.Abs
func (p Path) Abs() (pth Path, err error) {
	pt, err := filepath.Abs(string(p))
	return Path(pt), err
}

// Base returns the last element of path.
// Trailing path separators are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of separators, Base returns a single separator.
func (p Path) Base() Path {
	return Path(filepath.Base(string(p)))
}

// Clean is a wrapper to filepath.Clean
func (p Path) Clean() Path {
	return Path(filepath.Clean(string(p)))
}

// Dir is a wrapper to filepath.Dir
func (p Path) Dir() Path {
	return Path(filepath.Dir(string(p)))
}

// EvalSymlinks is a wrapper to filepath.EvalSymlinks
func (p Path) EvalSymlinks() (Path, error) {
	pt, err := filepath.EvalSymlinks(string(p))
	return Path(pt), err
}

// Ext is a wrapper to filepath.Ext
func (p Path) Ext() string {
	return filepath.Ext(string(p))
}

// FromSlash is a wrapper to filepath.FromSlash
func (p Path) FromSlash() Path {
	return Path(filepath.FromSlash(string(p)))
}

// IsAbs is a wrapper to filepath.IsAbs
func (p Path) IsAbs() bool {
	return filepath.IsAbs(string(p))
}

// Rel is a wrapper to filepath.Rel
func (p Path) Rel(targetpath Path) (Path, error) {
	rel, err := filepath.Rel(string(p), string(targetpath))
	return Path(rel), err
}

// Split is a wrapper to filepath.Split
func (p Path) Split() (dir, file Path) {
	d, f := filepath.Split(string(p))
	return Path(d), Path(f)
}

// SplitList is a wrapper to filepath.SplitList
func (p Path) SplitList() (lst []Path) {
	l := filepath.SplitList(string(p))
	lst = make([]Path, len(l))
	for i, el := range l {
		lst[i] = Path(el)
	}
	return
}

// ToSlash is a wrapper to filepath.ToSlash
func (p Path) ToSlash() string {
	return filepath.ToSlash(string(p))
}

// VolumeName is a wrapper to filepath.VolumeName
func (p Path) VolumeName() string {
	return filepath.VolumeName(string(p))
}

// WalkFunc is a wrapper to filepath.WalkFunc
type WalkFunc func(path Path, info os.FileInfo, err error) error

// Ext is a wrapper to filepath.Walk
func (p Path) Walk(walkFn WalkFunc) error {
	return filepath.Walk(string(p), func(path string, info os.FileInfo, err error) error {
		return walkFn(Path(path), info, err)
	})
}

/*
	wrappers of os package
*/

// Create is a wrapper to os.Create
func (p Path) Create() (file *os.File, err error) {
	return os.Create(string(p))
}

// Open is a wrapper to os.Open.
// It opens the named file for reading. If successful, methods on the returned 
// file can be used for reading; the associated file descriptor has mode O_RDONLY. 
// If there is an error, it will be of type *PathError.
func (p Path) Open() (file *os.File, err error) {
	return os.Open(string(p))

}

// OpenFile is a wrapper to os.OpenFile.
// It is the generalized open call; most users will use Open or Create instead. 
// It opens the named file with specified flag (O_RDONLY etc.) and perm, 
// (0666 etc.) if applicable. If successful, methods on the returned File can be 
// used for I/O. If there is an error, it will be of type *PathError.
func (p Path) OpenFile(flag int, perm os.FileMode) (file *os.File, err error) {
	return os.OpenFile(string(p), flag, perm)
}

// Mkdir is a wrappter to os.Mkdir
func (p Path) Mkdir(perm os.FileMode) error {
	return os.Mkdir(string(p), perm)
}

// MkdirAll is a wrappter to os.MkdirAll
func (p Path) MkdirAll(perm os.FileMode) error {
	return os.MkdirAll(string(p), perm)
}

// Remove is a wrappter to os.Remove
func (p Path) Remove() error {
	return os.Remove(string(p))
}

// RemoveAll is a wrappter to os.RemoveAll
func (p Path) RemoveAll() error {
	return os.RemoveAll(string(p))
}

// Rename is a wrappter to os.Rename
func (p Path) Rename(newname Path) error {
	return os.Rename(string(p), string(newname))
}

// Stat is a wrappter to os.Stat
func (p Path) Stat() (fi os.FileInfo, err error) {
	return os.Stat(string(p))
}

// Symlink is a wrappter to os.Symlink
func (p Path) Symlink(dst Path) error {
	return os.Symlink(string(p), string(dst))
}

/*
	wrappers of ioutil package
*/

// ReadDir is a wrappter to ioutil.ReadDir
func (p Path) ReadDir() (fi []os.FileInfo, err error) {
	return ioutil.ReadDir(string(p))
}

// ReadFile is a wrappter to ioutil.ReadFile
func (p Path) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(string(p))
}

// WriteFile is a wrappter to ioutil.WriteFile
func (p Path) WriteFile(data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(string(p), data, perm)
}

// TempDir is a wrappter to ioutil.TempDir
func (p Path) TempDir(prefix string) (name Path, err error) {
	nm, err := ioutil.TempDir(string(p), prefix)
	return Path(nm), err
}

/*
	wrapppers of exec package
*/

// Command is a wrappter to exec.Command
func (p Path) Command(arg ...string) *exec.Cmd {
	return exec.Command(string(p), arg...)
}

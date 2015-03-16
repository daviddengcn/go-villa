package villa

import(
	"fmt"
	"log"
	"path"
	"runtime"
)

// LinePos returns a string of format "filename:pos" of a position in the source.
// For skip == 0, it returns the position of the caller (the line call LinePos).
func LinePos(skip int) string {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", path.Base(file), line)
}

// Fatalf is similar to log.Fatalf with caller's position in the source
func Fatalf(format string, v ...interface{}) {
	log.Fatal(LinePos(1), ": ", fmt.Sprintf(format, v...))
}

// Fatalf is similar to log.Fatal with caller's position in the source
func Fatal(v ...interface{}) {
	log.Fatal(LinePos(1), ": ", fmt.Sprint(v...))
}
package file

import (
	"os"
	"path/filepath"
	"strings"
)

func GetDir() string {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

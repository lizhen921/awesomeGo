package path

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetDir() string {
	file := filepath.Dir(os.Args[0])
	fmt.Println(file)
	path, _ := filepath.Abs(file)
	fmt.Println(path)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	fmt.Println(ret)
	return ret
}

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	fmt.Println(file)
	path, _ := filepath.Abs(file)
	fmt.Println(path)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	fmt.Println(ret)
	return ret
}

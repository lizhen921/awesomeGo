package path

import (
	"testing"
)

func TestGetDir(t *testing.T) {
	//go 编译文件时生成的一个路径，大多代码所在位置都不是程序运行的位置
	GetDir()
}

func TestGetAppPath(t *testing.T) {
	GetAppPath()
}
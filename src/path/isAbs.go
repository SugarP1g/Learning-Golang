package main

import (
	"fmt"
	"path"
)

func main() {
	path1 := "/root/test"
	path2 := "/www/run/../../etc/passwd"
	path3 := "./test"

	// 检查路径是否为绝对路径。
	fmt.Println("path1 is abs: ", path.IsAbs(path1))  // true
	fmt.Println("path2 is abs: ", path.IsAbs(path2))  // true
	fmt.Println("path3 is abs: ", path.IsAbs(path3))  // false
}

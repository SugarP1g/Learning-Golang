package main

import (
	"fmt"
	"path"
)

func main() {
	filepath := "/www/html/test/test.sh"
	fmt.Println("[+] filename: ", filepath)
	// 扩展名为路径最后一个由斜杠分割的元素的最后一个点之后； 如果路径的这个部分没有点， 那么它的扩展名为空。
	fmt.Println("[-] ext: ", path.Ext(filepath))
	fmt.Println()
}

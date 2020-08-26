package main

import (
	"fmt"
	"path"
)

func main() {
	filepath := "/www/html/test/test.sh"
	fmt.Println("[+] filename: ", filepath)
	// func Split(path string) (dir, file string)
	basePath, filename := path.Split(filepath)
	fmt.Println("[-] dir: ", basePath)
	fmt.Println("[-] file: ", filename)
	fmt.Println()

	// 根据路径的最后一个斜杠， 将路径划分为目录部分和文件名部分。
	filepath = "/www/html/test/../../../../etc/shadow"
	fmt.Println("[+] filename: ", filepath)
	basePath, filename = path.Split(filepath)
	fmt.Println("[-] dir: ", basePath)
	fmt.Println("[-] file: ", filename)
	fmt.Println()

	// 如果路径中不包含斜杠， 那么返回一个空的目录， 并将文件名部分设置为路径本身。
	filepath = "test.sh"
	fmt.Println("[+] filename: ", filepath)
	basePath, filename = path.Split(filepath)
	fmt.Println("[-] dir: ", basePath)
	fmt.Println("[-] file: ", filename)
	fmt.Println()
}

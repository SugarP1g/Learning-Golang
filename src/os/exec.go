package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("calc.exe")
	fmt.Println(cmd.Path)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

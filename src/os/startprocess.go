package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Environ()
	fmt.Println(env)
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	proc, err := os.StartProcess("/bin/ls", []string{"ls", "-alt"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", proc)
	err = proc.Kill()
	if err != nil {
		fmt.Println(err)
	}
}

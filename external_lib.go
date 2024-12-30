package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func init() {
	fmt.Printf("mylib2.init()\n")
}


func ExternalFunc() string {

	fmt.Printf("Initialization with External function: Executed even if not invoked\n")
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
	return ""
}

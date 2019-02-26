package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec(username string) {
	fmt.Println("start")
	cmd := exec.Command("bin/vmiss", "bin/a.out")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	verr := cmd.Start()
	/*cerr := exec.Command("gcc", "*.c", "-o", "bin/"+username).Run()
	if cerr != nil {
		panic("cannnot compile:" + cerr.Error())
	}*/
	user := exec.Command("bin/"+username, "bin/a.out")
	uerr := user.Run()
	cmd.Wait()
	if verr != nil {
		panic("error:" + uerr.Error())
	}
}

package main

import (
	"os/exec"
	"os"
	"syscall"
)


/**

linux命令
 */
func main()  {

	binary,lookErr := exec.LookPath("ls")

	if lookErr != nil {
		panic(lookErr)
	}

	args := [] string{"ls","-a","-1","-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary,args,env)
	if execErr != nil {
		panic(execErr)
	}
}
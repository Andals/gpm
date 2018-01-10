package shell

import (
	"fmt"
	"testing"
)

func TestRunCmd(t *testing.T) {
	result := RunCmd("ls -l")
	fmt.Println(result)
}

func TestRunAsUser(t *testing.T) {
	result := RunAsUser("ls -l", "root")
	fmt.Println(result)
}

func TestRsync(t *testing.T) {
	host := "127.0.0.1"
	sou := "/tmp/go/*"
	dst := "/tmp/go1"
	sshUser := "ligang"

	result := Rsync(host, sou, dst, "", sshUser, 3)
	fmt.Println(result)
}

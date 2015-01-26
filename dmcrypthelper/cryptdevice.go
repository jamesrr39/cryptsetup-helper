package dmcrypthelper

import (
	"fmt"
	"os"
	"os/exec"
)

func Open(path string, containerName string) {

	fmt.Println("Attempting to open " + path)
	cmd := exec.Command("cryptsetup", "luksOpen", path, containerName)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

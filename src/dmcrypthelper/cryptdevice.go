package dmcrypthelper

import (
	"fmt"
	"os"
	"os/exec"
)

var CRYPTSETUP_PATH string = "/sbin/cryptsetup"

func Open(containerPath string, containerName string) {

	fmt.Println("Attempting to open " + containerPath)
	cmd := exec.Command(CRYPTSETUP_PATH, "luksOpen", containerPath, containerName)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func Close(devicePath string) {
	fmt.Println("Attempting to close " + devicePath)
	cmd := exec.Command(CRYPTSETUP_PATH, "luksClose", devicePath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

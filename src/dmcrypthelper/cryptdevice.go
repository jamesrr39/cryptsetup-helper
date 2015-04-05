package dmcrypthelper

import (
	"fmt"
	"os"
	"os/exec"
)

func Open(folderPath string, containerName string) {

	fmt.Println("Attempting to open " + folderPath + "/" + containerName)
	cmd := exec.Command("cryptsetup", "luksOpen", folderPath+"/"+containerName, containerName)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func Close(devicePath string) {
	fmt.Println("Attempting to close " + devicePath)
	cmd := exec.Command("cryptsetup", "luksClose", devicePath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

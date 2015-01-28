package mounthelper

import (
	"fmt"
	"os"
	"os/exec"
)

func MountDevice(devicePath string, mountPath string) {
	fmt.Println("Attempting to mount " + devicePath)
	cmd := exec.Command("mount", devicePath, mountPath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error mounting " + devicePath + " at " + mountPath)
	}
}

func UnmountDevice(path string) {
	fmt.Println("Attempting to unmount " + path)
	cmd := exec.Command("umount", path)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error unmounting " + path)
	}
}

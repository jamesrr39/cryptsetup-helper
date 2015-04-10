package mounthelper

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func MountDevice(devicePath string, mountPath string) {
	_, statErr := os.Stat(mountPath)
	if os.IsNotExist(statErr) {
		fmt.Println("Creating a new directory to mount the folder in at " + mountPath)
		mkdirErr := os.Mkdir(mountPath, 0700)
		if mkdirErr != nil {
			log.Fatal("Couldn't create directory at " + mountPath)
		}
	}
	fmt.Println("Attempting to mount " + devicePath)
	cmd := exec.Command("mount", devicePath, mountPath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
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

package device

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

// mount a device on the mountpath
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

// path Path that the container is mounted upon
func UnmountDevice(path string) {
	fmt.Println("Attempting to unmount " + path)
	err := syscall.Unmount(path, 0)
	if err != nil {
		fmt.Println("Error unmounting " + path)
	} else {
		fmt.Println("Successfully unmounted " + path)
	}
}

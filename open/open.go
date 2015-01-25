package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var containerName string
	var containerFilePath string
	var mountPath string
	var deviceFolderPath string
	flag.StringVar(&containerFilePath, "fp", "", "filepath to the device to be opened")
	flag.StringVar(&containerName, "n", "", "name of the device to be opened")
	flag.StringVar(&mountPath, "mp", "", "path the device should be mounted at")
	flag.StringVar(&deviceFolderPath, "dfp", "/dev/mapper/", "folder the device file should be created in, including the path separator ")

	flag.Parse()

	OpenDMCryptContainer(containerFilePath, containerName)
	devicePath := deviceFolderPath + containerName
	MountDevice(devicePath, mountPath)

}

func OpenDMCryptContainer(path string, containerName string) {

	fmt.Println("Attempting to open " + path)
	cmd := exec.Command("cryptsetup", "luksOpen", path, containerName)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func MountDevice(devicePath string, mountPath string) {
	fmt.Println("Attempting to mount " + devicePath)
	_, mountError := exec.Command("mount", devicePath, mountPath).Output()
	if mountError != nil {
		fmt.Println("Error mounting the device")
	} else {
		fmt.Println("Device mounted successfully at " + mountPath)
	}
}

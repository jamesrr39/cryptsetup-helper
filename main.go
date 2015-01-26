package main

import (
	"dmcrypthelper"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	containerName     string
	containerFilePath string
	mountPath         string
	deviceFolderPath  string
	verbose           bool
)

func init() {
	flag.StringVar(&containerFilePath, "fp", "", "filepath to the device to be opened")
	flag.StringVar(&containerName, "n", "", "name of the device to be opened")
	flag.StringVar(&mountPath, "mp", "", "path the device should be mounted at")
	flag.StringVar(&deviceFolderPath, "dfp", "/dev/mapper/", "folder the device file should be created in, including the path separator ")
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.Parse()
}

func main() {
	dmcrypthelper.Open(containerFilePath, containerName)
	devicePath := deviceFolderPath + containerName
	MountDevice(devicePath, mountPath)

}

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

package main

import (
	"device"
	"dmcrypthelper"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"
)

var (
	containerPath    string
	mountFolder      string
	deviceFolderPath string
)

func init() {

	flag.StringVar(&containerPath, "c", "", "container to be opened")
	flag.StringVar(&mountFolder, "mp", "", "folder the device should be mounted at")
	flag.StringVar(&deviceFolderPath, "dfp", "/dev/mapper/", "folder the device file should be created in, including the path separator ")
	flag.Parse()
}

func main() {

	containerFragments := strings.Split(containerPath, string(os.PathSeparator))
	containerName := containerFragments[len(containerFragments)-1]

	dmcrypthelper.Open(containerPath, containerName)
	devicePath := deviceFolderPath + string(os.PathSeparator) + containerName
	mountPath := mountFolder + string(os.PathSeparator) //+ containerName

	device, err := device.FromPath(devicePath)
	if err != nil {
		panic("There was an error finding the device at " + devicePath + ".\n" + err.Error())
	}

	// todo mount flags?
	syscall.Mount(devicePath, mountPath, device.FSType, 0, "")
}

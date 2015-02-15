package main

import (
	"dmcrypthelper"
	"flag"
	"mounthelper"
)

var (
	containerName     string
	containerFilePath string
	mountPath         string
	deviceFolderPath  string
)

func init() {
	flag.StringVar(&containerFilePath, "fp", "", "filepath to the device to be opened")
	flag.StringVar(&containerName, "n", "", "name of the device to be opened")
	flag.StringVar(&mountPath, "mp", "", "path the device should be mounted at")
	flag.StringVar(&deviceFolderPath, "dfp", "/dev/mapper/", "folder the device file should be created in, including the path separator ")
	flag.Parse()
}

func main() {
	dmcrypthelper.Open(containerFilePath, containerName)
	devicePath := deviceFolderPath + containerName
	mounthelper.MountDevice(devicePath, mountPath)

}

package main

import (
	"dmcrypthelper"
	"flag"
	"log"
	"mounthelper"
	"os/user"
)

var (
	containerName    string
	containerFolder  string
	mountFolder      string
	deviceFolderPath string
)

func init() {
	var currentUser, err = user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// todo - cryptsetup and mount app paths?
	flag.StringVar(&containerFolder, "fp", currentUser.HomeDir+"/volumes/", "folder of the device to be opened")
	flag.StringVar(&containerName, "n", "", "name of the device to be opened")
	flag.StringVar(&mountFolder, "mp", "/mnt/"+currentUser.Username+"/", "folder the device should be mounted at")
	flag.StringVar(&deviceFolderPath, "dfp", "/dev/mapper/", "folder the device file should be created in, including the path separator ")
	flag.Parse()
}

func main() {
	dmcrypthelper.Open(containerFolder, containerName)
	devicePath := deviceFolderPath + "/" + containerName
	mountPath := mountFolder + "/" + containerName
	mounthelper.MountDevice(devicePath, mountPath)

}

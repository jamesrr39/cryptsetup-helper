package main

import (
	"device"
	"dmcrypthelper"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
)

type ProgramMode uint8

const (
	PROGRAM_MODE_OPEN  ProgramMode = 1
	PROGRAM_MODE_CLOSE ProgramMode = 2
)

var (
	containerPath string
	command       ProgramMode
	mountFolder   string
)

func init() {
	log.SetFlags(log.Lshortfile)

	if 0 == len(os.Args) {
		log.Fatalln("No program mode or container path specified")
	}
	switch os.Args[1] {
	case "open":
		command = PROGRAM_MODE_OPEN
	case "close":
		command = PROGRAM_MODE_CLOSE
	default:
		log.Fatalln("Program Mode not found")
	}
	if len(os.Args) > 2 {
		containerPath = strings.TrimRight(os.Args[2], string(os.PathSeparator))
	} else {
		log.Fatalln("No container path specified")
	}
	flag.StringVar(&mountFolder, "mount-folder", "/media", "folder devices should be mounted at")
}

func main() {
	switch command {
	case PROGRAM_MODE_OPEN:
		Open()
	case PROGRAM_MODE_CLOSE:
		Close()
	}
}

func Open() {

	volume, err := device.FromPath(containerPath)
	if nil != err {
		log.Fatalf("There was an error finding the device at %s. %s", containerPath, err.Error())
	}

	dmdevicePtr, err := dmcrypthelper.Open(volume)
	if nil != err {
		log.Fatalf("There was an error opening the device. Was the correct passphrase entered? %s", err.Error())
	}
	dmdevice := *dmdevicePtr

	mountPath := mountFolder + string(os.PathSeparator) + dmdevice.UUID
	if _, err := os.Stat(mountPath); nil != err {
		if os.IsNotExist(err) {
			os.Mkdir(mountPath, os.FileMode(0700))
		}
	}

	log.Printf("about to attempt to mount %s at %s as %s", dmdevice.DevicePath, mountPath, dmdevice.FSType)
	if err := syscall.Mount(dmdevice.DevicePath, mountPath, dmdevice.FSType, syscall.MS_MGC_VAL, ""); nil != err {
		log.Fatalln(err)
	}
}

func Close() {
	// unmount

	mountedDevice, err := device.FromMountPoint(containerPath)
	if nil != err {
		switch err.(type) {
		case device.NoDeviceMountedOnThisPathError:
			log.Fatalf("Couldn't find a device at %s. %s", containerPath, err)
		default:
			log.Fatalf("Unknown error find a device at %s. %s", containerPath, err)
		}
	}

	if err := syscall.Unmount(containerPath, syscall.MNT_DETACH); nil != err {
		log.Fatalf("Error unmounting container at '%s': %s", containerPath, err)
	}

	// close container
	if err := dmcrypthelper.Close(mountedDevice.DevicePath); nil != err {
		log.Fatalf("Error closing the device at %s with cryptsetup: %s", mountedDevice.DevicePath)
	}

	// delete mount directory
	if _, err := ioutil.ReadDir(containerPath); nil != err {
		log.Fatalf("Error finding the container path direcory (%s). %s", containerPath, err)
	}
	if err := os.Remove(containerPath); nil != err {
		log.Fatalf("Error deleting the container path directory (%s). %s", containerPath, err)
	}

}

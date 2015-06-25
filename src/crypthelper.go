package main

import (
	"device"
	"dmcrypthelper"
	"flag"
	"fmt"
	"log"
	"os"
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

	switch os.Args[1] {
	case "open":
		command = PROGRAM_MODE_OPEN
	case "close":
		command = PROGRAM_MODE_CLOSE
	default:
		log.Fatalln("Program Mode not found")
	}
	if len(os.Args) > 2 {
		containerPath = os.Args[2]
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
	if err != nil {
		log.Fatalln("There was an error finding the device at " + containerPath + ".\n" + err.Error())
	}

	dmdevicePtr, _ := dmcrypthelper.Open(volume)
	dmdevice := *dmdevicePtr

	mountPath := mountFolder + string(os.PathSeparator) + dmdevice.UUID
	if _, err := os.Stat(mountPath); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(mountPath, os.FileMode(0700))
		}
	}

	// todo mount flags
	fmt.Println("trying to mount " + dmdevice.DevicePath + " at " + mountPath + " as " + dmdevice.FSType)
	if err := syscall.Mount(dmdevice.DevicePath, mountPath, dmdevice.FSType, 0, ""); err != nil {
		log.Fatalln(err)
	}
}

func Close() {
	panic("todo")
}

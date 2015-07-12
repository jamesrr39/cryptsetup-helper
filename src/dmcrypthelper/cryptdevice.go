package dmcrypthelper

import (
	devicepkg "device"
	"log"
	"os"
	"os/exec"
)

// Path to cryptsetup executable
var CRYPTSETUP_PATH string = "/sbin/cryptsetup"

func Open(device *devicepkg.Device) (*devicepkg.Device, error) {

	log.Println("Attempting to open " + device.DevicePath)
	cmd := exec.Command(CRYPTSETUP_PATH, "luksOpen", device.DevicePath, device.UUID)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); nil != err {
		return nil, err
	}
	openedDevicePath := "/dev/mapper/" + device.UUID
	openedDevice, err := devicepkg.FromPath(openedDevicePath)
	if nil != err {
		return nil, err
	}
	return openedDevice, nil
}

func Close(devicePath string) error {

	log.Println("Attempting to close " + devicePath)
	cmd := exec.Command(CRYPTSETUP_PATH, "luksClose", devicePath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

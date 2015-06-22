package dmcrypthelper

import (
	devicepkg "device"
	"fmt"
	"os"
	"os/exec"
)

var CRYPTSETUP_PATH string = "/sbin/cryptsetup"

func Open(device *devicepkg.Device) {

	fmt.Println("Attempting to open " + device.DevicePath)
	cmd := exec.Command(CRYPTSETUP_PATH, "luksOpen", device.DevicePath, device.UUID)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func Close(devicePath string) {
	fmt.Println("Attempting to close " + devicePath)
	cmd := exec.Command(CRYPTSETUP_PATH, "luksClose", devicePath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

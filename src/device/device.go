package device

import (
	"errors"
	"os/exec"
	"strings"
)

type Device struct {
	DevicePath string
	UUID       string
	FSType     string
}

type BlkidError struct {
	DevicePath string
}

func (e *BlkidError) Error() string {
	return "Error running '/sbin/blkid " + e.DevicePath + "'. Is a device available at " + e.DevicePath + "?"
}

var DeviceNotFoundError = errors.New("Device not found")

func FromPath(devicePath string) (*Device, error) {

	cmd := exec.Command("/sbin/blkid", devicePath)
	stdout, err := cmd.Output()

	if err != nil {
		return nil, DeviceNotFoundError
	}

	deviceInfo := string(stdout[:(len(stdout) - 1)])

	return &Device{
		DevicePath: devicePath,
		UUID:       GetUUIDFromBlkidOutput(deviceInfo),
		FSType:     GetFSTypeFromBlkidOutput(deviceInfo),
	}, nil
}

func GetUUIDFromBlkidOutput(blkidOutput string) string {
	const NoReplaceLimit = -1
	blkidOutputFragments := strings.Split(blkidOutput, " ")
	return strings.Replace(strings.Split(blkidOutputFragments[1], "=")[1], "\"", "", NoReplaceLimit)
}

func GetFSTypeFromBlkidOutput(blkidOutput string) string {
	const NoReplaceLimit = -1
	blkidOutputFragments := strings.Split(blkidOutput, " ")
	return strings.Replace(strings.Split(blkidOutputFragments[2], "=")[1], "\"", "", NoReplaceLimit)
}

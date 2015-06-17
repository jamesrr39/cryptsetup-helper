package device

import (
	"os/exec"
	"strings"
)

type Device struct {
	DevicePath string
	UUID       string
	FSType     string
}

func (e DeviceNotFoundError) Error() string {
	return "Error running '/sbin/blkid " + e.DevicePath + "'. Is a device available at " + e.DevicePath + "?"
}

type DeviceNotFoundError struct {
	DevicePath string
}

func FromPath(devicePath string) (*Device, error) {

	cmd := exec.Command("/sbin/blkid", devicePath)
	stdout, err := cmd.Output()

	if err != nil {
		return nil, DeviceNotFoundError{
			DevicePath: devicePath,
		}
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

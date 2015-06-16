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

func FromPath(devicePath string) (*Device, error) {

	cmd := exec.Command("/sbin/blkid", devicePath)

	stdout, err := cmd.Output()

	if err != nil {
		return nil, errors.New("Error running /sbin/blkid")
	}

	// fixme error handling
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

package device

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUUIDFromBlkidOutput(t *testing.T) {
	const mockUUID = "01234567-89ab-cdef-0123-456789abcdef"
	testOutput := "/dev/mapper/test: UUID=\"" + mockUUID + "\" TYPE=\"ext2\""
	assert.Equal(t, mockUUID, GetUUIDFromBlkidOutput(testOutput), "should give the test device UUID")
}

func TestGetFSTypeFromBlkidOutput(t *testing.T) {
	const mockFSType = "ext2"
	testOutput := "/dev/mapper/test: UUID=\"01234567-89ab-cdef-0123-456789abcdef\" TYPE=\"" + mockFSType + "\""
	assert.Equal(t, mockFSType, GetFSTypeFromBlkidOutput(testOutput), "should give the test device filesystem type")
}

func TestFromDevice_NoDevice(t *testing.T) {
	const devicePath = "/dev/mapper/1234567890"
	const expectedErrorMessage = "Error running '/sbin/blkid " + devicePath + "'. Is a device available at " + devicePath + "?"
	_, err := FromPath(devicePath)
	assert.Equal(t, DeviceNotFoundError{DevicePath: devicePath}, err, "should error out on not finding a device")
}

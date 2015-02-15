package main

import (
	"dmcrypthelper"
	"flag"
	"mounthelper"
)

var (
	devicePath string
)

func init() {
	flag.StringVar(&devicePath, "dp", "", "where the device file is located")
	flag.Parse()
}

func main() {
	mounthelper.UnmountDevice(devicePath)
	dmcrypthelper.Close(devicePath)
}

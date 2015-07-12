package userhelper

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func GetLognameUser() string {
	cmd := exec.Command("logname")
	cmd.Stderr = os.Stderr
	reader, writer, _ := os.Pipe()
	originalStdout := os.Stdout
	//os.Stdout = writer
	cmd.Stdout = os.Stdout
	cmd.Run()

	channelOut := make(chan string)
	go func() {
		var buffer bytes.Buffer
		io.Copy(&buffer, reader)
		channelOut <- buffer.String()
	}()

	writer.Close()
	os.Stdout = originalStdout
	stringOut := <-channelOut

	fmt.Print("I am " + stringOut)

	return stringOut
}

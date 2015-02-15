## cryptsetup-helper

Helper programs for cryptsetup, to make opening and closing volumes easier.

Written in Go.

### Prequisites
* cryptsetup
* go

### Build

    go build src/crypthelper-open.go
    go build src/crypthelper-close.go

### Install
From building, these programs can be used out of the box.
You may want to set the setuid bit to allow normal users using the programs.

    sudo chmod u+s crypthelper-open
    sudo chmod u+s crypthelper-close

If you are using ecryptfs, you need to move this program outside of the filesystem to use the setuid bit.

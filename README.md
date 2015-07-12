## cryptsetup-helper

Helper programs for cryptsetup, to make opening and closing volumes easier.

Written in Go.

### Requirements

* build
  * go
  * make
* runtime (the app shells out to a couple of command line programs)
  * /sbin/cryptsetup
  * /sbin/blkid


### Build

    make

### Install
build, and then to run:

    bin/crypthelper open|close container-path

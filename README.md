## cryptsetup-helper

Helper programs for cryptsetup, to make opening and closing volumes easier.

Written in Go.

### Prequisites
* cryptsetup
* go

### Build

    ./build.sh

### Install
From building, these programs can be used out of the box.
You may want to set the setuid bit to allow normal users using the programs.

    sudo chmod u+s crypthelper-open
    sudo chmod u+s crypthelper-close

If you are using a filesystem mounted with the nosuid option (ecryptfs forces this), you need to move this program outside of that filesystem to use the setuid bit. [More](http://unix.stackexchange.com/a/151227)

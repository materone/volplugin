package librbd

// #cgo LDFLAGS: -lrbd -lrados
// #include <rados/librados.h>
// #include <rbd/librbd.h>
// #include <stdlib.h>
// #include <errno.h>
// #include <string.h>
//
import "C"
import (
	"errors"
	"os/exec"
)

// Version returns the version of librbd.
func Version() (int, int, int) {
	var major, minor, extra C.int

	C.rbd_version(&major, &minor, &extra)
	return int(major), int(minor), int(extra)
}

func strerror(i C.int) error {
	return errors.New(C.GoString(C.strerror(-i)))
}

func modprobeRBD() error {
	return exec.Command("modprobe", "rbd").Run()
}

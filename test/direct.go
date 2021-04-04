package main

import (
	"fmt"
	evdev "github.com/gvalkov/golang-evdev"
	"os"
	"syscall"
	"unsafe"
)

var fd uintptr

var dev *evdev.InputDevice

const devnode = "/dev/input/event3"

func main() {

	devices, err := evdev.ListInputDevices("/dev/input/event*")
	if err != nil {
		fmt.Println(err)
	}
	for _, dev := range devices {
		fmt.Println(dev.Name, " ", dev.Fn)
	}

	file, err := os.Open(devnode)
	check(err)

	fd = file.Fd()

	_, _, err = syscall.Syscall(syscall.SYS_IOCTL, fd, evdev.EVIOCGRAB, 1)
	check(err)

	_, _, err = syscall.Syscall(syscall.SYS_IOCTL, fd, evdev.EVIOCGRAB, uintptr(unsafe.Pointer(nil)))
	check(err)

	some()
}

func some() {
	dev, err := evdev.Open(devnode)
	checkP(err)

	for {
		events, err := dev.Read()
		check(err)
		for _, event := range events {
			fmt.Println(event.String())
		}
	}

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func checkP(err error) {
	if err != nil {
		panic(err)
	}
}

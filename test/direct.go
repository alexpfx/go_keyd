package main

import (
	"fmt"
	"github.com/alexpfx/go_keyd/internal/input"
	evdev "github.com/gvalkov/golang-evdev"
)

var fd uintptr

var dev *evdev.InputDevice

const devnode = "/dev/input/event22"

func main() {
	devices, err := evdev.ListInputDevices("/dev/input/event*")
	if err != nil {
		fmt.Println(err)
	}

	for _, dev := range devices {
		fmt.Println(dev.Fn, " ", dev.File.Fd(), " ", dev.Name)
	}

	source := input.New(3)
	evChannel, err := source.Listen()
	if err != nil {
		panic(err)
	}
	for ev := range evChannel {
		fmt.Println(ev)
	}


	/*

	file, err := os.Open(devnode)
	check(err)


	fd = file.Fd()

	_, _, err = syscall.Syscall(syscall.SYS_IOCTL, fd, evdev.EVIOCGRAB, 1)
	check(err)

	_, _, err = syscall.Syscall(syscall.SYS_IOCTL, fd, evdev.EVIOCGRAB, uintptr(unsafe.Pointer(nil)))
	check(err)

	some()*/
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

package input

import (
	"fmt"
	evdev "github.com/gvalkov/golang-evdev"
	"os"
	"syscall"
)

const eviocgrab = uintptr(0x40044590)

func New(deviceCode int) Source {
	return inputFile{deviceId: deviceCode}
}

type Source interface {
	Listen() (chan Event, error)
}

type inputFile struct {
	deviceId int
}

func (i inputFile) Listen() (chan Event, error) {
	evChannel := make(chan Event)
	device, err := open(fmt.Sprintf("/dev/input/event%d", i.deviceId))
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			events, _ := device.Read()
			for _, e := range events {

				evChannel <- Event{
					EventType: EventType(e.Type),
					DeviceId:  0,
					Detail:    int(e.Code),
					Buttons:   0,
					Modifiers: 0,
				}

			}

			close(evChannel)
		}
	}()

	return nil, err
}


func release(dev *InputDevice) error{
	_, _, serr := syscall.RawSyscall(syscall.SYS_IOCTL, dev.fd.Fd(), eviocgrab, uintptr(0))
	err := convertErr(serr)
	if err != nil{
		return err
	}
	return nil
}
func grab(dev *InputDevice) error {
	_, _, errN := syscall.RawSyscall(syscall.SYS_IOCTL, dev.fd.Fd(), eviocgrab,  uintptr(1))
	err := convertErr(errN)
	if err != nil{
		return err
	}

	return nil
}

func convertErr(err syscall.Errno) error {
	if err != 0 {
		return fmt.Errorf("error: %d", err)
	}
	return nil
}

func open(devfile string) (*evdev.InputDevice, error) {
	fmt.Println(devfile)
	device, err := evdev.Open(devfile)
	check(err)
	_, err = os.Open(devfile)
	check(err)
	return device, err

}

func check(err error) {
	if err != nil{
		fmt.Println(err)
	}

}

type InputDevice struct {
	name string
	fd   *os.File
}

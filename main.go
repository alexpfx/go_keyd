package main

import (
	"fmt"
	"github.com/alexpfx/go_keyd/internal/devices"
	"github.com/alexpfx/go_keyd/internal/keymap"
	"regexp"
	"strings"
)

var lastNumber = regexp.MustCompile(`(\d+)$`)

func main() {
	keyboards := devices.FindKeyboards()

	deviceIds := make([]string, 0)

	for id := range keyboards {
		nameId := devices.GetDeviceName(id)
		name := strings.Split(nameId, "id=")[0]
		fmt.Println(strings.TrimSpace(name))
		deviceIds = append(deviceIds, id)
	}
	kmp := keymap.Load()

	monitor := devices.Monitor{
		Activate: nil,
		Escape:   nil,
		Keymap:   kmp,
	}

	ch := monitor.Start(deviceIds[0])

	for msg := range ch {
		fmt.Printf("type: %d id: %d name: %s\n", msg.Type, msg.KeyId, kmp.Get(msg.KeyId, 0))
	}

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

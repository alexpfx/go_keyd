package main

import (
	"fmt"
	monitor2 "github.com/alexpfx/go_keyd/internal/monitor"
	"regexp"
	"strconv"
	"strings"
)

var lastNumber = regexp.MustCompile(`(\d+)$`)

func main() {
	keyboards := monitor2.FindKeyboards()

	deviceIds := make([]string, 0)

	for id := range keyboards {
		nameId := monitor2.GetDeviceName(id)
		name := strings.Split(nameId, "id=")[0]
		fmt.Println(strings.TrimSpace(name))
		deviceIds = append(deviceIds, id)
	}
	/*kmp := keymap.Load()*/

	monitor := monitor2.InputMonitor{
		Filter: nil,
	}

	id, _ := strconv.Atoi(deviceIds[0])

	ch, cancel := monitor.Start(uint16(id))

	for msg := range ch {
		fmt.Println(msg)
	}

	cancel()
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

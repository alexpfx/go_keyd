package monitor

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func FindKeyboards() chan string {
	out := find(func(msg string) bool {
		return !strings.Contains(msg, "motion") && !strings.Contains(msg, "button")
	}, time.Second*2)

	return out
}

func GetDeviceName(id string) string {
	cmd := exec.Command("xinput", "--list", "--short", id)
	output, err := cmd.CombinedOutput()
	check(err)
	return string(output)
}

func find(accept func(msg string) bool, duration time.Duration) chan string {
	devices := getAllDevices()
	q := len(devices)

	och := make(chan string)
	doneCh := make(chan bool)

	for _, device := range devices {
		go spawnWithTimeout(device, duration, accept, och, doneCh)
	}

	go func() {
		for i := 0; i < q; i++ {
			<-doneCh
		}

		close(doneCh)
		close(och)
	}()

	return och
}


func getAllDevices() []string {
	ids := make([]string, 0)
	cmd := exec.Command("xinput", "--list", "--id-only")

	output, err := cmd.CombinedOutput()
	check(err)

	split := strings.Split(string(output), "\n")
	for _, s := range split {
		ids = append(ids, s)
	}

	return ids
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

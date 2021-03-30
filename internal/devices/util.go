package devices

import (
	"bufio"
	"context"
	"os/exec"
	"strings"
	"time"
)

func spawnUntilCancel(args [] string) (chan string, context.CancelFunc) {
	ch := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "xinput", args...)
	pipe, err := cmd.StdoutPipe()
	check(err)
	scanner := bufio.NewScanner(pipe)
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			ch <- strings.TrimSpace(text)
		}
	}()
	err = cmd.Start()

	return ch, cancel
}

func spawnWithTimeout(id string,
	duration time.Duration,
	accept func(msg string) bool,
	ch chan string, doneCh chan bool) {

	ctx, cancel := context.WithTimeout(context.Background(), duration)

	cmd := exec.CommandContext(ctx, "xinput", "test", id)

	pipe, err := cmd.StdoutPipe()
	check(err)

	scanner := bufio.NewScanner(pipe)
	go func() {
		for scanner.Scan() {
			t := scanner.Text()

			if accept(t) {
				ch <- id
			}
			cancel()
		}
	}()

	err = cmd.Run()

	doneCh <- true
}

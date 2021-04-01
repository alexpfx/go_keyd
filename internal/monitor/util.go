package monitor

import (
	"bufio"
	"bytes"
	"context"
	"log"
	"os/exec"
	"time"
)

var sep = []byte{'\n', '\n'}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if atEOF {
		return len(data), dropCR(data), nil
	}

	if i := bytes.Index(data, sep); i >= 0 {
		return i + len(sep), dropCR(data[0:i]), nil
	}
	return 0, nil, nil
}

func spawnUntilCancel(args []string) (chan []byte, context.CancelFunc) {
	ch := make(chan []byte, 1)
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "xinput", args...)
	pipe, err := cmd.StdoutPipe()
	check(err)
	scanner := bufio.NewScanner(pipe)
	scanner.Split(ScanLines)

	go func() {
		for scanner.Scan() {
			bs := []byte(scanner.Text())
			ch <- bs
		}
	}()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)

	}

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

package devices

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type EventType uint16

const (
	KeyPress   EventType = iota
	KeyRelease EventType = iota
)

type KeyListener struct {
	Devices []string
}

//RawMonitor inicia o processo que ouve as entradas do usuário e faz a filtragem básica
//transformando a saída em mensagem
type RawMonitor struct {
	Filter func(string) bool
}

type Event struct {
	Type  EventType
	KeyId uint16
}

type KeyStroke struct {
	Keys []uint16
}

func (m RawMonitor) monitor() (chan Event, context.CancelFunc) {
	ch, cancel := spawnUntilCancel([]string{"test", "xi2", "--root"})

	chEvent := make(chan Event)

	go func() {
		for {
			rawMsg := <-ch
			ev, err := parse(rawMsg)
			if err != nil {
				continue
			}
			chEvent <- ev
		}
	}()

	return chEvent, cancel
}

type Msg struct {
	
}
func parse(msg string) (Event, error) {
	fields := strings.Split(msg, "\n")
	if len(fields) < 1 {
		return Event{}, fmt.Errorf("acabando execucao")
	}



}

func parseKey(raw []string) Event {
	rawType := raw[1]
	code, _ := strconv.Atoi(raw[2])

	switch rawType {
	case "press":
		return Event{
			Type:  KeyPress,
			KeyId: uint16(code),
		}
	default:
		return Event{
			Type:  KeyRelease,
			KeyId: uint16(code),
		}

	}

}

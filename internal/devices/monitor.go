package devices

import (
	"fmt"
	"github.com/alexpfx/go_keyd/internal/keymap"
	"strconv"
	"strings"
)

type EventType uint16

const (
	KeyPress   EventType = iota
	KeyRelease EventType = iota
)


type Monitor struct {
	Activate []uint16
	Escape   []uint16
	Keymap   keymap.KeyMapper
}

func (m Monitor) Start(id string) chan Event {
	return m.monitor(id)
}

type Event struct {
	Type  EventType
	KeyId uint16
}

type KeyStroke struct {
	Keys [] uint16
}

func (m Monitor) monitor(id string) chan Event {
	ch, cancel := spawnUntilCancel(id)

	chEvent := make(chan Event)

	go func() {
		for {
			rawMsg := <-ch
			ev, err := parseRawMsg(rawMsg)
			if err != nil{
				fmt.Println(err)
				continue
			}
			chEvent <- ev
			if ev.Type == KeyRelease && ev.KeyId == 9{
				cancel()
				close(chEvent)
				close(ch)
			}
		}
	}()

	return chEvent
}

func parseRawMsg(msg string) (Event, error) {
	fields := strings.Fields(msg)
	if len(fields) < 1 {
		return Event{}, fmt.Errorf("acabando execucao")
	}
	evType := fields[0]
	switch evType {
	case "key":
		return parseKey(fields), nil
	default:
		return Event{}, fmt.Errorf("tipo de evento não tratado: %s", evType)
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

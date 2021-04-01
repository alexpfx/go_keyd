package monitor

import (
	"context"
	"fmt"
	"github.com/alexpfx/go_keyd/internal/xinput"
)

type KeyListener struct {
	Devices []string
}

//InputMonitor inicia o processo que ouve as entradas do usuário e faz a filtragem básica
//transformando a saída em eventos
type InputMonitor struct {
	Filter func(string) bool
}

func (m InputMonitor) Start(deviceId uint16) (chan xinput.Event, context.CancelFunc) {
	ch := make(chan xinput.Event)
	rawCh, cancelFunc := m.monitor()

	go func() {
		for r := range rawCh {
			fmt.Println(string(r.Payoff))
			ch <- xinput.Event{
				DeviceId:  0,
				Detail:    0,
				Modifiers: 0,
			}
		}
	}()

	return ch, cancelFunc
}

type KeyStroke struct {
	Keys []uint16
}

func (m InputMonitor) monitor() (chan xinput.RawEvent, context.CancelFunc) {
	ch, cancel := spawnUntilCancel([]string{"test-xi2", "--root"})

	chEvent := make(chan xinput.RawEvent)

	go func() {
		for {
			rawMsg := <-ch
			ev := xinput.Parse(rawMsg)
			if ev.EventType == xinput.Nothing {
				continue
			}
			chEvent <- ev
		}
	}()

	return chEvent, cancel
}

/*
func parseKey(raw []string) xinput.Event {
	rawType := raw[1]
	code, _ := strconv.Atoi(raw[2])

	switch rawType {
	case "press":
		return xinput.Event{
			Type:  KeyPress,
			KeyId: uint16(code),
		}
	default:
		return xinput.Event{
			Type:  KeyRelease,
			KeyId: uint16(code),
		}

	}

}*/

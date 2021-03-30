package monitor

import (
	"context"
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
	return m.monitor()
}

type KeyStroke struct {
	Keys []uint16
}

func (m InputMonitor) monitor() (chan xinput.Event, context.CancelFunc) {
	ch, cancel := spawnUntilCancel([]string{"test-xi2", "--root"})

	chEvent := make(chan xinput.Event)

	go func() {
		for {
			rawMsg := <-ch
			ev, err := xinput.Parse(rawMsg)
			if err != nil {
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

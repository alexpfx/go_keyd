package monitor

import (
	"context"
	"github.com/alexpfx/go_keyd/internal/xinput"
)

type KeyListener struct {
	Devices []string
}
type AcceptFunc func(event xinput.EventType) bool
type StopFunc func(event xinput.Event) bool

type KeyStroke struct {
	Keys []uint16
}

//Input inicia o processo que ouve as entradas do usuário e faz a filtragem básica
//transformando a saída em eventos
type Input struct {
	Accept AcceptFunc
	Stop   StopFunc
}

func NewKeyboardInput() Input {
	return Input{
		Accept: acceptKeyboard,
	}
}

func (m Input) Start() chan xinput.Event {
	ch := make(chan xinput.Event)
	rawCh, cancelFunc := m.rawInputMonitor()

	accept, stop := initFuncs(m)

	go func() {
		for rawEvent := range rawCh {
			if rawEvent.EventType == xinput.Nothing {
				continue
			}

			if !accept(rawEvent.EventType) {
				continue
			}

			ev := xinput.Parse(rawEvent)

			if stop(ev) {
				cancelFunc()
				close(rawCh)
				close(ch)
				continue
			}

			ch <- ev
		}
	}()

	return ch
}

func initFuncs(m Input) (AcceptFunc, StopFunc) {
	var accept = acceptAll
	var stop = stopWhenPress

	if m.Accept != nil {
		accept = m.Accept
	}
	if m.Stop != nil {
		stop = m.Stop
	}
	return accept, stop
}

var acceptKeyboard AcceptFunc = func(event xinput.EventType) bool {
	switch event {
	case xinput.KeyPress, xinput.KeyRelease:
		return true
	}
	return false
}

var acceptMotion AcceptFunc = func(eType xinput.EventType) bool {
	switch eType {
	case xinput.Motion,
		xinput.TouchBegin,
		xinput.TouchEnd,
		xinput.TouchUpdate:
		return true
	}
	return false
}

func (m Input) rawInputMonitor() (chan xinput.RawEvent, context.CancelFunc) {
	ch, cancel := spawnUntilCancel([]string{"test-xi2", "--root"})

	chEvent := make(chan xinput.RawEvent)

	go func() {
		for {
			rawMsg := <-ch
			ev := xinput.ParseRaw(rawMsg)
			if ev.EventType == xinput.Nothing {
				continue
			}
			chEvent <- ev
		}
	}()

	return chEvent, cancel
}

var acceptAll AcceptFunc = func(event xinput.EventType) bool {
	return true
}

var stopWhenPress StopFunc = func(event xinput.Event) bool {
	return false
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

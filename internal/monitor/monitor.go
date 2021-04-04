package monitor

import (
	"context"
	"github.com/alexpfx/go_keyd/internal/input"

)

type KeyListener struct {
	Devices []string
}
type AcceptFunc func(event input.EventType) bool
type StopFunc func(event input.Event) bool

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



func (m Input) Start() chan input.Event {
	ch := make(chan input.Event)
	rawCh, cancelFunc := m.rawInputMonitor()

	accept, stop := initFuncs(m)

	go func() {
		for rawEvent := range rawCh {
			if rawEvent.EventType == input.Nothing {
				continue
			}

			if !accept(rawEvent.EventType) {
				continue
			}

			ev := input.Parse(rawEvent)

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
	var stop = stopWhen

	if m.Accept != nil {
		accept = m.Accept
	}
	if m.Stop != nil {
		stop = m.Stop
	}
	return accept, stop
}

var acceptKeyboard AcceptFunc = func(event input.EventType) bool {
	switch event {
	case input.KeyPress, input.KeyRelease:
		return true
	}
	return false
}

var acceptMotion AcceptFunc = func(eType input.EventType) bool {
	switch eType {
	case input.Motion,
		input.TouchBegin,
		input.TouchEnd,
		input.TouchUpdate:
		return true
	}
	return false
}

func (m Input) rawInputMonitor() (chan input.RawEvent, context.CancelFunc) {
	ch, cancel := spawnUntilCancel([]string{"test-xi2", "--root"})

	chEvent := make(chan input.RawEvent)

	go func() {
		for {
			rawMsg := <-ch
			ev := input.ParseRaw(rawMsg)
			if ev.EventType == input.Nothing {
				continue
			}
			chEvent <- ev
		}
	}()

	return chEvent, cancel
}

var acceptAll AcceptFunc = func(event input.EventType) bool {
	return true
}

var stopWhen StopFunc = func(event input.Event) bool {


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

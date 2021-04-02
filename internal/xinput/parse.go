package xinput

import (
	"bytes"
	"regexp"
	"strconv"
)

var number = regexp.MustCompile(`(\d+)`)

var evTypeStr = []byte("EVENT type")

func Parse(rawEvent RawEvent) Event {
	return Event{
		EventType: rawEvent.EventType,
		DeviceId:  0,
		Detail:    0,
		Modifiers: 0,
	}
}

func ParseRaw(bs []byte) RawEvent {
	if len(bs) == 0 {
		return RawEvent{}
	}

	if !bytes.HasPrefix(bs, evTypeStr) {
		return RawEvent{}
	}

	if i := bytes.IndexByte(bs, ')'); i >= 0 {
		eventType := parse(bs[0 : i+1])
		if eventType == Nothing {
			return RawEvent{}
		}

		return RawEvent{eventType, bs[i+1:]}
	}
	return RawEvent{}

}

func parse(bs []byte) EventType {
	found := number.Find(bs)
	if found == nil {
		return Nothing
	}
	n, err := strconv.Atoi(string(found))
	if err != nil {
		return Nothing
	}

	return EventType(n)
}

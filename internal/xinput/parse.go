package xinput

import (
	"bytes"
	"regexp"
	"strconv"
)

var EmptyRawEvent = RawEvent{
	EventType: Nothing,
}
var evTypeStr = []byte("EVENT type")

func Parse(bs []byte) RawEvent {
	if len(bs) == 0 {
		return EmptyRawEvent
	}

	if !bytes.HasPrefix(bs, evTypeStr) {
		return EmptyRawEvent
	}

	if i := bytes.IndexByte(bs, ')'); i >= 0 {
		eventType := parse(bs[0 : i+1])
		if eventType == Nothing {
			return EmptyRawEvent
		}

		return RawEvent{eventType, bs[i:]}
	}
	return EmptyRawEvent

}

var number = regexp.MustCompile(`(\d+)`)

func parse(bs []byte) EventType {
	found := number.Find(bs)
	if found == nil {
		return Nothing
	}
	n, err := strconv.Atoi(string(found))
	if err != nil {
		return Nothing
	}
	if n == 2 {
		return KeyPress
	} else if n == 3 {
		return KeyRelease
	}
	return Nothing

}

package input

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var number = regexp.MustCompile(`(\d+)`)
var numberP = regexp.MustCompile(`\((\d+)\)`)
var hexa = regexp.MustCompile(`(?:0x)([A-Fa-f0-9]+)`)

var evTypeStr = []byte("EVENT type")

func Parse(rawEvent RawEvent) Event {
	payoff := rawEvent.Payoff
	fields := bytes.Split(payoff, []byte("\n    "))

	m := make(map[string]string)

	for _, by4Spaces := range fields {
		s := string(by4Spaces)
		if s == "" {
			continue
		}

		fieldValue := strings.SplitN(s, ":", 2)
		m[fieldValue[0]] = fieldValue[1]
	}
	device := parseDevice(m)
	detail := parseDetail(m)
	mod := parseModifier(m)

	event := Event{
		EventType: rawEvent.EventType,
		DeviceId:  device,
		Detail:    detail,
		Modifiers: mod,
	}
	fmt.Println(event)
	return event
}

func parseDetail(m map[string]string) int {
	nStr := number.FindStringSubmatch(m["detail"])[1]
	n, _ := strconv.Atoi(nStr)
	return n
}

func parseDevice(m map[string]string) int {
	nStr := numberP.FindStringSubmatch(m["deviceId"])[1]
	n, _ := strconv.Atoi(nStr)
	return n
}

func parseModifier(m map[string]string) int {
	str := m["modifiers"]
	mod := hexa.FindStringSubmatch(str)
	n, _ := strconv.ParseInt(mod[1], 16, 16)
	return int(n)
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

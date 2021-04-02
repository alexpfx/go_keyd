package xinput

type EventType int

const (
	Nothing = iota
	DeviceChanged
	KeyPress
	KeyRelease
	ButtonPress
	ButtonRelease
	Motion
	Enter
	Leave
	FocusIn
	FocusOut
	HierarchyChanged
	PropertyEvent
	RawKeyPress
	RawKeyRelease
	RawButtonPress
	RawButtonRelease
	RawMotion
	TouchBegin
	TouchUpdate
	TouchEnd
	TouchOwnership
	RawTouchBegin
	RawTouchUpdate
	RawTouchEnd
)

type Event struct {
	EventType EventType
	DeviceId  int
	Detail    uint16
	Modifiers uint16
}

var eventTypeMap = map[EventType]string{
	Nothing:          "Nothing",
	DeviceChanged:    "DeviceChanged",
	KeyPress:         "KeyPress",
	KeyRelease:       "KeyRelease",
	ButtonPress:      "ButtonPress",
	ButtonRelease:    "ButtonRelease",
	Motion:           "Motion",
	Enter:            "Enter",
	Leave:            "Leave",
	FocusIn:          "FocusIn",
	FocusOut:         "FocusOut",
	HierarchyChanged: "HierarchyChanged",
	PropertyEvent:    "PropertyEvent",
	RawKeyPress:      "RawKeyPress",
	RawKeyRelease:    "RawKeyRelease",
	RawButtonPress:   "RawButtonPress",
	RawButtonRelease: "RawButtonRelease",
	RawMotion:        "RawMotion",
	TouchBegin:       "TouchBegin",
	TouchUpdate:      "TouchUpdate",
	TouchEnd:         "TouchEnd",
	TouchOwnership:   "TouchOwnership",
	RawTouchBegin:    "RawTouchBegin",
	RawTouchUpdate:   "RawTouchUpdate",
	RawTouchEnd:      "RawTouchEnd",
}

func (r EventType) String() string {
	return r.GetName()
}

func (r EventType) GetName() string {
	n, ok := eventTypeMap[r]
	if !ok {
		return eventTypeMap[Nothing]
	}
	return n

}

type RawEvent struct {
	EventType EventType
	Payoff    []byte
}

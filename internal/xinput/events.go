package xinput

type EventType uint16

const (
	DeviceChanged    EventType = 1
	KeyPress         EventType = 2
	KeyRelease       EventType = 3
	ButtonPress      EventType = 4
	ButtonRelease    EventType = 5
	Motion           EventType = 6
	Enter            EventType = 7
	Leave            EventType = 8
	FocusIn          EventType = 9
	FocusOut         EventType = 10
	HierarchyChanged EventType = 11
	PropertyEvent    EventType = 12
	RawKeyPress      EventType = 13
	RawKeyRelease    EventType = 14
	RawButtonPress   EventType = 15
	RawButtonRelease EventType = 16
	RawMotion        EventType = 17
	TouchBegin       EventType = 18
	TouchUpdate      EventType = 19
	TouchEnd         EventType = 20
	TouchOwnership   EventType = 21
	RawTouchBegin    EventType = 22
	RawTouchUpdate   EventType = 23
	RawTouchEnd      EventType = 24
)

type Event struct {
	DeviceId  int
	Detail    uint16
	Modifiers uint16
}

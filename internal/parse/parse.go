package parse

type EventType uint16

const(
	KeyPress EventType = iota

)

type Event struct {

	DeviceId int
	Detail uint16
	Modifiers uint16
}

func Parse(str string){

}
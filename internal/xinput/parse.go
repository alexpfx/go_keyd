package xinput

import (
	"fmt"
)



func Parse(str string) (Event, error) {
	fmt.Println(str)

	return Event{
		DeviceId:  0,
		Detail:    0,
		Modifiers: 0,
	}, nil

}
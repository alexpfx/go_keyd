package main

import (
	"fmt"
)

func main() {



	/*keybMonitor := monitor.NewKeyboardInput()

	ch := keybMonitor.Start()

	for msg := range ch {
		fmt.Println(msg.EventType)
	}*/

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

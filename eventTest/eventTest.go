package main

import "C"
import (
	"fmt"
	"github.com/senro/robotgo"
)

//type Callback func(keycode C.int)

func callback (keycode int){
	fmt.Println("get keycode from robotgo:",keycode)
}

func main() {

	robotgo.AddEventListener()
	robotgo.AddCallback(callback)
}


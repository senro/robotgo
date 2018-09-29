package main

import "C"
import (
	"fmt"
	"github.com/senro/robotgo"
)

func callback (keycode int){
	fmt.Println("get keycode from robotgo:",keycode)
}

func main() {
	robotgo.AddReleasedCallback(callback)
	robotgo.AddEventListener()
}


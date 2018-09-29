// Copyright 2016 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-vgo/robotgo/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package robotgo

import "C"

type Callback func(keycode int)
var globalPressedCallback Callback
var globalReleasedCallback Callback
var globalTypedCallback Callback

func AddPressedCallback(callback Callback) {
	globalPressedCallback=callback
}
func AddReleasedCallback(callback Callback) {
	globalReleasedCallback=callback
}
func AddTypedCallback(callback Callback) {
	globalTypedCallback=callback
}

//export showPressedKeyCode
func showPressedKeyCode(keyCode C.int) int {
	//fmt.Println("show msg in go ",C.int(keyCode))
	globalPressedCallback(int(keyCode))
	//defer C.free(unsafe.Pointer(keyCode)) // will destruct in c
	return 1
}

//export showReleasedKeyCode
func showReleasedKeyCode(keyCode C.int) int {
	//fmt.Println("show msg in go ",C.int(keyCode))
	globalReleasedCallback(int(keyCode))
	//defer C.free(unsafe.Pointer(keyCode)) // will destruct in c
	return 1
}

//export showTypedKeyCode
func showTypedKeyCode(keyCode C.int) int {
	//fmt.Println("show msg in go ",C.int(keyCode))
	globalTypedCallback(int(keyCode))
	//defer C.free(unsafe.Pointer(keyCode)) // will destruct in c
	return 1
}
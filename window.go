package main

import (
	"fmt"

	"github.com/lxn/win"
)

func MoveChromeWindow() {
	fmt.Println("Finding Chrome window...")

	_ = win.HWND(0)
}

// type WindowPosition struct {
// 	X int32
// 	Y int32

// 	Width int32
// 	Height int32
// }

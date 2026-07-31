// package main

// import (
// 	"fmt"

// 	"github.com/lxn/win"
// )

// type ChromeWindow struct {
// 	Handle win.HWND
// 	Tittle string
// }

// func FindChromeWindows() ([]ChromeWindow, error) {
// 	var chromeWindows []ChromeWindow

// 	ok := win.EnumWindows(

// 		win.WNDENUMPROC(

// 			func(
// 				hwnd win.HWND,
// 				lparam win.LPARAM,
// 			) win.BOOL {
// 				fmt.Printf("Window Handle: %v\n", hwnd)

// 				// judul := win.Getwindowtext(hwnd)

// 				// if judul == "Google Chrome" {
// 				// 	chromeWindows = append(chromeWindows, ChromeWindow{
// 				// 		Handle: hwnd,
// 				// 		Tittle: judul,
// 				// 	})
// 				// }

// 				return win.TRUE
// 			},
// 		),

// 		win.LPARAM(0),
// 	)

// 	if !ok {
// 		fmt.Println("Gagal mengenumerasi jendela Chrome")

// 	}

// 	return chromeWindows, nil
// }

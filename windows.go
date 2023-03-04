//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"

	// "time"
	"unsafe"

	// "github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type MSG struct {
	HWND   uintptr
	UINT   uintptr
	WPARAM int16
	LPARAM int64
	DWORD  int32
	POINT  struct{ X, Y int64 }
}

func RegisterHotkey(a *App) {
	// key1, err1 := keys.Parse("Ctrl+Option+A") // KEY?

	user32 := syscall.MustLoadDLL("user32")
	defer user32.Release()
	reghotkey := user32.MustFindProc("RegisterHotKey")
	// peekmsg := user32.MustFindProc("PeekMessageW")
	getmsg := user32.MustFindProc("GetMessageW")

	// Hotkeys to listen to:
	keys := map[int16]*Hotkey{
		1: {1, ModAlt + ModShift, 'N'}, // ALT+SHIFT+N
		2: {2, ModAlt + ModShift, 'M'}, // ALT+SHIFT+M
		3: {3, ModAlt + ModCtrl, 'X'},  // ALT+CTRL+X
	}

	// Register hotkeys:
	for _, v := range keys {
		r1, _, err := reghotkey.Call(
			0, uintptr(v.Id), uintptr(v.Modifiers), uintptr(v.KeyCode))
		if r1 == 1 {
			fmt.Println("Registered", v)
		} else {
			fmt.Println("Failed to register", v, ", error:", err)
		}
	}

	for {
		var msg = &MSG{}
		getmsg.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0)
		// peekmsg.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0, 1)

		// Registered id is in the WPARAM field:
		if id := msg.WPARAM; id != 0 {
			fmt.Println("Hotkey pressed:", keys[id])
			if id == 1 {
				runtime.Show(a.ctx)
			}
			if id == 2 {
				runtime.Hide(a.ctx)
			}
			if id == 3 { // CTRL+ALT+X = Exit
				fmt.Println("CTRL+ALT+X pressed, goodbye...")
				runtime.Quit(a.ctx)
				return
			}
		}

		// time.Sleep(time.Millisecond * 50)
	}
}

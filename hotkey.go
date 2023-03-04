package main

import (
	"bytes"
	"fmt"
)

const (
    ModAlt = 1 << iota
    ModCtrl
    ModShift
    ModWin
)

type Hotkey struct {
    Id        int // Unique id
    Modifiers int // Mask of modifiers
    KeyCode   int // Key code, e.g. 'A'
}

func (h *Hotkey) String() string {
    mod := &bytes.Buffer{}
    if h.Modifiers&ModAlt != 0 {
        mod.WriteString("Alt+")
    }
    if h.Modifiers&ModCtrl != 0 {
        mod.WriteString("Ctrl+")
    }
    if h.Modifiers&ModShift != 0 {
        mod.WriteString("Shift+")
    }
    if h.Modifiers&ModWin != 0 {
        mod.WriteString("Win+")
    }
    return fmt.Sprintf("Hotkey[Id: %d, %s%c]", h.Id, mod, h.KeyCode)
}
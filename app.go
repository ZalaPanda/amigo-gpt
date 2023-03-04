package main

import (
	"context"
	"fmt"
	// "golang.design/x/hotkey"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	var app = &App{}
	go RegisterHotkey(app)
	// go fn()
	return app
}

// func fn() {
// 	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
// 	err := hk.Register()
// 	if err != nil {
// 		fmt.Printf("hotkey: failed to register hotkey: %v", err)
// 	}

// 	fmt.Printf("hotkey: %v is registered\n", hk)
// 	<-hk.Keydown()
// 	fmt.Printf("hotkey: %v is down\n", hk)
// 	<-hk.Keyup()
// 	fmt.Printf("hotkey: %v is up\n", hk)
// 	hk.Unregister()
// 	fmt.Printf("hotkey: %v is unregistered\n", hk)
// }

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

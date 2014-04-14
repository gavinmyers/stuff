package main

import (
	"fmt"
	"gopkg.in/qml.v0"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

type Player struct {
	qml.Object
}

func (r *Player) HandleClick(xPos, yPos int) {
  r.Set("x",r.Int("x") - 1)
  r.Set("y",r.Int("y") - 1)

}

func (r *Player) Paint(p *qml.Painter) {
  r.Set("targetX",0)
  r.Set("x",0)
  r.Set("targetY",0)
  r.Set("y",0)
}

func run() error {
	qml.Init(nil)

	qml.RegisterTypes("GoExtensions", 1, 0, []qml.TypeSpec{{
		Init: func(r *Player, obj qml.Object) { r.Object = obj },
	}})

	engine := qml.NewEngine()
	component, err := engine.LoadFile("painting.qml")
	if err != nil {
		return err
	}

	win := component.CreateWindow(nil)
	win.Show()
	win.Wait()

	return nil
}

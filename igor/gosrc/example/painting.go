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
type Game struct {
	qml.Object
}

func (r *Player) HandleClick(xPos, yPos int) {
	r.Set("x", r.Int("x")+40)
	r.Set("y", r.Int("y")+40)
}

func (r *Player) Update() {
}
func (r *Game) Update() {
}
func (r *Player) Paint(p *qml.Painter) {
}

func (r *Game) Paint(p *qml.Painter) {
}

func run() error {
	qml.Init(nil)

	qml.RegisterTypes("GoExtensions", 1, 0, []qml.TypeSpec{
		{Init: func(g *Game, obj qml.Object) { g.Object = obj }},
		{Init: func(g *Player, obj qml.Object) { g.Object = obj }},
	})

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

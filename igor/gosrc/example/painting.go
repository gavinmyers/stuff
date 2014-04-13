package main

import (
	"fmt"
	"gopkg.in/qml.v0"
	"gopkg.in/qml.v0/gl"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

type Tile struct {
	qml.Object
}

type Sprite struct {
	qml.Object
}

func (r *Tile) Paint(p *qml.Painter) {
	width := gl.Float(r.Int("width"))
	height := gl.Float(r.Int("height"))

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Color4f(1.0, 1.0, 1.0, 1.0)
	gl.Begin(gl.QUADS)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(width, 0)
	gl.Vertex2f(width, height)
	gl.Vertex2f(0, height)
	gl.End()

	gl.LineWidth(0.5)
	gl.Color4f(1.0, 1.0, 1.0, 1.0)
	gl.Begin(gl.LINES)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(width, height)
	gl.Vertex2f(width, 0)
	gl.Vertex2f(0, height)
	gl.End()
}

func run() error {
	qml.Init(nil)

	qml.RegisterTypes("GoExtensions", 1, 0, []qml.TypeSpec{{
		Init: func(r *Tile, obj qml.Object) { r.Object = obj },
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

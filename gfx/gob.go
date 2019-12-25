package gfx

import (
	"encoding/gob"

	"github.com/buchanae/ink/color"
	"github.com/buchanae/ink/dd"
)

func init() {
	gob.Register(Shader{})
	gob.Register(Layer{})
	gob.Register(color.RGBA{})
	gob.Register(dd.XY{})
	gob.Register(dd.Mesh{})
	gob.Register(dd.Rect{})
	gob.Register(dd.Quad{})
	gob.Register(dd.Triangle{})
	gob.Register(dd.Circle{})
	gob.Register(dd.Triangles{})
	gob.Register([]color.RGBA{})
	gob.Register([]dd.XY{})
}
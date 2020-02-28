// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"

	"github.com/hneemann/nplot"
	"github.com/hneemann/nplot/palette"
	"github.com/hneemann/nplot/vg"
	"github.com/hneemann/nplot/vg/draw"
)

// PaletteThumbnailers creates a slice of nplot.Thumbnailers that can be used to
// add legend entries for the colors in a color palette.
func PaletteThumbnailers(p palette.Palette) []nplot.Thumbnailer {
	colors := p.Colors()
	thumbnailers := make([]nplot.Thumbnailer, len(colors))
	for i, c := range colors {
		thumbnailers[i] = paletteThumbnailer{color: c}
	}
	return thumbnailers
}

// paletteThumbnailer implements the Thumbnailer interface
// for color palettes.
type paletteThumbnailer struct {
	color color.Color
}

// Thumbnail satisfies the nplot.Thumbnailer interface.
func (t paletteThumbnailer) Thumbnail(c *draw.Canvas) {
	pts := []vg.Point{
		{X: c.Min.X, Y: c.Min.Y},
		{X: c.Min.X, Y: c.Max.Y},
		{X: c.Max.X, Y: c.Max.Y},
		{X: c.Max.X, Y: c.Min.Y},
	}
	poly := c.ClipPolygonY(pts)
	c.FillPolygon(t.color, poly)
}

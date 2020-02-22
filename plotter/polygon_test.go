// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter_test

import (
	"image/color"
	"testing"

	"github.com/hneemann/plot"
	"github.com/hneemann/plot/cmpimg"
	"github.com/hneemann/plot/plotter"
	"github.com/hneemann/plot/vg"
	"github.com/hneemann/plot/vg/draw"
	"github.com/hneemann/plot/vg/recorder"
)

func TestPolygon_holes(t *testing.T) {
	cmpimg.CheckPlot(ExamplePolygon_holes, t, "polygon_holes.png", "polygon_holes.svg", "polygon_holes.pdf", "polygon_holes.eps")
}

func TestPolygon_hexagons(t *testing.T) {
	cmpimg.CheckPlot(ExamplePolygon_hexagons, t, "polygon_hexagons.png")
}

// This test ensures that the plotter doesn't panic if there are
// polygons wholly outside of the plotting range.
func TestPolygon_clip(t *testing.T) {
	poly, err := plotter.NewPolygon(
		plotter.XYs{{0, 0}, {1, 0}, {1, 1}, {0, 1}, {0, 0}},
	)
	if err != nil {
		t.Fatal(err)
	}
	poly.Color = color.Black // Give the polygon a color to fill.
	p, err := plot.New()
	if err != nil {
		t.Fatal(err)
	}

	// Set the plotting range so that the polygon is outside of it.
	p.X.Min = 2
	p.X.Max = 5

	p.Add(poly)
	c := new(recorder.Canvas)
	dc := draw.NewCanvas(c, vg.Centimeter, vg.Centimeter)
	p.Draw(dc) // If this does not panic, then the test passes.
}

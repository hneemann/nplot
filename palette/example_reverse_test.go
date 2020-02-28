// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package palette_test

import (
	"fmt"
	"log"

	"github.com/hneemann/nplot"
	"github.com/hneemann/nplot/palette"
	"github.com/hneemann/nplot/palette/moreland"
	"github.com/hneemann/nplot/plotter"
)

// This example creates a color bar and a second color bar where the
// direction of the colors are reversed.
func ExampleReverse() {
	p, err := nplot.New()
	if err != nil {
		log.Panic(err)
	}
	l := &plotter.ColorBar{ColorMap: moreland.Kindlmann()}
	l2 := &plotter.ColorBar{ColorMap: palette.Reverse(moreland.Kindlmann())}
	l.ColorMap.SetMin(0.5)
	l.ColorMap.SetMax(2.5)
	l2.ColorMap.SetMin(2.5)
	l2.ColorMap.SetMax(4.5)

	p.Add(l, l2)
	p.HideY()
	p.X.Padding = 0
	p.Title.Text = "A ColorMap and its Reverse"

	if err = p.Save(300, 48, "testdata/reverse.png"); err != nil {
		log.Panic(err)
	}
}

// This example creates a color palette from a reversed ColorMap.
func ExampleReverse_palette() {
	p, err := nplot.New()
	if err != nil {
		log.Panic(err)
	}
	thumbs := plotter.PaletteThumbnailers(palette.Reverse(moreland.Kindlmann()).Palette(10))
	for i, t := range thumbs {
		p.Legend.Add(fmt.Sprint(i), t)
	}
	p.HideAxes()
	p.X.Padding = 0
	p.Y.Padding = 0

	if err = p.Save(35, 120, "testdata/reverse_palette.png"); err != nil {
		log.Panic(err)
	}
}

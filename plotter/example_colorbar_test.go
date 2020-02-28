// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter_test

import (
	"image/color"
	"log"

	"github.com/hneemann/nplot"
	"github.com/hneemann/nplot/palette/moreland"
	"github.com/hneemann/nplot/plotter"
)

func ExampleColorBar_horizontal() {
	p, err := nplot.New()
	if err != nil {
		log.Panic(err)
	}
	l := &plotter.ColorBar{ColorMap: moreland.ExtendedBlackBody()}
	l.ColorMap.SetMin(0.5)
	l.ColorMap.SetMax(1.5)
	p.Add(l)
	p.HideY()
	p.X.Padding = 0
	p.Title.Text = "Title"

	if err = p.Save(300, 48, "testdata/colorBarHorizontal.png"); err != nil {
		log.Panic(err)
	}
}

// This example shows how to create a ColorBar on a log-transformed axis.
func ExampleColorBar_horizontal_log() {
	p, err := nplot.New()
	if err != nil {
		log.Panic(err)
	}
	colorMap, err := moreland.NewLuminance([]color.Color{color.Black, color.White})
	if err != nil {
		log.Panic(err)
	}
	l := &plotter.ColorBar{ColorMap: colorMap}
	l.ColorMap.SetMin(1)
	l.ColorMap.SetMax(100)
	p.Add(l)
	p.HideY()
	p.X.Padding = 0
	p.Title.Text = "Title"
	p.X.Scale = nplot.LogScale{}
	p.X.Tick.Marker = nplot.LogTicks{}

	if err = p.Save(300, 48, "testdata/colorBarHorizontalLog.png"); err != nil {
		log.Panic(err)
	}
}

func ExampleColorBar_vertical() {
	p, err := nplot.New()
	if err != nil {
		log.Panic(err)
	}
	l := &plotter.ColorBar{ColorMap: moreland.ExtendedBlackBody()}
	l.ColorMap.SetMin(0.5)
	l.ColorMap.SetMax(1.5)
	l.Vertical = true
	p.Add(l)
	p.HideX()
	p.Y.Padding = 0
	p.Title.Text = "Title"

	if err = p.Save(40, 300, "testdata/colorBarVertical.png"); err != nil {
		log.Panic(err)
	}
}

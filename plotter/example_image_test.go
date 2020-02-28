// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter_test

import (
	"image/png"
	"log"
	"os"

	"github.com/hneemann/nplot"
	"github.com/hneemann/nplot/plotter"
	"github.com/hneemann/nplot/vg"
)

// An example of embedding an image in a nplot.
func ExampleImage() {
	p, err := nplot.New()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	p.Title.Text = "A Logo"

	// load an image
	f, err := os.Open("testdata/image_plot_input.png")
	if err != nil {
		log.Fatalf("error opening image file: %v\n", err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatalf("error decoding image file: %v\n", err)
	}

	p.Add(plotter.NewImage(img, 100, 100, 200, 200))

	const (
		w = 5 * vg.Centimeter
		h = 5 * vg.Centimeter
	)

	err = p.Save(w, h, "testdata/image_plot.png")
	if err != nil {
		log.Fatalf("error saving image nplot: %v\n", err)
	}
}

// An example of embedding an image in a nplot with non-linear axes.
func ExampleImage_log() {
	p, err := nplot.New()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	p.Title.Text = "A Logo"

	// load an image
	f, err := os.Open("testdata/gopher.png")
	if err != nil {
		log.Fatalf("error opening image file: %v\n", err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatalf("error decoding image file: %v\n", err)
	}

	p.Add(plotter.NewImage(img, 100, 100, 10000, 10000))

	// Transform axes.
	p.X.Scale = nplot.LogScale{}
	p.Y.Scale = nplot.LogScale{}
	p.X.Tick.Marker = nplot.LogTicks{}
	p.Y.Tick.Marker = nplot.LogTicks{}

	const (
		w = 5 * vg.Centimeter
		h = 5 * vg.Centimeter
	)

	err = p.Save(w, h, "testdata/image_plot_log.png")
	if err != nil {
		log.Fatalf("error saving image nplot: %v\n", err)
	}
}

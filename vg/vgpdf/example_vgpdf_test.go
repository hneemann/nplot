// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vgpdf_test

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hneemann/nplot"
	"github.com/hneemann/nplot/plotter"
	"github.com/hneemann/nplot/vg"
	"github.com/hneemann/nplot/vg/draw"
	"github.com/hneemann/nplot/vg/vgpdf"
)

// Example_embedFonts shows how one can embed (or not) fonts inside
// a PDF nplot.
func Example_embedFonts() {
	p, err := nplot.New()
	if err != nil {
		log.Fatalf("could not create nplot: %v", err)
	}

	pts := plotter.XYs{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 1}}
	line, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatalf("could not create line: %v", err)
	}
	p.Add(line)
	p.X.Label.Text = "X axis"
	p.Y.Label.Text = "Y axis"

	c := vgpdf.New(100, 100)

	// enable/disable embedding fonts
	c.EmbedFonts(true)
	p.Draw(draw.New(c))

	f, err := os.Create("testdata/enable-embedded-fonts.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = c.WriteTo(f)
	if err != nil {
		log.Fatalf("could not write canvas: %v", err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalf("could not save canvas: %v", err)
	}
}

// Example_multipage shows how one can create a PDF with multiple pages.
func Example_multipage() {
	c := vgpdf.New(5*vg.Centimeter, 5*vg.Centimeter)

	for i, col := range []color.RGBA{{B: 255, A: 255}, {R: 255, A: 255}} {
		if i > 0 {
			// Add a new page.
			c.NextPage()
		}

		p, err := nplot.New()
		if err != nil {
			log.Fatalf("could not create nplot: %v", err)
		}

		pts := plotter.XYs{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 1}}
		line, err := plotter.NewLine(pts)
		if err != nil {
			log.Fatalf("could not create line: %v", err)
		}
		line.Color = col
		p.Add(line)
		p.Title.Text = fmt.Sprintf("Plot %d", i+1)
		p.X.Label.Text = "X axis"
		p.Y.Label.Text = "Y axis"

		// Write nplot to page.
		p.Draw(draw.New(c))
	}

	f, err := os.Create("testdata/multipage.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = c.WriteTo(f)
	if err != nil {
		log.Fatalf("could not write canvas: %v", err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalf("could not save canvas: %v", err)
	}
}

// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vgsvg_test

import (
	"log"

	"github.com/hneemann/nplot"
	"github.com/hneemann/nplot/plotter"
	"github.com/hneemann/nplot/vg"
)

func Example() {
	p, err := nplot.New()
	if err != nil {
		log.Fatalf("could not create nplot: %v", err)
	}
	p.Title.Text = "Scatter nplot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	scatter, err := plotter.NewScatter(plotter.XYs{{X: 1, Y: 1}, {X: 0, Y: 1}, {X: 0, Y: 0}})
	if err != nil {
		log.Fatalf("could not create scatter: %v", err)
	}
	p.Add(scatter)

	err = p.Save(5*vg.Centimeter, 5*vg.Centimeter, "testdata/scatter.svg")
	if err != nil {
		log.Fatalf("could not save SVG nplot: %v", err)
	}
}

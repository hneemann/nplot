// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob // import "github.com/hneemann/nplot/gob"

import (
	"encoding/gob"
	"image/color"

	"github.com/hneemann/nplot"
	"github.com/hneemann/nplot/plotter"
)

func init() {
	// register types for proper gob-encoding/decoding
	gob.Register(color.Gray16{})

	// nplot.Ticker
	gob.Register(nplot.ConstantTicks{})
	gob.Register(nplot.DefaultTicks{})
	gob.Register(nplot.LogTicks{})

	// nplot.Normalizer
	gob.Register(nplot.LinearScale{})
	gob.Register(nplot.LogScale{})

	// nplot.Plotter
	gob.Register(plotter.BarChart{})
	gob.Register(plotter.Histogram{})
	gob.Register(plotter.BoxPlot{})
	gob.Register(plotter.YErrorBars{})
	gob.Register(plotter.XErrorBars{})
	gob.Register(plotter.Function{})
	gob.Register(plotter.GlyphBoxes{})
	gob.Register(plotter.Grid{})
	gob.Register(plotter.Labels{})
	gob.Register(plotter.Line{})
	gob.Register(plotter.QuartPlot{})
	gob.Register(plotter.Scatter{})

	// plotter.XYZer
	gob.Register(plotter.XYZs{})
	gob.Register(plotter.XYValues{})

}

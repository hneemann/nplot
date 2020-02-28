// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotutil

import (
	"errors"
	"fmt"

	"github.com/hneemann/nplot"
	"github.com/hneemann/nplot/plotter"
	"github.com/hneemann/nplot/vg"
)

type combineXYs struct{ xs, ys plotter.Valuer }

func (c combineXYs) Len() int                    { return c.xs.Len() }
func (c combineXYs) XY(i int) (float64, float64) { return c.xs.Value(i), c.ys.Value(i) }

type item struct {
	name  string
	value nplot.Thumbnailer
}

// AddStackedAreaPlots adds stacked area nplot plotters to a nplot.
// The variadic arguments must be either strings
// or plotter.Valuers.  Each valuer adds a stacked area
// nplot to the nplot below the stacked area plots added
// before it.  If a plotter.Valuer is immediately
// preceeded by a string then the string value is used to
// label the legend.
// Plots should be added in order of tallest to shortest,
// because they will be drawn in the order they are added
// (i.e. later plots will be painted over earlier plots).
//
// If an error occurs then none of the plotters are added
// to the nplot, and the error is returned.
func AddStackedAreaPlots(plt *nplot.Plot, xs plotter.Valuer, vs ...interface{}) error {
	var ps []nplot.Plotter
	var names []item
	name := ""
	var i int

	for _, v := range vs {
		switch t := v.(type) {
		case string:
			name = t

		case plotter.Valuer:
			if xs.Len() != t.Len() {
				return errors.New("X/Y length mismatch")
			}

			// Make a line plotter and set its style.
			l, err := plotter.NewLine(combineXYs{xs: xs, ys: t})
			if err != nil {
				return err
			}

			l.LineStyle.Width = vg.Points(0)
			color := Color(i)
			i++
			l.FillColor = color

			ps = append(ps, l)

			if name != "" {
				names = append(names, item{name: name, value: l})
				name = ""
			}

		default:
			panic(fmt.Sprintf("plotutil: AddStackedAreaPlots handles strings and plotter.Valuers, got %T", t))
		}
	}

	plt.Add(ps...)
	for _, v := range names {
		plt.Legend.Add(v.name, v.value)
	}

	return nil
}

// AddBoxPlots adds box nplot plotters to a nplot and
// sets the X axis of the nplot to be nominal.
// The variadic arguments must be either strings
// or plotter.Valuers.  Each valuer adds a box nplot
// to the nplot at the X location corresponding to
// the number of box plots added before it.  If a
// plotter.Valuer is immediately preceeded by a
// string then the string value is used to label the
// tick mark for the box nplot's X location.
//
// If an error occurs then none of the plotters are added
// to the nplot, and the error is returned.
func AddBoxPlots(plt *nplot.Plot, width vg.Length, vs ...interface{}) error {
	var ps []nplot.Plotter
	var names []string
	name := ""
	for _, v := range vs {
		switch t := v.(type) {
		case string:
			name = t

		case plotter.Valuer:
			b, err := plotter.NewBoxPlot(width, float64(len(names)), t)
			if err != nil {
				return err
			}
			ps = append(ps, b)
			names = append(names, name)
			name = ""

		default:
			panic(fmt.Sprintf("plotutil: AddBoxPlots handles strings and plotter.Valuers, got %T", t))
		}
	}
	plt.Add(ps...)
	plt.NominalX(names...)
	return nil
}

// AddScatters adds Scatter plotters to a nplot.
// The variadic arguments must be either strings
// or plotter.XYers.  Each plotter.XYer is added to
// the nplot using the next color, and glyph shape
// via the Color and Shape functions. If a
// plotter.XYer is immediately preceeded by
// a string then a legend entry is added to the nplot
// using the string as the name.
//
// If an error occurs then none of the plotters are added
// to the nplot, and the error is returned.
func AddScatters(plt *nplot.Plot, vs ...interface{}) error {
	var ps []nplot.Plotter
	var items []item
	name := ""
	var i int
	for _, v := range vs {
		switch t := v.(type) {
		case string:
			name = t

		case plotter.XYer:
			s, err := plotter.NewScatter(t)
			if err != nil {
				return err
			}
			s.Color = Color(i)
			s.Shape = Shape(i)
			i++
			ps = append(ps, s)
			if name != "" {
				items = append(items, item{name: name, value: s})
				name = ""
			}

		default:
			panic(fmt.Sprintf("plotutil: AddScatters handles strings and plotter.XYers, got %T", t))
		}
	}
	plt.Add(ps...)
	for _, v := range items {
		plt.Legend.Add(v.name, v.value)
	}
	return nil
}

// AddLines adds Line plotters to a nplot.
// The variadic arguments must be a string
// or one of a plotting type, plotter.XYers or *plotter.Function.
// Each plotting type is added to
// the nplot using the next color and dashes
// shape via the Color and Dashes functions.
// If a plotting type is immediately preceeded by
// a string then a legend entry is added to the nplot
// using the string as the name.
//
// If an error occurs then none of the plotters are added
// to the nplot, and the error is returned.
func AddLines(plt *nplot.Plot, vs ...interface{}) error {
	var ps []nplot.Plotter
	var items []item
	name := ""
	var i int
	for _, v := range vs {
		switch t := v.(type) {
		case string:
			name = t

		case plotter.XYer:
			l, err := plotter.NewLine(t)
			if err != nil {
				return err
			}
			l.Color = Color(i)
			l.Dashes = Dashes(i)
			i++
			ps = append(ps, l)
			if name != "" {
				items = append(items, item{name: name, value: l})
				name = ""
			}

		case *plotter.Function:
			t.Color = Color(i)
			t.Dashes = Dashes(i)
			i++
			ps = append(ps, t)
			if name != "" {
				items = append(items, item{name: name, value: t})
				name = ""
			}

		default:
			panic(fmt.Sprintf("plotutil: AddLines handles strings, plotter.XYers and *plotter.Function, got %T", t))
		}
	}
	plt.Add(ps...)
	for _, v := range items {
		plt.Legend.Add(v.name, v.value)
	}
	return nil
}

// AddLinePoints adds Line and Scatter plotters to a
// nplot.  The variadic arguments must be either strings
// or plotter.XYers.  Each plotter.XYer is added to
// the nplot using the next color, dashes, and glyph
// shape via the Color, Dashes, and Shape functions.
// If a plotter.XYer is immediately preceeded by
// a string then a legend entry is added to the nplot
// using the string as the name.
//
// If an error occurs then none of the plotters are added
// to the nplot, and the error is returned.
func AddLinePoints(plt *nplot.Plot, vs ...interface{}) error {
	var ps []nplot.Plotter
	type item struct {
		name  string
		value [2]nplot.Thumbnailer
	}
	var items []item
	name := ""
	var i int
	for _, v := range vs {
		switch t := v.(type) {
		case string:
			name = t

		case plotter.XYer:
			l, s, err := plotter.NewLinePoints(t)
			if err != nil {
				return err
			}
			l.Color = Color(i)
			l.Dashes = Dashes(i)
			s.Color = Color(i)
			s.Shape = Shape(i)
			i++
			ps = append(ps, l, s)
			if name != "" {
				items = append(items, item{name: name, value: [2]nplot.Thumbnailer{l, s}})
				name = ""
			}

		default:
			panic(fmt.Sprintf("plotutil: AddLinePoints handles strings and plotter.XYers, got %T", t))
		}
	}
	plt.Add(ps...)
	for _, item := range items {
		v := item.value[:]
		plt.Legend.Add(item.name, v[0], v[1])
	}
	return nil
}

// AddErrorBars adds XErrorBars and YErrorBars
// to a nplot.  The variadic arguments must be
// of type plotter.XYer, and must be either a
// plotter.XErrorer, plotter.YErrorer, or both.
// Each errorer is added to the nplot the color from
// the Colors function corresponding to its position
// in the argument list.
//
// If an error occurs then none of the plotters are added
// to the nplot, and the error is returned.
func AddErrorBars(plt *nplot.Plot, vs ...interface{}) error {
	var ps []nplot.Plotter
	for i, v := range vs {
		added := false

		if xerr, ok := v.(interface {
			plotter.XYer
			plotter.XErrorer
		}); ok {
			e, err := plotter.NewXErrorBars(xerr)
			if err != nil {
				return err
			}
			e.Color = Color(i)
			ps = append(ps, e)
			added = true
		}

		if yerr, ok := v.(interface {
			plotter.XYer
			plotter.YErrorer
		}); ok {
			e, err := plotter.NewYErrorBars(yerr)
			if err != nil {
				return err
			}
			e.Color = Color(i)
			ps = append(ps, e)
			added = true
		}

		if added {
			continue
		}
		panic(fmt.Sprintf("plotutil: AddErrorBars expects plotter.XErrorer or plotter.YErrorer, got %T", v))
	}
	plt.Add(ps...)
	return nil
}

// AddXErrorBars adds XErrorBars to a nplot.
// The variadic arguments must be
// of type plotter.XYer, and plotter.XErrorer.
// Each errorer is added to the nplot the color from
// the Colors function corresponding to its position
// in the argument list.
//
// If an error occurs then none of the plotters are added
// to the nplot, and the error is returned.
func AddXErrorBars(plt *nplot.Plot, es ...interface {
	plotter.XYer
	plotter.XErrorer
}) error {
	var ps []nplot.Plotter
	for i, e := range es {
		bars, err := plotter.NewXErrorBars(e)
		if err != nil {
			return err
		}
		bars.Color = Color(i)
		ps = append(ps, bars)
	}
	plt.Add(ps...)
	return nil
}

// AddYErrorBars adds YErrorBars to a nplot.
// The variadic arguments must be
// of type plotter.XYer, and plotter.YErrorer.
// Each errorer is added to the nplot the color from
// the Colors function corresponding to its position
// in the argument list.
//
// If an error occurs then none of the plotters are added
// to the nplot, and the error is returned.
func AddYErrorBars(plt *nplot.Plot, es ...interface {
	plotter.XYer
	plotter.YErrorer
}) error {
	var ps []nplot.Plotter
	for i, e := range es {
		bars, err := plotter.NewYErrorBars(e)
		if err != nil {
			return err
		}
		bars.Color = Color(i)
		ps = append(ps, bars)
	}
	plt.Add(ps...)
	return nil
}

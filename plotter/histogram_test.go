// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter_test

import (
	"testing"
	"time"

	"github.com/hneemann/nplot"
	"github.com/hneemann/nplot/cmpimg"
	"github.com/hneemann/nplot/plotter"
	"github.com/hneemann/nplot/vg"
)

func TestHistogram(t *testing.T) {
	cmpimg.CheckPlot(ExampleHistogram, t, "histogram.png")
}

func TestSingletonHistogram(t *testing.T) {
	done := make(chan struct{}, 1)
	go func() {
		defer close(done)
		p, err := nplot.New()
		if err != nil {
			t.Fatalf("unexpected error from nplot.New: %v", err)
		}

		hist, err := plotter.NewHist(plotter.Values([]float64{1.0}), 60)
		if err != nil {
			t.Fatalf("unexpected error from NewHist: %v", err)
		}
		hist.Normalize(1)

		p.Add(hist)

		_, err = p.WriterTo(4*vg.Inch, 4*vg.Inch, "png")
		if err != nil {
			t.Fatalf("unexpected error from WriterTo: %v", err)
		}
	}()

	select {
	case <-time.After(10 * time.Second):
		t.Error("histogram timed out")
	case <-done:
	}
}
func TestHistogramLogScale(t *testing.T) {
	cmpimg.CheckPlot(ExampleHistogram_logScaleY, t, "histogram_logy.png")
}

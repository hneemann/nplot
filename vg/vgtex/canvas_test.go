// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vgtex_test

import (
	"testing"

	"github.com/hneemann/plot/cmpimg"
)

func TestTexCanvas(t *testing.T) {
	cmpimg.CheckPlot(Example, t, "scatter.tex")
}

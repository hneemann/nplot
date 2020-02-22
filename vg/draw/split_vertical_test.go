// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw_test

import (
	"github.com/hneemann/plot/vg"
	"github.com/hneemann/plot/vg/draw"
)

// SplitVertical returns the lower and upper portions of c after
// splitting it along a horizontal line distance y from the
// bottom of c.
func SplitVertical(c draw.Canvas, y vg.Length) (lower, upper draw.Canvas) {
	return draw.Crop(c, 0, 0, 0, c.Min.Y-c.Max.Y+y), draw.Crop(c, 0, 0, y, 0)
}

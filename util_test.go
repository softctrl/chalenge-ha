//
// Author:
//  Carlos Timoshenko
//  carlostimoshenkorodrigueslopes@gmail.com
//
//  https://github.com/softctrl
//
// This project is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 3 of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
//
package main

import (
	m "math"
	"testing"
)

const DEGREE = 123456789

// Basic aritmethics:
func TestDegreeToRadian(t *testing.T) {

	if r := DegreeToRadian(DEGREE * 180 / m.Pi); r != DEGREE {
		t.Error("Value received: %f", r)
	}

}

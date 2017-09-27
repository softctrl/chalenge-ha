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

const DEGREE_180 = 180

var RECORD = []string{"1", "180", "180"}

var HA_LOCATION = NewLocation(0, 51.925146, 4.478617)
var PARIS_LOCATION = NewLocation(1, 48.8589506, 2.2768484)
var DUBLIN_LOCATION = NewLocation(2, 53.3244427, -6.3861302)

var DISTANCE = 374.98929098677536

func TestNewLocation(t *testing.T) {

	if loc := NewLocation(1, DEGREE_180, DEGREE_180); loc.lat != m.Pi && loc.lng != m.Pi {
		t.Error("Wrong location created : ", loc.ToString())
	}

}

func TestNewLocationFromRecord(t *testing.T) {

	if loc := NewLocationFromRecord(RECORD); loc.lat != m.Pi && loc.lng != m.Pi {
		t.Error("Wrong location created : ", loc.ToString())
	}

}

func TestDistFrom(t *testing.T) {

	if dist := PARIS_LOCATION.DistFrom(HA_LOCATION); dist != DISTANCE {
		t.Error("Wrong distance : ", dist)
	}

}

func TestCompareTo(t *testing.T) {

	PARIS_LOCATION.DistFrom(HA_LOCATION)
	DUBLIN_LOCATION.DistFrom(HA_LOCATION)
	if comp := PARIS_LOCATION.CompareTo(DUBLIN_LOCATION); comp != -1 {
		t.Error("Wrong comparation : ", comp)
	}

}

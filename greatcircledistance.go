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
)

// The earth's radius in Km.
const EARTHS_RADIUS float64 = 6372 // Km / 6371e3 m

// Based on https://en.wikipedia.org/wiki/Great-circle_distance
// So, this aproach give us a 0.5% if error, i recomend that in a production
// code we use harversine formula to achieve this information with accuraci,
// or even better, let the PostGis do this.
// DistanceInKM calculate the distance between two locations in Km.
func DistanceInKM(pos1, pos2 *Location) float64 {

	return EARTHS_RADIUS * centralAngle(pos1, pos2)

}

// Based on https://en.wikipedia.org/wiki/Great-circle_distance
// centralAngle calculate the central angle between two positions using the
// spherical law of cosines.
func centralAngle(loc1, loc2 *Location) float64 {

	return m.Acos(m.Sin(loc1.lat)*m.Sin(loc2.lat) +
		m.Cos(loc1.lat)*m.Cos(loc2.lat)*m.Cos(loc2.lng-loc1.lng))

}

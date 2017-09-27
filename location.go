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
	"fmt"
	sc "strconv"
)

// It will keep the location data.
type Location struct {
	id   int     // Identifier
	lat  float64 // Latitude (radians)
	lng  float64 // Longitude (radians)
	dist float64 // Distance to the reference.
}

// Just to create a new pos struct. Also convert Degree values to Radian values.
func NewLocation(id int, lat, long float64) *Location {

	return &Location{id, DegreeToRadian(lat), DegreeToRadian(long), -1}

}

// NewLocationFromRecord just creates a valid Location from a informed record.
func NewLocationFromRecord(record []string) *Location {

	id := 0

	if len(record) != 3 { // Sorry but i need all values here.
		return nil // Let it crash..
	}

	if _id, e := sc.Atoi(record[0]); e == nil { // Get ID:
		id = _id
	}

	var lat float64 = 0
	if _lat, e := sc.ParseFloat(record[1], 64); e == nil { // Get LATITUDE:
		lat = _lat
	}

	var lng float64 = 0
	if _lng, e := sc.ParseFloat(record[2], 64); e == nil { // Get LONGITUDE:
		lng = _lng
	}

	return NewLocation(id, lat, lng)

}

// Calculates the distance from [fromLoc] to the [__obj] instance.
func (__obj *Location) DistFrom(fromLoc *Location) float64 {
	__obj.dist = DistanceInKM(__obj, fromLoc)
	return __obj.dist
}

// CompareTo validate if the [__obj] is greater than, equal to or less than
// [other] value.
func (__obj *Location) CompareTo(other *Location) int {

	var result = 0               // Equal to.
	if __obj.dist < other.dist { // Less than.
		result = -1
	} else if __obj.dist > other.dist { // Greater than.
		result = 1
	}
	return result

}

// ToString generate a huma readable version for the [__obj] object.
func (__obj *Location) ToString() string {
	return fmt.Sprintf("Location(id=%d, lat=%f, lng=%f, dist=%f)", __obj.id,
		__obj.lat, __obj.lng, __obj.dist)
}

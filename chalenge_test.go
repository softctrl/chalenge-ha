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
	"testing"
)

func TestNewList_C_Front_F_Front(t *testing.T) {

	if list := NewList(DUBLIN_LOCATION, 5); list.C_Front() != list.F_Front() {
		t.Error("Wrong values for closest and further lists: ", list.C_Front().ToString(), list.F_Front().ToString())
	}

}

func Test_C_OrdInsert_F_OrdInsert_Insert(t *testing.T) {

	list := NewList(DUBLIN_LOCATION, 5)
	list.Insert(PARIS_LOCATION)
	locF := list.closestList.Front().Value.(*Location)
	locB := list.furthestList.Back().Value.(*Location)
	if locF != locB {
		t.Error("Wrong values: ", locF, locB)
	}

}

func TestChalenge(t *testing.T) {

	MAX := 5

	ds, err := NewCSVDataSource("./geoData.csv")
	if err != nil {
		t.Error("Error when create a NewCSVDataSource: ", err.Error())
	}

	var list *List = nil
	if loc, err := ds.Next(); err == nil {

		loc.DistFrom(HA_LOCATION)
		list = NewList(loc, MAX)
		for {

			if loc, err := ds.Next(); err == nil {

				loc.DistFrom(HA_LOCATION)
				list.Insert(loc)

			} else {
				break

			}

		}

	} else {
		t.Error("Error reading the csv file: ", err.Error())
	}

	if list.Closest().Len() != MAX || list.Closest().Len() != MAX {
		t.Error("Wrong number of elements: ", list.Closest().Len(), list.Closest().Len())
	}

	// Closest:
	if list.Closest().Front().Value.(*Location).id != 442406 {
		t.Error("When expected 442406 it got: ", list.Closest().Front().Value.(*Location).id)
	}

	e := list.Closest().Front().Next()
	if e.Value.(*Location).id != 285782 {
		t.Error("When expected 285782 it got: ", e.Value.(*Location).id)
	}

	e = e.Next()
	if e.Value.(*Location).id != 429151 {
		t.Error("When expected 429151 it got: ", e.Value.(*Location).id)
	}

	e = e.Next()
	if e.Value.(*Location).id != 512818 {
		t.Error("When expected 512818 it got: ", e.Value.(*Location).id)
	}

	e = e.Next()
	if e.Value.(*Location).id != 25182 {
		t.Error("When expected 25182 it got: ", e.Value.(*Location).id)
	}

}

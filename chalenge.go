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
	"container/list"
	"fmt"
)

// List will keep two lists one for the closests places and other for the
// furthest ones.
type List struct {
	closestList  *list.List
	furthestList *list.List
	nth          int
}

// NewList creates a new List with a initial value
func NewList(loc *Location, nth int) *List {

	lst := &List{closestList: list.New(), furthestList: list.New(), nth: nth}
	lst.closestList.PushFront(loc)
	lst.furthestList.PushFront(loc)
	return lst

}

// Closest return closest list.
func (__obj *List) Closest() *list.List {

	return __obj.closestList

}

// Furtherst return closest list.
func (__obj *List) Furtherst() *list.List {

	return __obj.furthestList

}

// C_Front return the Front value of the closest list.
func (__obj *List) C_Front() *Location {

	return __obj.closestList.Front().Value.(*Location)

}

// C_Front return the Front value of the furthest list.
func (__obj *List) F_Front() *Location {

	return __obj.furthestList.Front().Value.(*Location)

}

// _C_OrdInsert make an ordered insert into the closest list of locations.
func (__obj *List) _C_OrdInsert(loc *Location) {

	for e := __obj.closestList.Front(); e != nil; e = e.Next() {

		if loc.CompareTo((e.Value.(*Location))) < 0 {
			__obj.closestList.InsertBefore(loc, e)
			break
		}

	}

}

// _F_OrdInsert make an ordered insert into the furthest list of locations.
func (__obj *List) _F_OrdInsert(loc *Location) {

	for e := __obj.furthestList.Back(); e != nil; e = e.Prev() {

		if loc.CompareTo((e.Value.(*Location))) < 0 {
			__obj.furthestList.InsertAfter(loc, e)
			break
		}

	}

}

// _CRemoveLast remove lest element.
func (__obj *List) _CRemoveLast() {

	__obj.closestList.Remove(__obj.closestList.Back()) // Remove the last one.

}

// _FRemoveLast remove lest element.
func (__obj *List) _FRemoveLast() {

	__obj.furthestList.Remove(__obj.furthestList.Back()) // Remove the last one.

}

// Here the magics happens :)
// Insert insert a new value into a priority queue. Also at the same time i
// am limiting the maximun size to the 5th elements, and also i am doing an
// online sort, so i do not need to get all data from memory and y only need
// to go through the data only once, and the sort will only take 5 elements
// at time
func (__obj *List) Insert(loc *Location) *List {

	// Closest elements (OK)
	if loc.CompareTo(__obj.C_Front()) < 0 { // If the new value is less than front:

		__obj.closestList.PushFront(loc) // Put new value on the front.

	} else {

		__obj._C_OrdInsert(loc) // Try to make an ordered insert.

	}

	// Furthest elements (OK)
	if loc.CompareTo(__obj.F_Front()) > 0 { // If the new value is bigger than front:

		__obj.furthestList.PushFront(loc) // Put new value on the front.

	} else {

		__obj._F_OrdInsert(loc) // Try to make an ordered insert.

	}

	if __obj.closestList.Len() > __obj.nth { // If you have more then nth elements:

		__obj._CRemoveLast() // Remove the last one.

	}

	if __obj.furthestList.Len() > __obj.nth { // If you have more then nth elements:

		__obj._FRemoveLast() // Remove the last one.

	}

	return __obj

}

// Print creates a human readable version with the closest list and the
// further list.
func (__obj *List) Print() {

	var count = 0
	fmt.Println("----- Closest ones")
	for e := __obj.closestList.Front(); e != nil; e = e.Next() {
		count = count + 1
		fmt.Printf("[%d] <%d>- %f\n", count, (e.Value.(*Location)).id, (e.Value.(*Location)).dist /*ToString()*/)
	}

	count = 0
	fmt.Println("\n+++++ Furthest ones")
	for e := __obj.furthestList.Front(); e != nil; e = e.Next() {
		count = count + 1
		fmt.Printf("[%d] <%d>- %f\n", count, (e.Value.(*Location)).id, (e.Value.(*Location)).dist /*ToString()*/)
	}

}

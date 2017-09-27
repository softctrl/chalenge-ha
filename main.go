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
)

// Parameters that users can inform to change my behaviour.
var config *Configuration

// it reminds me of the static block on Java :D
func init() {

	config = ParseConfiguration()

}

// Start point for the Program.
func main() {

	// Create a location for the reference, it maybe the HousingAnywhere location.
	haLocation := NewLocation(0, *config.lat, *config.lng)

	// So, here we will identify your data source.
	var ds DataSource = nil

	if *config.dSource == 0 { // If CSV file:

		ds, _ = NewCSVDataSource(*config.csvFile)

	} else if *config.dSource == 1 { // If PostgreSQL connection:

		fmt.Println("### PostgreSQL support under development.")
		return

	}

	// Start to read all Data:
	if loc, err := ds.Next(); err == nil {

		loc.DistFrom(haLocation)
		var list *List = NewList(loc, *config.nth)

		for {

			if loc, err := ds.Next(); err == nil {

				loc.DistFrom(haLocation)
				list.Insert(loc)

			} else {
				break
			}

		}

		// Show us your --power--, sorry, show me the list:
		list.Print()
	}

}

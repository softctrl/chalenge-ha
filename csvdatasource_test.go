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
	"io"
	"testing"
)

func TestCSVDataSource(t *testing.T) {

	csv, err := NewCSVDataSource("./test_csv_data_source.csv")
	if err != nil {
		t.Error("Error on NewCSVDataSource : ", err.Error())
	}

	if csv.reader.FieldsPerRecord != 2 {
		t.Error("Wrong number of FieldsPerRecord: ", csv.reader.FieldsPerRecord)
	}
	count := 0
	for {
		if _, err := csv.Next(); err == nil {
			count = count + 1
		} else {

			if err == io.EOF {
				break
			} else {
				t.Error("Error when reading the file: ", err.Error())
			}

		}
	}
	if count != 10 {
		t.Error("The file return more then 10 lines. Resuled: ", count)
	}

}

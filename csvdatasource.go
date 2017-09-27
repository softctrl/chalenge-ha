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
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// _CSVDataSource will implements the Datasource interface:
type _CSVDataSource struct {
	file   string     // The csv file address
	reader csv.Reader // A csv Reader
}

// NewCSVDataSource creates a new _CSVDataSource with the informed csv file.
func NewCSVDataSource(file string) (*_CSVDataSource, error) {

	csvDs := _CSVDataSource{}
	csvDs.file = file

	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	csvDs.reader = *csv.NewReader(csvFile)
	_, err = csvDs.reader.Read()
	if err != nil && err == io.EOF {
		fmt.Printf("Empty File?!?!?!\n%s\n", err.Error())
	}

	return &csvDs, nil

}

// All new data source need to implement this method.
// Next move to the next register if there exists.
func (__obj _CSVDataSource) Next() (*Location, error) {

	rec, err := __obj.reader.Read()
	if err != nil {
		return nil, err
	} else {
		return NewLocationFromRecord(rec), nil
	}

}

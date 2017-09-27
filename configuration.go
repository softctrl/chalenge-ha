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
	f "flag"
	"fmt"
)

// All Parameters
const (
	DATA_SOURCE_PARAM = "ds"
	CSV_FILE_PARAM    = "csv"
	LATITUDE_PARAM    = "lt"
	LONGITUDE_PARAM   = "lg"
	NTH_PARAM         = "n"
)

// All default Parameter values
const (
	DEFAULT_DATA_SOURCE = 0 // 0 - CSV / 1 - PostgreSQL
	DEFAULT_HELP_CSV    = "./geoData.csv"
	DEFAULT_LATITUDE    = 51.925146
	DEFAULT_LONGITUDE   = 4.478617
	DEFAULT_NTH         = 5 // Default
)

// All Parameter helpers
const (
	HELP_DATA_SOURCE = "Configure the source of the data. you can inform 0 for csv or 1 for PostgreSQL (under development)"
	HELP_CSV_FILE    = "Inform a valid csv file to be processed."
	HELP_LATITUDE    = "Inform the reference latitude to measure the distance."
	HELP_LONGITUDE   = "Inform the reference longitude to measure the distance."
	HELP_NTH         = "Inform the max relevant elements you want."
)

// This struct will keep all configuration informed by the user.
type Configuration struct {
	dSource *int     // 0-CSV / 1-PostgreSQL(under development).
	csvFile *string  // If dSource=0, need to informe where is the csn file.
	lat     *float64 // The reference latitude to measure distance.
	lng     *float64 // The reference longitude to measure distance.
	nth     *int     // Max elements.
}

// ToString creates a human readable version of the [__obj] instance.
func (__obj *Configuration) ToString() string {
	return fmt.Sprintf("Configuration[dSource= %d, lat= %f, lng = %f, nth = %d]",
		__obj.dSource, __obj.lat, __obj.lng, __obj.nth)
}

// ParseConfiguration parses a valid Configuration from user inputs.
func ParseConfiguration() *Configuration {

	res := &Configuration{f.Int(DATA_SOURCE_PARAM, DEFAULT_DATA_SOURCE, HELP_DATA_SOURCE),
		f.String(CSV_FILE_PARAM, DEFAULT_HELP_CSV, HELP_CSV_FILE),
		f.Float64(LATITUDE_PARAM, DEFAULT_LATITUDE, HELP_LATITUDE),
		f.Float64(LONGITUDE_PARAM, DEFAULT_LONGITUDE, HELP_LONGITUDE),
		f.Int(NTH_PARAM, DEFAULT_NTH, HELP_NTH)}
	f.Parse()
	return res

}

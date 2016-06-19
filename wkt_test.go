/*
Copyright 2016, RadiantBlue Technologies, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/venicegeo/geojson-go/geojson"
)

var inputWKTFiles = [...]string{
	"test/point.wkt",
	"test/linestring.wkt",
	"test/polygon.wkt",
	"test/polygon2.wkt",
	"test/multipoint.wkt",
	"test/multipoint2.wkt"}

var outputGeojsons = [...]string{
	"{\"type\":\"Point\",\"coordinates\":[30,10]}",
	"{\"type\":\"LineString\",\"coordinates\":[[30,10],[10,30],[40,40]]}",
	"{\"type\":\"Polygon\",\"coordinates\":[[[30,10],[40,40],[20,40],[10,20],[30,10]]]}",
	"{\"type\":\"Polygon\",\"coordinates\":[[[35,10],[45,45],[15,40],[10,20],[35,10]],[[20,30],[35,35],[30,20],[20,30]]]}",
	"{\"type\":\"MultiPoint\",\"coordinates\":[[10,40],[40,30],[20,20],[30,10]]}",
	"{\"type\":\"MultiPoint\",\"coordinates\":[[10,40],[40,30],[20,20],[30,10]]}"}

// TestWKT tests WKT parsing
func TestWKT(t *testing.T) {
	var (
		gj            interface{}
		err           error
		bytes         []byte
		input, output string
	)
	for inx, fileName := range inputWKTFiles {
		if bytes, err = ioutil.ReadFile(fileName); err != nil {
			log.Panicf("Parse error: %v\n", err)
		}
		input = string(bytes)
		gj = geojson.WKT(input)

		if bytes, err = geojson.Write(gj); err != nil {
			log.Panicf("Write error: %v\n", err)
		}
		output = string(bytes)
		if output == outputGeojsons[inx] {
			fmt.Printf("Parsed: %v\nOutput: %v\n", input, output)
		} else {
			t.Errorf("Expected: %v\nFound: %v\n", outputGeojsons[inx], output)
		}
	}
}

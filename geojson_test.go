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
	"testing"

	"github.com/venicegeo/geojson-go/geojson"
)

var inputFiles = [...]string{
	"test/point.geojson",
	"test/linestring.geojson",
	"test/polygon.geojson",
	"test/polygon-dateline.geojson",
	"test/polygon-hole.geojson",
	"test/multipoint.geojson",
	"test/multilinestring.geojson",
	"test/multipolygon.geojson",
	"test/geometrycollection.geojson",
	"test/sample.geojson",
	"test/boundingbox.geojson"}

// TestGeoJSON tests GeoJSON readers
func TestGeoJSON(t *testing.T) {
	for _, fileName := range inputFiles {
		process(fileName)
	}
}

func TestToGeometryArray(t *testing.T) {
	var (
		gj     interface{}
		err    error
		result []interface{}
	)
	if gj, err = geojson.ParseFile("test/sample.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	result = geojson.ToGeometryArray(gj)
	fmt.Printf("Geometries: %#v\n", result)
}

func TestBbox(t *testing.T) {
	for inx, fileName := range inputFiles {
		gj, _ := geojson.ParseFile(fileName)
		bboxIfc := gj.(geojson.BoundingBoxIfc)
		bbox := bboxIfc.ForceBbox()
		fmt.Printf("%v BBox: %v\n", inx+1, bbox.String())
	}
}

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

package geojson

import (
	"fmt"
	"log"
	"testing"
)

var inputFiles = [...]string{
	"test/point.geojson",
	"test/point2.geojson",
	"test/point3.geojson",
	"test/linestring.geojson",
	"test/polygon.geojson",
	"test/polygon-dateline.geojson",
	"test/polygon-hole.geojson",
	"test/multipoint.geojson",
	"test/multilinestring.geojson",
	"test/multipolygon.geojson",
	"test/geometrycollection.geojson",
	"test/sample.geojson",
	"test/boundingbox.geojson",
	"test/feature.geojson",
	"test/featureCollectionWithGeometryCollection.geojson",
	"test/featureCollection.geojson"}

func testProcess(filename string) {
	var (
		gj    interface{}
		err   error
		bytes []byte
	)
	if gj, err = ParseFile(filename); err != nil {
		log.Panicf("Parse error: %v\n", err)
	}
	fmt.Printf("%T: %#v\n", gj, gj)

	if bytes, err = Write(gj); err != nil {
		log.Panicf("Write error: %v\n", err)
	}
	fmt.Printf("%v\n", string(bytes))
}

// TestGeoJSON tests GeoJSON readers
func TestGeoJSON(t *testing.T) {
	for _, fileName := range inputFiles {
		testProcess(fileName)
	}
}

// Testing multilinestring geometrycollection multipoint multipolygon
func TestMultiAndColection(t *testing.T) {
	var gj interface{}
	var err error
	if gj, err = ParseFile("test/multilinestring.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	multiLineString := gj.(*MultiLineString)
	_ = multiLineString.String()

	if gj, err = ParseFile("test/geometrycollection.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	geometryCollection := gj.(*GeometryCollection)
	_ = geometryCollection.String()

	if gj, err = ParseFile("test/multipolygon.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	multiPolygon := gj.(*MultiPolygon)
	_ = multiPolygon.String()
	if gj, err = ParseFile("test/multipoint.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	multiPoint := gj.(*MultiPoint)
	_ = multiPoint.String()

	if gj, err = ParseFile("test/multipolygon2.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	multiPolygon2 := gj.(*MultiPolygon)
	_ = multiPolygon2.String()

}

// Testing FromMap
func TestFromMap(t *testing.T) {
	var gj interface{}
	var err error
	if gj, err = ParseFile("test/featureCollection.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	featureCollection := gj.(*FeatureCollection)
	map1 := featureCollection.Map()
	if FromMap(map1) == nil {
		t.Errorf("Failed to parse file: %v", err)
		t.Log(FromMap(map1))
	}

	if gj, err = ParseFile("test/feature.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	feature := gj.(*Feature)
	map2 := make(map[string]interface{})
	map2["type"] = feature.Type
	map2["properties"] = feature.Properties
	map2["geometry"] = feature.Geometry
	if FromMap(map2) == nil {
		t.Errorf("Failed to parse file: %v", err)
		t.Log(FromMap(map2))
	}

	if err = testWriteFile(map2); err != nil {
		t.Errorf("Failed to write file: %v", err)
	}

}

// Testing Writing
func testWriteFile(input map[string]interface{}) error {
	var err error
	err = WriteFile(input, "test/TestWriteFile.geojson")
	return err
}

func TestNullInputs(t *testing.T) {
	bb, _ := NewBoundingBox(nil)
	if "" != bb.String() {
		fmt.Print(bb.String())
		t.Error("Couldn't handle nil bounding box")
	}
	point := bb.Centroid()
	if point != nil {
		t.Error("Expected a nil Centroid for an empty bounding box")
	}
	fc := NewFeatureCollection(nil)
	if fc.String() != `{"type":"FeatureCollection","features":[]}` {
		t.Errorf("Received %v for empty Feature Collection.", fc.String())
	}
	f := NewFeature(nil, nil, nil)
	if f.String() != `{"type":"Feature","geometry":null}` {
		t.Errorf("Received %v for an empty Feature.", f.String())
	}
	fc.Features = append(fc.Features, f)
	if fc.String() != `{"type":"FeatureCollection","features":[{"type":"Feature","geometry":null}]}` {
		t.Errorf("Received %v for a feature collection with a single empty feature", fc.String())
	}
}

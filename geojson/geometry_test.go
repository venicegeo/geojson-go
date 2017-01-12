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
	"testing"
)

func TestToGeometryArray(t *testing.T) {
	var (
		gj     interface{}
		err    error
		result []interface{}
	)
	if gj, err = ParseFile("test/sample.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	result = ToGeometryArray(gj)
	fmt.Printf("Geometries: %#v\n", result)
}

func TestRTPoint(t *testing.T) {
	var (
		gj     interface{}
		err    error
		m      map[string]interface{}
		p1     *Point
		p2     *Point
		result = `{"type":"Point","coordinates":[100,0]}`
	)
	if gj, err = ParseFile("test/point.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	p1 = gj.(*Point)
	m = p1.Map()
	p2 = FromMap(m).(*Point)
	if p2.String() != result {
		t.Errorf("Round trip point failed: %v", p2.String())
	}
}

func TestToMultiPoint(t *testing.T) {
	var (
		gj     interface{}
		err    error
		result *MultiPoint
	)
	if gj, err = ParseFile("test/sample.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	result = ToMultiPoint(gj)
	if len(result.Coordinates) != 10 {
		t.Errorf("Found %v points, expected 10.\n%#v\n", len(result.Coordinates), result)
	}
}

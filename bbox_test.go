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
	"testing"

	"github.com/venicegeo/geojson-go/geojson"
)

const bbox0 = ""
const bbox1 = "10,10,20,20"
const bbox2 = "10,10,20,20,30"
const bbox3 = "10,10,20,foo"
const bbox4 = "10,10,20,20,30,30"
const bbox5 = "40,10,20,20,30,30"

// TestGeoJSON tests GeoJSON readers
func TestBBox(t *testing.T) {
	var (
		err  error
		bbox geojson.BoundingBox
	)
	if bbox, err = geojson.NewBoundingBox(bbox0); err != nil {
		t.Error(err)
	}
	if bbox.Valid() != nil {
		t.Errorf("\"%v\" is supposed to be a valid bounding box but it returned %v.", bbox0, bbox.Valid().Error())
	}
	if bbox, err = geojson.NewBoundingBox(bbox1); err != nil {
		t.Error(err)
	}
	if bbox.Valid() != nil {
		t.Errorf("\"%v\" is supposed to be a valid bounding box.", bbox1)
	}
	if _, err = geojson.NewBoundingBox(bbox2); err == nil {
		t.Errorf("\"%v\" is supposed to be an invalid bounding box.", bbox2)
	}
	if _, err = geojson.NewBoundingBox(bbox3); err == nil {
		t.Errorf("\"%v\" is supposed to be an invalid bounding box.", bbox3)
	}
	if bbox, err = geojson.NewBoundingBox(bbox4); err != nil {
		t.Error(err)
	}
	if bbox.Valid() != nil {
		t.Errorf("\"%v\" is supposed to be a valid bounding box but it returned %v.", bbox4, bbox.Valid().Error())
	}
	if _, err = geojson.NewBoundingBox(bbox5); err == nil {
		t.Errorf("\"%v\" is supposed to be an invalid bounding box.", bbox5)
	}
}
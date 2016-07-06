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
	"log"
	"testing"

	"github.com/venicegeo/geojson-go/geojson"
)

const bbox0 = ""
const bbox1 = "10,10,20,20"
const bbox2 = "10,10,20,20,30"
const bbox3 = "10,10,20,foo"
const bbox4 = "10,10,20,20,30,30"
const bbox5 = "40,10,20,20,30,30"
const bbox6 = "10,40,20,20,30,30"
const bbox7 = "-180,10,-170,20"
const bbox8 = "170,10,180,20"
const bbox9 = "-180,-70,-179,70"

// TestGeoJSON tests GeoJSON readers
func TestBBox(t *testing.T) {
	var (
		err             error
		bbox, otherBbox geojson.BoundingBox
		gjIfc           interface{}
		fc              *geojson.FeatureCollection
		mpoly           *geojson.MultiPolygon
		ok              bool
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
	if bbox, err = geojson.NewBoundingBox(bbox5); err != nil {
		t.Error(err)
	}
	if bbox.Valid() != nil {
		t.Errorf("\"%v\" is supposed to be a valid bounding box but it returned %v.", bbox5, bbox.Valid().Error())
	}
	if !bbox.Antimeridian() {
		t.Errorf("\"%v\" crosses the antimeridian but the check returned false.", bbox5)
	}
	if _, err = geojson.NewBoundingBox(bbox6); err == nil {
		t.Errorf("\"%v\" is supposed to be an invalid bounding box.", bbox6)
	}
	if bbox, err = geojson.NewBoundingBox(bbox7); err != nil {
		t.Error(err)
	}
	if bbox.Valid() != nil {
		t.Errorf("\"%v\" is supposed to be a valid bounding box but it returned %v.", bbox7, bbox.Valid().Error())
	}
	if otherBbox, err = geojson.NewBoundingBox(bbox8); err != nil {
		t.Error(err)
	}
	if bbox.Valid() != nil {
		t.Errorf("\"%v\" is supposed to be a valid bounding box but it returned %v.", bbox8, bbox.Valid().Error())
	}
	if bbox, err = geojson.NewBoundingBox([]geojson.BoundingBox{bbox, otherBbox}); err == nil {
		if !bbox.Antimeridian() {
			t.Errorf("Joining \"%v\" and \"%v\" created \"%v\" which should cross the antimeridian but the check returned false.", bbox7, bbox8, bbox.String())
		}
	} else {
		t.Errorf("Joining \"%v\" and \"%v\" is supposed to be a valid bounding box.", bbox7, bbox8)
	}
	if otherBbox, err = geojson.NewBoundingBox(bbox9); err == nil {
		if !bbox.Overlaps(otherBbox) {
			t.Error("These bounding boxes (7+8, 9) should overlap.")
			log.Printf("9: %v; 7+8: %v", otherBbox.String(), bbox.String())
		}
		if !otherBbox.Overlaps(bbox) {
			t.Error("These bounding boxes (9, 7+8) should overlap.")
			log.Printf("9: %v; 7+8: %v", otherBbox.String(), bbox.String())
		}
	} else {
		t.Error(err)
	}
	if gjIfc, err = geojson.ParseFile("test/featureCollection.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	if fc, ok = gjIfc.(*geojson.FeatureCollection); ok {
		otherBbox = fc.ForceBbox()
		if bbox.Overlaps(otherBbox) {
			t.Error("These bounding boxes (7+8, fc) should not overlap.")
			log.Printf("7+8: %v; fc: %v", bbox.String(), otherBbox.String())
		}
	} else {
		t.Errorf("Expected *geojson.FeatureCollection, got %T", gjIfc)
	}
	if gjIfc, err = geojson.ParseFile("test/multipolygon2.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	if mpoly, ok = gjIfc.(*geojson.MultiPolygon); ok {
		bbox = mpoly.ForceBbox()
		if !bbox.Antimeridian() {
			t.Errorf("The multipolygon should cross the antimeridian. %v", bbox.String())
		}
	} else {
		t.Errorf("Expected *geojson.FeatureCollection, got %T", gjIfc)
	}
}

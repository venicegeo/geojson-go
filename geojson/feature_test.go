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
	"math"
	"testing"
)

type functor func(string)

func testFeaturePropertyString(t *testing.T, f *Feature, prop, expected string) {
	if f.PropertyString(prop) != expected {
		t.Errorf("Expected '%v', got '%v'", expected, f.PropertyString(prop))
	}
}

func testFeaturePropertyInt(t *testing.T, f *Feature, prop string, expected int) {
	if f.PropertyInt(prop) != expected {
		t.Errorf("Expected '%v', got '%v'", expected, f.PropertyInt(prop))
	}
}

func testFeaturePropertyFloat(t *testing.T, f *Feature, prop string, expected float64) {
	pf := f.PropertyFloat(prop)
	if math.IsNaN(expected) {
		if !math.IsNaN(pf) {
			t.Errorf("Expected NaN, got '%v'", pf)
		}
	} else if pf != expected {
		t.Errorf("Expected '%v', got '%v'", expected, pf)
	}
}

// TestFeature tests Feature stuff
func TestFeature(t *testing.T) {
	properties := make(Map)
	properties["foo"] = "bar"
	properties["bar"] = 123
	properties["float"] = 0.0
	properties["int"] = int64(11)
	map2 := make(map[string]string)
	map2["foo2"] = "bar2"
	properties["bas"] = map2
	f := NewFeature(nil, "12345", properties)
	testFeaturePropertyString(t, f, "foo", "bar")
	testFeaturePropertyString(t, f, "bar", "123")
	testFeaturePropertyString(t, f, "float", "0")
	testFeaturePropertyString(t, f, "int", "11")
	testFeaturePropertyInt(t, f, "foo", 0)
	testFeaturePropertyInt(t, f, "bar", 123)
	testFeaturePropertyInt(t, f, "float", 0)
	testFeaturePropertyInt(t, f, "int", 11)
	testFeaturePropertyFloat(t, f, "foo", math.NaN())
	testFeaturePropertyFloat(t, f, "bar", 123)
	testFeaturePropertyFloat(t, f, "float", 0)
	testFeaturePropertyFloat(t, f, "int", 11)
}

func TestRTFeature(t *testing.T) {
	var (
		gj     interface{}
		err    error
		m      Map
		f1     *Feature
		f2     *Feature
		result = `{"type":"Feature","geometry":{"type":"LineString","coordinates":[[102,0],[103,1],[104,0],[105,1]]},"properties":{"prop0":"value0","prop1":0},"id":98765}`
	)
	if gj, err = ParseFile("test/feature.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	f1 = gj.(*Feature)
	m = f1.Map()
	f2 = FromMap(m).(*Feature)
	if f2.String() != result {
		t.Errorf("Round trip feature failed: %v", f2.String())
	}
}

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
	"math"
	"strings"
	"testing"
)

type functor func(string)
type Nato int

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
	var strAry [2]string
	strAry[0] = "big"
	strAry[1] = "bad"
	properties["sling"] = strAry
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
	f.PropertyStringSlice("sling")
	f.PropertyStringSlice("bas")
	/*Test FillProperties Map and FeatureCollectionFromMap*/
	var gj interface{}
	var err error
	if gj, err = ParseFile("test/featureCollection.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	featureCollection := gj.(*FeatureCollection)
	featureCollection.FillProperties()
	fcMap := featureCollection.Map()
	featureCollection = FeatureCollectionFromMap(fcMap)
	/*Test FeatureFromMap*/
	if gj, err = ParseFile("test/feature.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	Feat1 := gj.(*Feature)
	ToGeometryArray(Feat1)
	map3 := make(map[string]interface{})
	map3["type"] = Feat1.Type
	map3["properties"] = Feat1.Properties
	map3["geometry"] = Feat1.Geometry
	if strings.Compare(Feat1.IDStr(), "98765") != 0 {
		t.Errorf("Feature was read with a 98765 id and a %v was returned", Feat1.IDStr())
	}
	map3["id"] = 1234
	Feat2 := FeatureFromMap(map3)
	map3["id"] = "1234"
	Feat2 = FeatureFromMap(map3)
	map3["id"] = nil
	Feat2 = FeatureFromMap(map3)
	if strings.Compare(Feat2.IDStr(), "") != 0 {
		t.Errorf("Feature was given an id of nil and %v was returned", Feat2.IDStr())
	}
	if Feat2 == nil {
		t.Errorf("feature.geojson failed FeatureFromMap Test")
		t.Log(Feat1)
		t.Log(Feat2)
	}
	//Test Floatify Stringify and intify
	stringer1 := Nato(1)
	float321 := float32(5.1)
	float641 := float64(6.1)
	int1 := 1
	int641 := int64(1)
	string1 := "Test"
	interfaceLoop := make([]interface{}, 6)
	interfaceLoop[0] = stringer1
	interfaceLoop[1] = float321
	interfaceLoop[2] = float641
	interfaceLoop[3] = int1
	interfaceLoop[4] = int641
	interfaceLoop[5] = string1

	for index, element := range interfaceLoop {
		t.Log(stringify(element))
		t.Log(intify(element))
		t.Log(floatify(element))
		t.Log(index)
	}

}
func (p Nato) String() string {
	switch p {
	case 1:
		return "A"
	case 2:
		return "B"
	case 3:
		return "C"
	case 4:
		return "D"
	}
	return fmt.Sprintf("Pill(%d)", p)
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

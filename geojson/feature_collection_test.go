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

import "testing"

// TestFeatureCollection tests Feature stuff
func TestFeatureCollection(t *testing.T) {
	var (
		err error
		gj  interface{}
		gj2 interface{}
		fc  *FeatureCollection
		fcm Map
		fi  interface{}
		ok  bool
	)
	if gj, err = ParseFile("test/featureCollection.geojson"); err != nil {
		t.Error(err.Error())
	}
	if fc, ok = gj.(*FeatureCollection); ok {
		if len(fc.Features) != 4 {
			t.Errorf("Expected 4 features, got %v", len(fc.Features))
		}
		fcm = fc.Map()
		if fi, ok = fcm["features"]; ok {
			if _, ok = fi.([]interface{}); ok {

			} else {
				t.Errorf("Expected an array of interfaces. %#v", fi)
			}
		} else {
			t.Error("Expected features in map.")
		}
		gj2 = FromMap(fcm)
		if _, ok = gj2.(*FeatureCollection); !ok {
			t.Errorf("Expected *FeatureCollection, got %T", gj2)
		}
	} else {
		t.Errorf("Expected *FeatureCollection, got %T", gj)
	}
}

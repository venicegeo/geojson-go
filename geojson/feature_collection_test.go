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

// TestRTFC round trips a Feature Collection
func TestRTFC(t *testing.T) {
	var (
		gj     interface{}
		err    error
		m      map[string]interface{}
		fc1    *FeatureCollection
		fc2    *FeatureCollection
		result = `{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[102,0.5]},"properties":{"prop0":"value0"}},{"type":"Feature","geometry":{"type":"LineString","coordinates":[[102,0],[103,1],[104,0],[105,1]]},"properties":{"prop0":"value0","prop1":0}},{"type":"Feature","geometry":{"type":"Polygon","coordinates":[[[100,0],[101,0],[101,1],[100,1],[100,0]]]},"properties":{"prop0":"value0","prop1":{"this":"that"}}}]}`
	)
	if gj, err = ParseFile("test/sample.geojson"); err != nil {
		t.Errorf("Failed to parse file: %v", err)
	}
	fc1 = gj.(*FeatureCollection)
	m = fc1.Map()
	fc2 = FromMap(m).(*FeatureCollection)
	if fc2.String() != result {
		t.Errorf("Round trip feature failed: %v", fc2.String())
	}
}

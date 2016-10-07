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
	"testing"
)

// TestGeoJSON tests GeoJSON readers
func TestWFS(t *testing.T) {
	var (
		wfsURL      string
		featureType string
		fc          *FeatureCollection
		err         error
	)
	wfsURL = "http://foo.com/wfs"
	featureType = "bar"
	if fc, err = FromWFS(wfsURL, featureType); err == nil {
		t.Error("This should have returned an error.")
	}
	wfsURL = "http://gsp-geose-LoadBala-4EP8UFUE9SXL-919040015.us-east-1.elb.amazonaws.com:80/geoserver/piazza/wfs"
	featureType = "piazza:8e31e022-4e1f-4a32-b341-4eb019ab45bc"
	if fc, err = FromWFS(wfsURL, featureType); err != nil {
		t.Fatalf("Failed to load WFS layer: %v", err.Error())
	}
	if len(fc.Features) == 0 {
		t.Error("Failed to create any features")
	}
}

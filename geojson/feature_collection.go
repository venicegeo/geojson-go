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

import "encoding/json"

// FEATURECOLLECTION is a GeoJSON FeatureCollection
const FEATURECOLLECTION = "FeatureCollection"

// The FeatureCollection object represents an array of features
type FeatureCollection struct {
	Type     string      `json:"type"`
	Features []*Feature  `json:"features"`
	Bbox     BoundingBox `json:"bbox,omitempty"`
}

// FeatureCollectionFromBytes constructs a FeatureCollection from a GeoJSON byte array
// and returns its pointer
func FeatureCollectionFromBytes(bytes []byte) (*FeatureCollection, error) {
	var result FeatureCollection
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	for _, feature := range result.Features {
		feature.resolveGeometry()
	}
	return &result, nil
}

// NewFeatureCollection is the normal factory method for a FeatureCollection
func NewFeatureCollection(features []*Feature) *FeatureCollection {
	return &FeatureCollection{Type: FEATURECOLLECTION, Features: features}
}

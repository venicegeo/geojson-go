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

// FEATURE is a GeoJSON Feature
const FEATURE = "Feature"

// The Feature object represents an array of features
type Feature struct {
	Type       string                 `json:"type"`
	Geometry   interface{}            `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
	ID         string                 `json:"id,omitempty"`
	Bbox       BoundingBox            `json:"bbox,omitempty"`
}

// FeatureFromBytes constructs a Feature from a GeoJSON byte array
func FeatureFromBytes(bytes []byte) (Feature, error) {
	var result Feature
	err := json.Unmarshal(bytes, &result)
	result.resolveGeometry()
	return result, err
}

// NewFeature is the normal factory method for a feature
func NewFeature(geometry interface{}, id string, properties map[string]interface{}) *Feature {
	return &Feature{Type: FEATURE, Geometry: geometry, Properties: properties, ID: id}
}

// Since unmarshaled objects don't come back as real geometries,
// This method reconstructs them
func (feature *Feature) resolveGeometry() {
	feature.Geometry = NewGeometry(feature.Geometry.(map[string]interface{}))
}

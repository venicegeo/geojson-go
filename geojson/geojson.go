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

// FEATURECOLLECTION is a GeoJSON FeatureCollection
const FEATURECOLLECTION = "FeatureCollection"

// The GeoJSON object represents a geometry, feature, or collection of features.
type GeoJSON struct {
	Type string `json:"type"`
}

// The Feature object represents an array of features
type Feature struct {
	Type       string                 `json:"type"`
	Geometry   interface{}            `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
	ID         string                 `json:"id, omitempty"`
	// Bbox     Bbox     `json:"bbox, omitempty"`
}

// NewFeature constructs a Feature from a GeoJSON byte array
func NewFeature(bytes []byte) (Feature, error) {
	var result Feature
	err := json.Unmarshal(bytes, &result)
	result.resolveGeometry()
	return result, err
}

// Since unmarshaled objects don't come back as real geometries,
// This method reconstructs them
func (feature *Feature) resolveGeometry() {
	feature.Geometry = NewGeometry(feature.Geometry.(map[string]interface{}))
}

// The FeatureCollection object represents an array of features
type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
	//	Bbox     Bbox      `json:"bbox, omitempty"`
}

// NewFeatureCollection constructs a FeatureCollection from a GeoJSON byte array
func NewFeatureCollection(bytes []byte) (FeatureCollection, error) {
	var result FeatureCollection
	err := json.Unmarshal(bytes, &result)
	for inx := 0; inx < len(result.Features); inx++ {
		result.Features[inx].resolveGeometry()
	}
	return result, err
}

// Parse parses a GeoJSON string into a GeoJSON object
func Parse(bytes []byte) (interface{}, error) {
	var result interface{}
	var geojson GeoJSON
	err := json.Unmarshal(bytes, &geojson)
	switch geojson.Type {
	case "Point":
		result, err = NewPoint(bytes)
	case "LineString":
		result, err = NewLineString(bytes)
	case "Polygon":
		result, err = NewPolygon(bytes)
	case "MultiPoint":
		result, err = NewMultiPoint(bytes)
	case "MultiLineString":
		result, err = NewMultiLineString(bytes)
	case "MultiPolygon":
		result, err = NewMultiPolygon(bytes)
	case "GeometryCollection":
		result, err = NewGeometryCollection(bytes)
	case "Feature":
		result, err = NewFeature(bytes)
	case "FeatureCollection":
		result, err = NewFeatureCollection(bytes)
	}
	return result, err
}

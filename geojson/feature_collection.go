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

// GeoJSON FeatureCollection constants
const (
	FEATURECOLLECTION = "FeatureCollection"
	FEATURES          = "features"
)

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
		feature.ResolveGeometry()
	}
	return &result, nil
}

// ForceBbox returns a bounding box, creating one by brute force if needed
func (fc *FeatureCollection) ForceBbox() BoundingBox {
	if len(fc.Bbox) > 0 {
		return fc.Bbox
	}
	var result BoundingBox
	for _, feature := range fc.Features {
		result = mergeBboxes(result, feature.ForceBbox())
	}
	return result
}

// NewFeatureCollection is the normal factory method for a FeatureCollection
func NewFeatureCollection(features []*Feature) *FeatureCollection {
	if features == nil {
		features = make([]*Feature, 0)
	}
	return &FeatureCollection{Type: FEATURECOLLECTION, Features: features}
}

// String returns the string representation
func (fc *FeatureCollection) String() string {
	var result string
	if bytes, err := json.Marshal(fc); err == nil {
		result = string(bytes)
	} else {
		result = err.Error()
	}
	return result
}

// Map returns a map of the FeatureCollection's members
// This may be useful in wrapping a Feature Collection with foreign members
func (fc *FeatureCollection) Map() map[string]interface{} {
	result := make(map[string]interface{})
	result["type"] = fc.Type
	features := make([]interface{}, len(fc.Features))
	for inx, feature := range fc.Features {
		features[inx] = feature.Map()
	}

	result[FEATURES] = features
	result[BBOX] = fc.Bbox
	return result
}

// FeatureCollectionFromMap constructs a FeatureCollection from a map
// and returns its pointer
func FeatureCollectionFromMap(input map[string]interface{}) *FeatureCollection {
	result := NewFeatureCollection(nil)
	featuresIfc := input[FEATURES]
	switch it := featuresIfc.(type) {
	case []interface{}:
		for _, featureIfc := range it {
			if featureMap, ok := featureIfc.(map[string]interface{}); ok {
				feature := FeatureFromMap(featureMap)
				result.Features = append(result.Features, feature)
			}
		}
	case []map[string]interface{}:
	}
	if bboxIfc, ok := input[BBOX]; ok {
		result.Bbox, _ = NewBoundingBox(bboxIfc)
	}
	return result
}

// FillProperties iterates through all features to ensure that all properties
// are present on all features to meet the needs of some relational databases
func (fc *FeatureCollection) FillProperties() {
	properties := make(map[string]bool)

	// Loop 1: construct a set (actually a map)
	for _, feature := range fc.Features {
		for key := range feature.Properties {
			if _, ok := properties[key]; !ok {
				properties[key] = true
			}
		}
	}

	// Loop 2: make sure each feature has each property from the set
	for _, feature := range fc.Features {
		for key := range properties {
			if _, ok := feature.Properties[key]; !ok {
				feature.Properties[key] = nil
			}
		}
	}
}

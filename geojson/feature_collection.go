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

// Map returns a map of the FeatureCollection's members
// This may be useful in wrapping a Feature Collection with foreign members
func (fc *FeatureCollection) Map() map[string]interface{} {
	result := make(map[string]interface{})
	result["type"] = fc.Type
	result["features"] = fc.Features
	result["bbox"] = fc.Bbox
	return result
}

// FeatureCollectionFromMap constructs a FeatureCollection from a map
// and returns its pointer
func FeatureCollectionFromMap(input map[string]interface{}) *FeatureCollection {
	var result FeatureCollection
	result.Type = input["type"].(string)
	featuresIfc := input["features"]
	if featuresArray, ok := featuresIfc.([]interface{}); ok {
		for _, featureIfc := range featuresArray {
			if featureMap, ok := featureIfc.(map[string]interface{}); ok {
				feature := FeatureFromMap(featureMap)
				result.Features = append(result.Features, feature)
			}
		}
	}
	if bboxIfc, ok := input["bbox"]; ok {
		result.Bbox = NewBoundingBox(bboxIfc)
	}
	return &result
}

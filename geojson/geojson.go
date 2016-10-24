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
	"encoding/json"
	"io/ioutil"
)

// GeoJSON Constants
const (
	TYPE = "type"
	BBOX = "bbox"
)

// Parse parses a GeoJSON string into a GeoJSON object pointer
func Parse(bytes []byte) (interface{}, error) {
	var (
		result interface{}
		gj     map[string]interface{}
		err    error
	)
	if err = json.Unmarshal(bytes, &gj); err != nil {
		return nil, err
	}
	switch gj[TYPE] {
	case "Point":
		result, err = PointFromBytes(bytes)
	case "LineString":
		result, err = LineStringFromBytes(bytes)
	case "Polygon":
		result, err = PolygonFromBytes(bytes)
	case "MultiPoint":
		result, err = MultiPointFromBytes(bytes)
	case "MultiLineString":
		result, err = MultiLineStringFromBytes(bytes)
	case "MultiPolygon":
		result, err = MultiPolygonFromBytes(bytes)
	case "GeometryCollection":
		result, err = GeometryCollectionFromBytes(bytes)
	case "Feature":
		result, err = FeatureFromBytes(bytes)
	case "FeatureCollection":
		result, err = FeatureCollectionFromBytes(bytes)
	}
	return result, err
}

// ParseFile parses a file containing a GeoJSON string into a GeoJSON object pointer
func ParseFile(filename string) (interface{}, error) {
	var (
		bytes []byte
		err   error
	)
	if bytes, err = ioutil.ReadFile(filename); err != nil {
		return nil, err
	}
	return Parse(bytes)
}

// FromMap parses a map containing a GeoJSON object
// into a GeoJSON object pointer
func FromMap(input map[string]interface{}) interface{} {
	var (
		typeIfc interface{}
		ok      bool
	)

	if typeIfc, ok = input["type"]; ok {
		switch typeIfc.(string) {
		case FEATURE:
			return FeatureFromMap(input)
		case FEATURECOLLECTION:
			return FeatureCollectionFromMap(input)
		default:
			return newGeometry(input)
		}

	}
	return nil
}

// Write writes a GeoJSON object into a byte array
func Write(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}

// WriteFile writes a GeoJSON object to the file specified
func WriteFile(input interface{}, filename string) error {
	var (
		bytes []byte
		err   error
	)
	if bytes, err = Write(input); err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0666)
}

// Mapper is an interface for all GeoJSON objects to return itself as a Map
type Mapper interface {
	Map() map[string]interface{}
}

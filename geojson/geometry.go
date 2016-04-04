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

// GeoJSON Constants
const (
	COORDINATES        = "coordinates"
	POINT              = "Point"
	LINESTRING         = "LineString"
	POLYGON            = "Polygon"
	MULTIPOINT         = "MultiPoint"
	MULTILINESTRING    = "MultiLineString"
	MULTIPOLYGON       = "MultiPolygon"
	GEOMETRYCOLLECTION = "GeometryCollection"
)

// The Point object contains a single position
type Point struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
	// Bbox     Bbox     `json:"bbox, omitempty"`
}

// PointFromBytes constructs a point from a GeoJSON byte array
func PointFromBytes(bytes []byte) (Point, error) {
	var result Point
	err := json.Unmarshal(bytes, &result)
	return result, err
}

// NewPoint is the normal factory method for a Point
func NewPoint(coordinates []float64) *Point {
	return &Point{Type: POINT, Coordinates: coordinates}
}

// The LineString object contains a array of two or more positions
type LineString struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
	// Bbox     Bbox     `json:"bbox, omitempty"`
}

// LineStringFromBytes constructs a LineString from a GeoJSON byte array
func LineStringFromBytes(bytes []byte) (LineString, error) {
	var result LineString
	err := json.Unmarshal(bytes, &result)
	return result, err
}

// NewLineString is the normal factory method for a LineString
func NewLineString(coordinates [][]float64) *LineString {
	return &LineString{Type: LINESTRING, Coordinates: coordinates}
}

// The Polygon object contains a array of one or more linear rings
type Polygon struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
	// Bbox     Bbox     `json:"bbox, omitempty"`
}

// PolygonFromBytes constructs a Polygon from a GeoJSON byte array
func PolygonFromBytes(bytes []byte) (Polygon, error) {
	var result Polygon
	err := json.Unmarshal(bytes, &result)
	return result, err
}

// NewPolygon is the normal factory method for a Polygon
func NewPolygon(coordinates [][][]float64) *Polygon {
	return &Polygon{Type: POLYGON, Coordinates: coordinates}
}

// The MultiPoint object contains a array of one or more points
type MultiPoint struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
	// Bbox     Bbox     `json:"bbox, omitempty"`
}

// MultiPointFromBytes constructs a MultiPoint from a GeoJSON byte array
func MultiPointFromBytes(bytes []byte) (MultiPoint, error) {
	var result MultiPoint
	err := json.Unmarshal(bytes, &result)
	return result, err
}

// NewMultiPoint is the normal factory method for a MultiPoint
func NewMultiPoint(coordinates [][]float64) *MultiPoint {
	return &MultiPoint{Type: MULTIPOINT, Coordinates: coordinates}
}

// The MultiLineString object contains a array of one or more line strings
type MultiLineString struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
	// Bbox     Bbox     `json:"bbox, omitempty"`
}

// MultiLineStringFromBytes constructs a MultiLineString from a GeoJSON byte array
func MultiLineStringFromBytes(bytes []byte) (MultiLineString, error) {
	var result MultiLineString
	err := json.Unmarshal(bytes, &result)
	return result, err
}

// NewMultiLineString is the normal factory method for a LineString
func NewMultiLineString(coordinates [][][]float64) *MultiLineString {
	return &MultiLineString{Type: MULTILINESTRING, Coordinates: coordinates}
}

// The MultiPolygon object contains a array of one or more polygons
type MultiPolygon struct {
	Type        string          `json:"type"`
	Coordinates [][][][]float64 `json:"coordinates"`
	// Bbox     Bbox     `json:"bbox, omitempty"`
}

// MultiPolygonFromBytes constructs a MultiPolygon from a GeoJSON byte array
func MultiPolygonFromBytes(bytes []byte) (MultiPolygon, error) {
	var result MultiPolygon
	err := json.Unmarshal(bytes, &result)
	return result, err
}

// NewMultiPolygon is the normal factory method for a MultiPolygon
func NewMultiPolygon(coordinates [][][][]float64) *MultiPolygon {
	return &MultiPolygon{Type: MULTIPOLYGON, Coordinates: coordinates}
}

// The GeometryCollection object contains a array of one or more polygons
type GeometryCollection struct {
	Type       string        `json:"type"`
	Geometries []interface{} `json:"geometries"`
	// Bbox     Bbox     `json:"bbox, omitempty"`
}

// GeometryCollectionFromBytes constructs a GeometryCollection from a GeoJSON byte array
func GeometryCollectionFromBytes(bytes []byte) (GeometryCollection, error) {
	var result GeometryCollection
	err := json.Unmarshal(bytes, &result)
	var geometries []interface{}
	for inx := 0; inx < len(result.Geometries); inx++ {
		gmap := result.Geometries[inx].(map[string]interface{})
		geometry := NewGeometry(gmap)
		geometries = append(geometries, geometry)
		if err != nil {
			break
		}
	}
	if err != nil {
		result.Geometries = geometries
	}
	return result, err
}

// NewGeometryCollection is the normal factory method for a GeometryCollection
func NewGeometryCollection(geometries []interface{}) *GeometryCollection {
	return &GeometryCollection{Type: LINESTRING, Geometries: geometries}
}

// This quasi-recursive function determines drills into the
// multidimensional array of interfaces to build a proper
// coordinate array of the right dimension
func interfaceArrayToArray(input []interface{}) interface{} {
	var result interface{}
	var coords []float64
	var coords2 [][]float64
	var coords3 [][][]float64
	var coords4 [][][][]float64
	for inx := 0; inx < len(input); inx++ {
		switch curr1 := input[inx].(type) {
		case float64:
			coords = append(coords, curr1)
			result = coords
		case []interface{}:
			switch curr2 := interfaceArrayToArray(curr1).(type) {
			case []float64:
				coords2 = append(coords2, curr2)
				result = coords2
			case [][]float64:
				coords3 = append(coords3, curr2)
				result = coords3
			case [][][]float64:
				coords4 = append(coords4, curr2)
				result = coords4
			}
		}
	}
	return result
}

// NewGeometry constructs a Geometry from a map that represents a
// GeoJSON Geometry Object
func NewGeometry(input map[string]interface{}) interface{} {
	var result interface{}
	iType := input["type"].(string)
	coordinates := input[COORDINATES].([]interface{})
	switch iType {
	case POINT:
		var point Point
		point.Type = iType
		point.Coordinates = interfaceArrayToArray(coordinates).([]float64)
		result = point
	case LINESTRING:
		var lineString LineString
		lineString.Type = iType
		lineString.Coordinates = interfaceArrayToArray(coordinates).([][]float64)
		result = lineString
	case POLYGON:
		var polygon Polygon
		polygon.Type = iType
		polygon.Coordinates = interfaceArrayToArray(coordinates).([][][]float64)
		result = polygon
	case MULTIPOINT:
		var multiPoint MultiPoint
		multiPoint.Type = iType
		multiPoint.Coordinates = interfaceArrayToArray(coordinates).([][]float64)
		result = multiPoint
	case MULTILINESTRING:
		var multiLineString MultiLineString
		multiLineString.Type = iType
		multiLineString.Coordinates = interfaceArrayToArray(coordinates).([][][]float64)
		result = multiLineString
	case MULTIPOLYGON:
		var multiPolygon MultiPolygon
		multiPolygon.Type = iType
		multiPolygon.Coordinates = interfaceArrayToArray(coordinates).([][][][]float64)
		result = multiPolygon
	case GEOMETRYCOLLECTION:
		var geometryCollection GeometryCollection
		geometryCollection.Type = iType
		geometries := input["geometries"].([]interface{})
		for inx := 0; inx < len(geometries); inx++ {
			geometryCollection.Geometries = append(geometryCollection.Geometries,
				NewGeometry(geometries[inx].(map[string]interface{})))
		}
		result = geometryCollection
	}
	return result
}

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
	"fmt"
	"strconv"
	"strings"
)

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
	GEOMETRIES         = "geometries"
)

// The Point object contains a single position
type Point struct {
	Type        string      `json:"type"`
	Coordinates []float64   `json:"coordinates"`
	Bbox        BoundingBox `json:"bbox,omitempty"`
}

// PointFromBytes constructs a point from a GeoJSON byte array
// and returns its pointer
func PointFromBytes(bytes []byte) (*Point, error) {
	var result Point
	err := json.Unmarshal(bytes, &result)
	return &result, err
}

// ForceBbox returns a bounding box, creating one by brute force if needed
func (point Point) ForceBbox() BoundingBox {
	if len(point.Bbox) > 0 {
		return point.Bbox
	}
	if bbox, err := NewBoundingBox(point.Coordinates); err == nil {
		return bbox
	}
	return []float64{}
}

// String returns the string representation
func (point Point) String() string {
	var result string
	if bytes, err := json.Marshal(point); err == nil {
		result = string(bytes)
	} else {
		result = err.Error()
	}
	return result
}

// Map returns a map of the Geometry's members
func (point Point) Map() map[string]interface{} {
	result := make(map[string]interface{})
	result[COORDINATES] = point.Coordinates
	result[TYPE] = POINT
	return result
}

// NewPoint is the normal factory method for a Point
func NewPoint(coordinates []float64) *Point {
	if len(coordinates) > 1 {
		return &Point{Type: POINT, Coordinates: coordinates}
	}
	return nil
}

// WKT returns the Well Known Text representation of the point
func (point Point) WKT() string {
	var result string
	switch len(point.Coordinates) {
	case 2:
		result = fmt.Sprintf("POINT (%v)", array1ToWKTCoordinates(point.Coordinates))
	case 3:
		result = fmt.Sprintf("POINT Z (%v)", array1ToWKTCoordinates(point.Coordinates))
	case 4:
		result = fmt.Sprintf("POINT ZM (%v)", array1ToWKTCoordinates(point.Coordinates))
	default:
		result = "POINT EMPTY"
	}
	return result
}

func array1ToWKTCoordinates(input []float64) string {
	var result string
	switch len(input) {
	case 2:
		result = fmt.Sprintf("%f %f", input[0], input[1])
	case 3:
		result = fmt.Sprintf("%f %f %f", input[0], input[1], input[2])
	case 4:
		result = fmt.Sprintf("%f %f %f %f", input[0], input[1], input[2], input[3])
	}
	return result
}

// The LineString object contains a array of two or more positions
type LineString struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
	Bbox        BoundingBox `json:"bbox,omitempty"`
}

// LineStringFromBytes constructs a LineString from a GeoJSON byte array
// and returns its pointer
func LineStringFromBytes(bytes []byte) (*LineString, error) {
	var result LineString
	err := json.Unmarshal(bytes, &result)
	return &result, err
}

// ForceBbox returns a bounding box, creating one by brute force if needed
func (ls LineString) ForceBbox() BoundingBox {
	if len(ls.Bbox) > 0 {
		return ls.Bbox
	}
	if bbox, err := NewBoundingBox(ls.Coordinates); err == nil {
		return bbox
	}
	return []float64{}
}

// String returns the string representation
func (ls LineString) String() string {
	var result string
	if bytes, err := json.Marshal(ls); err == nil {
		result = string(bytes)
	} else {
		result = err.Error()
	}
	return result
}

// Map returns a map of the Geometry's members
func (ls LineString) Map() map[string]interface{} {
	result := make(map[string]interface{})
	result[COORDINATES] = ls.Coordinates
	result[TYPE] = LINESTRING
	return result
}

// NewLineString is the normal factory method for a LineString
func NewLineString(coordinates [][]float64) *LineString {
	return &LineString{Type: LINESTRING, Coordinates: coordinates}
}

// The Polygon object contains a array of one or more linear rings
type Polygon struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
	Bbox        BoundingBox   `json:"bbox,omitempty"`
}

// PolygonFromBytes constructs a Polygon from a GeoJSON byte array
// and returns its pointer
func PolygonFromBytes(bytes []byte) (*Polygon, error) {
	var result Polygon
	err := json.Unmarshal(bytes, &result)
	return &result, err
}

// ForceBbox returns a bounding box, creating one by brute force if needed
func (polygon Polygon) ForceBbox() BoundingBox {
	if len(polygon.Bbox) > 0 {
		return polygon.Bbox
	}
	if bbox, err := NewBoundingBox(polygon.Coordinates); err == nil {
		return bbox
	}
	return []float64{}
}

// String returns the string representation
func (polygon Polygon) String() string {
	var result string
	if bytes, err := json.Marshal(polygon); err == nil {
		result = string(bytes)
	} else {
		result = err.Error()
	}
	return result
}

// Map returns a map of the Geometry's members
func (polygon Polygon) Map() map[string]interface{} {
	result := make(map[string]interface{})
	result[COORDINATES] = polygon.Coordinates
	result[TYPE] = POLYGON
	return result
}

// NewPolygon is the normal factory method for a Polygon
func NewPolygon(coordinates [][][]float64) *Polygon {
	return &Polygon{Type: POLYGON, Coordinates: coordinates}
}

// The MultiPoint object contains a array of one or more points
type MultiPoint struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
	Bbox        BoundingBox `json:"bbox,omitempty"`
}

// MultiPointFromBytes constructs a MultiPoint from a GeoJSON byte array
// and returns its pointer
func MultiPointFromBytes(bytes []byte) (*MultiPoint, error) {
	var result MultiPoint
	err := json.Unmarshal(bytes, &result)
	return &result, err
}

// ForceBbox returns a bounding box, creating one by brute force if needed
func (mp MultiPoint) ForceBbox() BoundingBox {
	if len(mp.Bbox) > 0 {
		return mp.Bbox
	}
	if bbox, err := NewBoundingBox(mp.Coordinates); err == nil {
		return bbox
	}
	return []float64{}
}

// String returns the string representation
func (mp MultiPoint) String() string {
	var result string
	if bytes, err := json.Marshal(mp); err == nil {
		result = string(bytes)
	} else {
		result = err.Error()
	}
	return result
}

// Map returns a map of the Geometry's members
func (mp MultiPoint) Map() map[string]interface{} {
	result := make(map[string]interface{})
	result[COORDINATES] = mp.Coordinates
	result[TYPE] = MULTIPOINT
	return result
}

// NewMultiPoint is the normal factory method for a MultiPoint
func NewMultiPoint(coordinates [][]float64) *MultiPoint {
	return &MultiPoint{Type: MULTIPOINT, Coordinates: coordinates}
}

// The MultiLineString object contains a array of one or more line strings
type MultiLineString struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
	Bbox        BoundingBox   `json:"bbox,omitempty"`
}

// MultiLineStringFromBytes constructs a MultiLineString from a GeoJSON byte array
// and returns its pointer
func MultiLineStringFromBytes(bytes []byte) (*MultiLineString, error) {
	var result MultiLineString
	err := json.Unmarshal(bytes, &result)
	return &result, err
}

// ForceBbox returns a bounding box, creating one by brute force if needed
func (mls MultiLineString) ForceBbox() BoundingBox {
	if len(mls.Bbox) > 0 {
		return mls.Bbox
	}
	if bbox, err := NewBoundingBox(mls.Coordinates); err == nil {
		return bbox
	}
	return []float64{}
}

// String returns the string representation
func (mls MultiLineString) String() string {
	var result string
	if bytes, err := json.Marshal(mls); err == nil {
		result = string(bytes)
	} else {
		result = err.Error()
	}
	return result
}

// Map returns a map of the Geometry's members
func (mls MultiLineString) Map() map[string]interface{} {
	result := make(map[string]interface{})
	result[COORDINATES] = mls.Coordinates
	result[TYPE] = MULTILINESTRING
	return result
}

// NewMultiLineString is the normal factory method for a MultiLineString
func NewMultiLineString(coordinates [][][]float64) *MultiLineString {
	return &MultiLineString{Type: MULTILINESTRING, Coordinates: coordinates}
}

// The MultiPolygon object contains a array of one or more polygons
type MultiPolygon struct {
	Type        string          `json:"type"`
	Coordinates [][][][]float64 `json:"coordinates"`
	Bbox        BoundingBox     `json:"bbox,omitempty"`
}

// MultiPolygonFromBytes constructs a MultiPolygon from a GeoJSON byte array
// and returns its pointer
func MultiPolygonFromBytes(bytes []byte) (*MultiPolygon, error) {
	var result MultiPolygon
	err := json.Unmarshal(bytes, &result)
	return &result, err
}

// ForceBbox returns a bounding box, creating one by brute force if needed
func (mp MultiPolygon) ForceBbox() BoundingBox {
	if len(mp.Bbox) > 0 {
		return mp.Bbox
	}
	if bbox, err := NewBoundingBox(mp.Coordinates); err == nil {
		return bbox
	}
	return []float64{}
}

// String returns the string representation
func (mp MultiPolygon) String() string {
	var result string
	if bytes, err := json.Marshal(mp); err == nil {
		result = string(bytes)
	} else {
		result = err.Error()
	}
	return result
}

// Map returns a map of the Geometry's members
func (mp MultiPolygon) Map() map[string]interface{} {
	result := make(map[string]interface{})
	result[COORDINATES] = mp.Coordinates
	result[TYPE] = MULTIPOLYGON
	return result
}

// NewMultiPolygon is the normal factory method for a MultiPolygon
func NewMultiPolygon(coordinates [][][][]float64) *MultiPolygon {
	return &MultiPolygon{Type: MULTIPOLYGON, Coordinates: coordinates}
}

// The GeometryCollection object contains a array of one or more polygons
type GeometryCollection struct {
	Type       string        `json:"type"`
	Geometries []interface{} `json:"geometries"`
	Bbox       BoundingBox   `json:"bbox,omitempty"`
}

// GeometryCollectionFromBytes constructs a GeometryCollection from a GeoJSON byte array
func GeometryCollectionFromBytes(bytes []byte) (*GeometryCollection, error) {
	var result GeometryCollection
	err := json.Unmarshal(bytes, &result)
	var geometries []interface{}
	for _, curr := range result.Geometries {
		gmap := curr.(map[string]interface{})
		geometry := newGeometry(gmap)
		geometries = append(geometries, geometry)
	}
	result.Geometries = geometries
	return &result, err
}

// ForceBbox returns a bounding box, creating one by brute force if needed
func (gc GeometryCollection) ForceBbox() BoundingBox {
	if len(gc.Bbox) > 0 {
		return gc.Bbox
	}
	var result BoundingBox
	for _, geometry := range gc.Geometries {
		if bboxIfc, ok := geometry.(BoundingBoxIfc); ok {
			result = mergeBboxes(result, bboxIfc.ForceBbox())
		}
	}
	return result
}

// String returns the string representation
func (gc GeometryCollection) String() string {
	var result string
	if bytes, err := json.Marshal(gc); err == nil {
		result = string(bytes)
	} else {
		result = err.Error()
	}
	return result
}

// Map returns a map of the Geometry's members
func (gc GeometryCollection) Map() map[string]interface{} {
	result := make(map[string]interface{})
	geometries := make([]map[string]interface{}, len(gc.Geometries))
	for inx, geometry := range gc.Geometries {
		switch gt := geometry.(type) {
		case map[string]interface{}:
			geometries[inx] = gt
		case Mapper:
			geometries[inx] = gt.Map()
		}
	}
	result[GEOMETRIES] = geometries
	result[TYPE] = GEOMETRYCOLLECTION
	return result
}

// NewGeometryCollection is the normal factory method for a GeometryCollection
func NewGeometryCollection(geometries []interface{}) *GeometryCollection {
	if geometries == nil {
		geometries = make([]interface{}, 0)
	}
	return &GeometryCollection{Type: GEOMETRYCOLLECTION, Geometries: geometries}
}

// This quasi-recursive function determines drills into the
// multidimensional array of interfaces to build a proper
// coordinate array of the right dimension
func interfaceToArray(input interface{}) interface{} {
	var (
		result  interface{}
		coords  []float64
		coords2 [][]float64
		coords3 [][][]float64
		coords4 [][][][]float64
	)
	switch it := input.(type) {
	case []float64, [][]float64, [][][]float64, [][][][]float64:
		result = it
	case []interface{}:
		for inx := 0; inx < len(it); inx++ {
			switch curr1 := it[inx].(type) {
			case float64:
				coords = append(coords, curr1)
				result = coords
			case []interface{}:
				switch curr2 := interfaceToArray(curr1).(type) {
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
	}
	return result
}

// newGeometry constructs a Geometry from an interface that represents a
// GeoJSON Geometry Object
func newGeometry(input interface{}) interface{} {
	var (
		result      interface{}
		coordinates interface{}
	)
	switch it := input.(type) {
	case map[string]interface{}:
		if _, ok := it[COORDINATES]; ok {
			coordinates = interfaceToArray(it[COORDINATES])
		}
		iType := it[TYPE].(string)
		switch iType {
		case POINT:
			result = NewPoint(coordinates.([]float64))
		case LINESTRING:
			result = NewLineString(coordinates.([][]float64))
		case POLYGON:
			result = NewPolygon(coordinates.([][][]float64))
		case MULTIPOINT:
			result = NewMultiPoint(coordinates.([][]float64))
		case MULTILINESTRING:
			result = NewMultiLineString(coordinates.([][][]float64))
		case MULTIPOLYGON:
			result = NewMultiPolygon(coordinates.([][][][]float64))
		case GEOMETRYCOLLECTION:
			geometries := it["geometries"].([]interface{})
			for inx, geometry := range geometries {
				geometries[inx] = newGeometry(geometry)
			}
			result = NewGeometryCollection(geometries)
		}
	case *Point, *LineString, *Polygon, *MultiPoint, *MultiLineString, *MultiPolygon, *GeometryCollection:
		result = it
	}
	return result
}

// ToGeometryArray takes a GeoJSON object and returns an array of
// its constituent geometry objects
func ToGeometryArray(gjObject interface{}) []interface{} {
	var result []interface{}
	switch typedGJ := gjObject.(type) {
	case *FeatureCollection:
		// re-enter with dereferenced pointer
		result = ToGeometryArray(*typedGJ)
	case FeatureCollection:
		for _, current := range typedGJ.Features {
			result = append(result, current.Geometry)
		}
	case *Feature:
		// re-enter with dereferenced pointer
		result = ToGeometryArray(*typedGJ)
	case Feature:
		result = append(result, typedGJ.Geometry)
	case *interface{}:
		// re-enter with dereferenced pointer
		result = ToGeometryArray(*typedGJ)
	default:
		// Hopefully this is a Geometry object
		result = append(result, typedGJ)
	}
	return result
}

// WKT returns a GeoJSON object based on the Well-Known Text input
func WKT(input string) interface{} {
	parts := strings.SplitN(input, " ", 2)
	name := parts[0]
	predicate := parseWKTPredicate(parts[1])

	switch name {
	case "POINT":
		if coords, ok := predicate.([]float64); ok {
			return NewPoint(coords)
		}
		return NewPoint(nil)
	case "POLYGON":
		if coords, ok := predicate.([][][]float64); ok {
			return NewPolygon(coords)
		}
		return NewPolygon(nil)
	case "LINESTRING":
		if coords, ok := predicate.([][]float64); ok {
			return NewLineString(coords)
		}
		return NewLineString(nil)
	case "MULTIPOINT":
		switch mpp := predicate.(type) {
		case [][]float64:
			return NewMultiPoint(mpp)
		default:
			fmt.Printf("%T\n", mpp)
			return NewMultiPoint(nil)
		}
	}
	return nil
}

func parseWKTPredicate(input string) interface{} {

	parts := strings.SplitN(input, "(", 2)
	switch string(parts[1][0]) {
	case "(":
		var result2 [][]float64
		var result3 [][][]float64
		for _, part := range strings.Split(parts[1], ")") {
			if strings.Index(part, "(") >= 0 {
				currIfc := parseWKTPredicate(part)
				switch curr := currIfc.(type) {
				case []float64:
					result2 = append(result2, curr)
				case [][]float64:
					result3 = append(result3, curr)
				}
				// part = strings.Split(part, "(")[1]
				// result = append(result, parseWKTCoordinates(strings.Trim(part, " )")))
			}
		}
		if result2 != nil {
			return result2
		} else if result3 != nil {
			return result3
		}
	default:
		switch strings.Index(parts[1], ",") {
		case -1:
			parts = strings.SplitN(parts[1], ")", 2)
			return parseWKTCoordinates(parts[0])
		default:
			var result [][]float64
			for _, part := range strings.Split(parts[1], ",") {
				result = append(result, parseWKTCoordinates(strings.Trim(part, " )")))
			}
			return result
		}
	}
	return nil
}

func parseWKTCoordinates(input string) []float64 {
	var coords []float64
	for _, part := range strings.Split(input, " ") {
		coord, _ := strconv.ParseFloat(part, 10)
		coords = append(coords, coord)
	}
	return coords
}

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
	"errors"
	"strconv"
	"strings"
)

// BoundingBoxIfc is for objects that have a bounding box property
type BoundingBoxIfc interface {
	ForceBbox() BoundingBox
}

// The BoundingBox type supports bbox elements in GeoJSON
// Note that since this is an array,
// it is passed by value instead of pointer
// (unlike other GeoJSON objects)
type BoundingBox []float64

// NewBoundingBox creates a BoundingBox from a large number of inputs
// including a string and an n-dimensional coordinate array
func NewBoundingBox(input interface{}) (BoundingBox, error) {
	var (
		result, bbox2 BoundingBox
		err           error
		coordValue    float64
	)

	switch inputType := input.(type) {
	case []string:
		for _, coord := range inputType {
			if coordValue, err = strconv.ParseFloat(coord, 64); err == nil {
				result = append(result, coordValue)
			} else {
				return result, errors.New("Failed to parse bounding box: " + err.Error())
			}
		}
	case string:
		if inputType != "" {
			return NewBoundingBox(strings.Split(inputType, ","))
		}
	case []float64:
		switch len(inputType) {
		case 0, 1:
			// No op
		case 2:
			result = append(inputType, inputType[:]...)
			// Drop extraneous coordinates like measurements that do not pertain to BBoxes
		default:
			result = append(inputType[0:3], inputType[0:3]...)
		}
	case [][]float64:
		for _, curr := range inputType {
			if bbox2, err = NewBoundingBox(curr); err == nil {
				result = mergeBboxes(result, bbox2)
			} else {
				return result, err
			}
		}
	case [][][]float64:
		for _, curr := range inputType {
			if bbox2, err = NewBoundingBox(curr); err == nil {
				result = mergeBboxes(result, bbox2)
			} else {
				return result, err
			}
		}
	case [][][][]float64:
		for _, curr := range inputType {
			if bbox2, err = NewBoundingBox(curr); err == nil {
				result = mergeBboxes(result, bbox2)
			} else {
				return result, err
			}
		}
	case []BoundingBox:
		for _, curr := range inputType {
			result = mergeBboxes(result, curr)
		}
	}

	return result, result.Valid()
}

func mergeBboxes(first, second BoundingBox) BoundingBox {
	length := len(first)
	if length == 0 {
		return second
	}
	// Generally this covers the case where second is empty for some reason
	if length != len(second) {
		return first
	}

	// For X, we must consider the antimeridian case
	if (first[0] == -180) && (second[length/2] == 180) {
		first[0] = second[0]
	} else if (first[length/2] == 180) && (second[0] == -180) {
		first[length/2] = second[length/2]
	} else {
		if second[0] < first[0] {
			first[0] = second[0]
		}
		if second[length/2] > first[length/2] {
			first[length/2] = second[length/2]
		}
	}

	// Consider the minimum values
	for inx := 1; inx < length/2; inx++ {
		if second[inx] < first[inx] {
			first[inx] = second[inx]
		}
	}
	// Consider the maximum values
	for inx := 1 + length/2; inx < length; inx++ {
		if second[inx] > first[inx] {
			first[inx] = second[inx]
		}
	}
	return first
}

// Equals returns true if all points in the bounding boxes are equal
func (bb BoundingBox) Equals(test BoundingBox) bool {
	bblen := len(bb)
	testlen := len(test)
	if (bblen == 0) && (testlen == 0) {
		return true
	}
	if (bblen == 0) || (testlen == 0) || (bblen != testlen) {
		return false
	}
	for inx := 0; inx < bblen; inx++ {
		if bb[inx] != test[inx] {
			return false
		}
	}
	return true
}

// Overlaps returns true if the interiors of the two bounding boxes
// have any area in common
func (bb BoundingBox) Overlaps(test BoundingBox) bool {
	bblen := len(bb)
	testlen := len(test)

	if (bblen == 0) || (testlen == 0) || (bblen != testlen) {
		return false
	}
	result := true
	bbDimensions := bblen / 2
	if bb.Antimeridian() && test.Antimeridian() {
		// no op
	} else if bb.Antimeridian() {
		result = result && ((test[bbDimensions] >= bb[0]) && (test[bbDimensions] <= 180) ||
			(test[0] <= bb[bbDimensions]) && (test[0] >= -180))
	} else if test.Antimeridian() {
		result = result && ((bb[bbDimensions] >= test[0]) && (bb[bbDimensions] <= 180) ||
			(bb[0] <= test[bbDimensions]) && (bb[0] >= -180))
	} else {
		result = result && (bb[0] < test[bbDimensions]) && (bb[bbDimensions] > test[0])
	}
	result = result && (bb[1] < test[1+bbDimensions]) && (bb[1+bbDimensions] > test[1])
	if bbDimensions > 2 {
		result = result && (bb[2] < test[2+bbDimensions]) && (bb[2+bbDimensions] > test[2])
	}
	return result
}

// String returns a string representation as minx,miny,maxx,maxy
func (bb BoundingBox) String() string {
	var result string
	switch len(bb) {
	case 4:
		result = strconv.FormatFloat(bb[0], 'f', 3, 32) + "," +
			strconv.FormatFloat(bb[1], 'f', 3, 32) + "," +
			strconv.FormatFloat(bb[2], 'f', 3, 32) + "," +
			strconv.FormatFloat(bb[3], 'f', 3, 32)
	case 6:
		result = strconv.FormatFloat(bb[0], 'f', 3, 32) + "," +
			strconv.FormatFloat(bb[1], 'f', 3, 32) + "," +
			strconv.FormatFloat(bb[2], 'f', 3, 32) + "," +
			strconv.FormatFloat(bb[3], 'f', 3, 32) + "," +
			strconv.FormatFloat(bb[4], 'f', 3, 32) + "," +
			strconv.FormatFloat(bb[5], 'f', 3, 32)
	}
	return result
}

// Valid returns nil if the bounding box is valid
// or an error object if it is invalid
func (bb BoundingBox) Valid() error {
	switch len(bb) {
	case 0:
		return nil
	case 4:
		if bb[1] > bb[3] {
			return errors.New("Bounding Box values must be in south-westerly to north-easterly order.")
		}
		return nil
	case 6:
		if bb[1] > bb[4] || bb[2] > bb[5] {
			return errors.New("Bounding Box values must be in south-westerly to north-easterly order.")
		}
		return nil
	}
	return errors.New("Bounding Box must have 0, 4, or 6 values.")
}

// Antimeridian returns true if the BoundingBox crosses the antimeridian
func (bb BoundingBox) Antimeridian() bool {
	if bb.Valid() == nil {
		switch len(bb) {
		case 4:
			if bb[0] > bb[2] {
				return true
			}
		case 6:
			if bb[0] > bb[3] {
				return true
			}
		}
	}
	return false
}

// Centroid returns the center of the BoundingBox,
// or nil if it has no coordinates or is otherwise invalid
func (bb BoundingBox) Centroid() *Point {
	var result *Point
	if bb.Valid() == nil {
		switch len(bb) {
		case 0:
		case 4:
			var coordinates [2]float64
			coordinates[0] = 0.5 * (bb[0] + bb[2])
			coordinates[1] = 0.5 * (bb[1] + bb[3])
			result = NewPoint(coordinates[0:])
		case 6:
			var coordinates [3]float64
			coordinates[0] = 0.5 * (bb[0] + bb[3])
			coordinates[1] = 0.5 * (bb[1] + bb[4])
			coordinates[2] = 0.5 * (bb[2] + bb[5])
			result = NewPoint(coordinates[0:])
		}
	}
	return result
}

// Geometry returns the BoundingBox as a GeoJSON object
// (point, line, or polygon)
// if the BoundingBox is two-dimensional
func (bb BoundingBox) Geometry() interface{} {
	var result interface{}
	switch len(bb) {
	case 4:
		if (bb[0] == bb[2]) && (bb[1] == bb[3]) {
			result = NewPoint(bb[0:2])
		} else {
			coordinates := make([][][]float64, 1)
			coordinates[0] = make([][]float64, 5)
			coordinates[0][0] = make([]float64, 2)
			coordinates[0][0][0] = bb[0]
			coordinates[0][0][1] = bb[1]
			coordinates[0][1] = make([]float64, 2)
			coordinates[0][1][0] = bb[2]
			coordinates[0][1][1] = bb[1]
			coordinates[0][2] = make([]float64, 2)
			coordinates[0][2][0] = bb[2]
			coordinates[0][2][1] = bb[3]
			coordinates[0][3] = make([]float64, 2)
			coordinates[0][3][0] = bb[0]
			coordinates[0][3][1] = bb[3]
			coordinates[0][4] = make([]float64, 2)
			coordinates[0][4][0] = bb[0]
			coordinates[0][4][1] = bb[1]
			result = NewPolygon(coordinates)
		}
	}
	return result
}

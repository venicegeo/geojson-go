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
	case string:
		if inputType != "" {
			coords := strings.Split(inputType, ",")
			for _, coord := range coords {
				if coordValue, err = strconv.ParseFloat(coord, 64); err == nil {
					result = append(result, coordValue)
				} else {
					return result, errors.New("Failed to parse bounding box: " + err.Error())
				}
			}
		}
	case []float64:
		result = append(inputType, inputType[:]...)
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
	}

	if result.Valid() {
		return result, nil
	}
	return []float64{}, errors.New("Bounding box has invalid number of coordinates")
}

func mergeBboxes(first, second BoundingBox) BoundingBox {
	length := len(first)
	if length == 0 {
		return second
	}
	for inx := 0; inx < length/2; inx++ {
		if second[inx] < first[inx] {
			first[inx] = second[inx]
		}
	}
	for inx := length / 2; inx < length; inx++ {
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
	for inx := 0; inx < bbDimensions; inx++ {
		result = result && (bb[inx] < test[inx+bbDimensions]) && (bb[inx+bbDimensions] > test[inx])
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

// Valid returns true if the bounding box is valid
func (bb BoundingBox) Valid() bool {
	switch len(bb) {
	case 0:
		return true
	case 4:
		return true
	case 6:
		return true
	}
	return false
}

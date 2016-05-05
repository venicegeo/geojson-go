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

// BoundingBoxIfc is for objects that have a bounding box property
type BoundingBoxIfc interface {
	ForceBbox() BoundingBox
}

// The BoundingBox type supports bbox elements in GeoJSON
type BoundingBox []float64

func bboxFromCoords(input interface{}) BoundingBox {
	var (
		result BoundingBox
	)
	switch inputType := input.(type) {
	case []float64:
		result = append(inputType, inputType[:]...)
	case [][]float64:
		for _, curr := range inputType {
			result = mergeBboxes(result, bboxFromCoords(curr))
		}
	case [][][]float64:
		for _, curr := range inputType {
			result = mergeBboxes(result, bboxFromCoords(curr))
		}
	case [][][][]float64:
		for _, curr := range inputType {
			result = mergeBboxes(result, bboxFromCoords(curr))
		}
	}
	return result
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

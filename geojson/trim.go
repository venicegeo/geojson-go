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

import "math"

// Trim allows something with coordinates to have its values trimmed
var Trim interface {
	Trim(precision int)
}

func trim1(coordinates []float64, precision int) []float64 {
	var result []float64
	for _, coordinate := range coordinates {
		result = append(result, math.Pow10(-precision)*math.Trunc(math.Pow10(precision)*coordinate))
	}
	return result
}
func trim2(coordinates [][]float64, precision int) [][]float64 {
	var result [][]float64
	for _, coordinate := range coordinates {
		result = append(result, trim1(coordinate, precision))
	}
	return result
}
func trim3(coordinates [][][]float64, precision int) [][][]float64 {
	var result [][][]float64
	for _, coordinate := range coordinates {
		result = append(result, trim2(coordinate, precision))
	}
	return result
}
func trim4(coordinates [][][][]float64, precision int) [][][][]float64 {
	var result [][][][]float64
	for _, coordinate := range coordinates {
		result = append(result, trim3(coordinate, precision))
	}
	return result
}

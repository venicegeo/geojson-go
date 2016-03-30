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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/venicegeo/geojson-go/geojson"
)

func test() {
	process("test/point.geojson")
	process("test/linestring.geojson")
	process("test/polygon.geojson")
	process("test/polygon-dateline.geojson")
	process("test/polygon-hole.geojson")
	process("test/multipoint.geojson")
	process("test/multilinestring.geojson")
	process("test/multipolygon.geojson")
	process("test/geometrycollection.geojson")
	process("test/sample.geojson")
}

func process(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("File error: %v\n", err)
	}
	geojson, err := geojson.Parse([]byte(file))
	if err != nil {
		log.Panicf("Parse error: %v\n", err)
	}
	fmt.Printf("%#v\n", geojson)
	var bytes []byte
	bytes, err = json.Marshal(geojson)
	if err != nil {
		log.Panicf("Marshal error: %v\n", err)
	}
	fmt.Printf("%v\n", string(bytes))
}

func main() {

	var args = os.Args[1:]
	if len(args) > 0 {
		if args[0] == "-t" {
			test()
		} else {
			for inx := 0; inx < len(args); inx++ {
				process(args[inx])
			}
		}
	}
}

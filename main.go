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
	"fmt"
	"log"
	"os"

	"github.com/venicegeo/geojson-go/geojson"
)

func process(filename string) {
	var (
		gj    interface{}
		err   error
		bytes []byte
	)
	if gj, err = geojson.ParseFile(filename); err != nil {
		log.Panicf("Parse error: %v\n", err)
	}
	fmt.Printf("%T: %#v\n", gj, gj)

	if bytes, err = geojson.Write(gj); err != nil {
		log.Panicf("Write error: %v\n", err)
	}
	fmt.Printf("%v\n", string(bytes))
}

func main() {

	var args = os.Args[1:]
	if len(args) > 0 {
		for inx := range args {
			process(args[inx])
		}
	}
}

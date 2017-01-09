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
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// FromWFS returns a Feature Collection from the WFS provided.
// This is a convenience method only and not intended for serious processing.
// This function does not currently support WFS layers with large numbers of features.
func FromWFS(wfsURL, featureType string) (*FeatureCollection, error) {
	var (
		result   *FeatureCollection
		err      error
		request  *http.Request
		response *http.Response
		b        []byte
		ok       bool
	)

	v := url.Values{}
	v.Set("service", "wfs")
	v.Set("count", "9999")
	v.Set("outputFormat", "application/json")
	v.Set("version", "2.0.0")
	v.Set("request", "GetFeature")
	v.Set("typeName", featureType)

	qurl := wfsURL + "?" + v.Encode()

	if request, err = http.NewRequest("GET", qurl, nil); err != nil {
		return result, err
	}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := http.Client{Transport: transport}
	if response, err = httpClient.Do(request); err != nil {
		return result, err
	}

	defer response.Body.Close()
	b, _ = ioutil.ReadAll(response.Body)

	// Check for HTTP errors
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return result, fmt.Errorf("Received %v: \"%v\" when performing a GetFeature request on %v\n%v", response.StatusCode, response.Status, qurl, string(b))
	}

	gjIfc, _ := Parse(b)
	if result, ok = gjIfc.(*FeatureCollection); ok {
		return result, nil
	}
	return result, fmt.Errorf("Did not receive valid GeoJSON on request %v", qurl)
}

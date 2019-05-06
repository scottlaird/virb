/*

  Copyright 2019 Google LLC

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package virb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Fetch talks to the camera over the network.  It takes a
// request structure, serializes it to JSON, and passes it to the
// camera over an HTTP POST to /virb, and then parses the JSON result
// into the provided resp structure.
func fetch(host string, req interface{}, resp interface{}) error {
	url := fmt.Sprintf("http://%s/virb", host)
	cmd, err := json.Marshal(req)
	if err != nil {
		return err
	}
	//fmt.Printf("** cmd is <%s>\n", string(cmd))
	hresp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewReader(cmd))
	if err != nil {
		return err
	}
	//fmt.Printf("Received response:\n")
	defer hresp.Body.Close()
	body, err := ioutil.ReadAll(hresp.Body)
	if err != nil {
		return err
	}
	//fmt.Printf("<%s>\n", string(body))
	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}
	return nil
}

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

type SensorsRequest struct {
	Command string `json:"command"`
	Names []string `json:"names,omitempty"`
}
type SensorsResponse struct {
	Sensors []Sensors_Object
	Result int
}

type Sensors_Object struct {
	Data string  // Provided example is "0.731707"
	DataType string // "double" or "int"
	HasData int
	Name string
	Type string
	Units string
}

func Sensors(host string, sensors []string) (*SensorsResponse, error) {
	var resp SensorsResponse
	var req SensorsRequest
	req.Command="sensors"
	req.Names = sensors

	err := fetch(host, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

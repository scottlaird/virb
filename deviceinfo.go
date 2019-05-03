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

type DeviceInfoRequest struct {
	Command string `json:"command"`
}
type DeviceInfoResponse struct {
	DeviceInfo []DeviceInfo_Object
	Result     int
}

type DeviceInfo_Object struct {
	DeviceId   int
	Firmware   int
	MacAddress string
	Model      string
	PartNumber string
	Type       string
}

func DeviceInfo(host string) (*DeviceInfoResponse, error) {
	var resp DeviceInfoResponse
	var req DeviceInfoRequest
	req.Command = "deviceInfo"

	err := fetch(host, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

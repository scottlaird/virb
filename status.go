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

type StatusRequest struct {
	Command string `json:"command"`
}
type StatusResponse struct {
	ApiMin int
	ApiMax int
	BatteryLevel float64
	BatteryChargingState int
	TotalSpace int64
	AvailableSpace int64
	GpsAccuracy int
	GpsLatitude float64
	GpsLongitude float64
	Altitude float64
	Speed float64
	AntSensor int
	BtSensor int
	BtHeadset int
	WifiSensor int
	RecordingTimeRemaining int
	PhotosRemaining int
	PhotoCount int
	RecordingTime int
	WifiSignalStrength int
	WifiMode string
	State string
	LastMediaEventTime int
	Result int
}
// Status sends a Status command request to the camera and returns a StatusResponse if successful.
func Status(host string) (*StatusResponse, error) {
	var resp StatusResponse
	var req StatusRequest
	req.Command="status"

	err := fetch(host, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}


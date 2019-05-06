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

type LivePreviewRequest struct {
	Command               string `json:"command"`
	StreamType            string `json:"streamType"` // Required, "rtp".
	MaxResolutionVertical string `json:"maxResolutionVertical,omitempty"`
	LiveStreamActive      string `json:"liveStreamActive,omitempty"`
}
type LivePreviewResponse struct {
	Url    string
	Result int
}

func LivePreview(host string, streamType string, maxResolution int, active string) (*LivePreviewResponse, error) {
	var resp LivePreviewResponse
	var req LivePreviewRequest
	req.Command = "livePreview"
	req.StreamType = streamType
	if maxResolution > 0 {
		req.MaxResolutionVertical = string(maxResolution)
	}
	req.LiveStreamActive = active

	err := fetch(host, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

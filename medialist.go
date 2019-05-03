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

type MediaListRequest struct {
	Command string `json:"command"`
	Path string `json:"path,omitempty"`
}
type MediaListResponse struct {
	Media []MediaList_MediaObject
	Result int
}
type MediaList_MediaObject struct {
	Date int64
	Duration float64
	Fav string
	FileSize int64
	FitURL string
	GroupID int
	Index int
	LensMode string
	LowResVideoPath string
	Name string
	ThumbURL string
	Type string
	Url string
}

// Status sends a Status command request to the camera and returns a StatusResponse if successful.
func MediaList(host string, path string) (*MediaListResponse, error) {
	var resp MediaListResponse
	var req MediaListRequest
	req.Command="mediaList"
	if len(path)>0 {
		req.Path = path
	}

	err := fetch(host, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}


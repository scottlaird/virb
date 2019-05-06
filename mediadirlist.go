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

type MediaDirListRequest struct {
	Command string `json:"command"`
}
type MediaDirListResponse struct {
	MediaDirs []MediaDirs_Object
	Result    int
}

type MediaDirs_Object struct {
	Date int
	Path string
	Type string
}

func MediaDirList(host string) (*MediaDirListResponse, error) {
	var resp MediaDirListResponse
	var req MediaDirListRequest
	req.Command = "mediaDirList"

	err := fetch(host, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

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

type CommandListRequest struct {
	Command string `json:"command"`
	Path string `json:"path,omitempty"`
}
type CommandListResponse struct {
	CommandList []CommandList_CommandObject
	Result int
}
type CommandList_CommandObject struct {
	Command string
	Version int
}

// CommandList lists all commands that the camera supports.  It does
// this by sending a commandList command request to the camera and
// returns a CommandListResponse if successful.
func CommandList(host string) (*CommandListResponse, error) {
	var resp CommandListResponse
	var req CommandListRequest
	req.Command="commandList"

	err := fetch(host, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}


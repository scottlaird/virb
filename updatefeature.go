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

type UpdateFeatureRequest struct {
	Command string `json:"command"`
	Feature string `json:"feature"`
	Value   string `json:"value"`
}
type UpdateFeatureResponse struct {
	Features []UpdateFeatures_Object
	Result   int
}

type UpdateFeatures_Object struct {
	Enabled         int
	Feature         string
	Options         []string
	OptionSummaries []string
	Type            int
	Value           string
}

func UpdateFeature(host string, feature string, value string) (*UpdateFeatureResponse, error) {
	var resp UpdateFeatureResponse
	var req UpdateFeatureRequest
	req.Command = "updateFeature"
	req.Feature = feature
	req.Value = value

	err := fetch(host, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

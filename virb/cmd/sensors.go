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

package cmd

import (
	"fmt"

	"github.com/scottlaird/virb"
	"github.com/spf13/cobra"
)

// sensorsCmd represents the sensors command
var sensorsCmd = &cobra.Command{
	Use:   "sensors",
	Short: "List all sensors on the camera",
	Long:  `sensors [<sensorname>]`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := virb.Sensors(hostname, args)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", *resp)
	},
}

func init() {
	rootCmd.AddCommand(sensorsCmd)
}

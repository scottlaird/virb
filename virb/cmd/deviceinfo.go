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

	"github.com/spf13/cobra"
	"github.com/scottlaird/virb"
)

// deviceinfoCmd represents the deviceinfo command
var deviceinfoCmd = &cobra.Command{
	Use:   "deviceinfo",
	Short: "Fetch device info from a Garmin Virb camera",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := virb.DeviceInfo(hostname)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", *resp)
	},
}

func init() {
	rootCmd.AddCommand(deviceinfoCmd)
}

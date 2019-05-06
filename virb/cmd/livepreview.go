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

var streamType string
var maxResolution int
var active bool

// livepreviewCmd represents the livepreview command
var livepreviewCmd = &cobra.Command{
	Use:   "livepreview",
	Short: "Start a live preview stream from Garmin Virb camera",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		act := "0"
		if active {
			act = "1"
		}
		resp, err := virb.LivePreview(hostname, streamType, maxResolution, act)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", *resp)
	},
}

func init() {
	rootCmd.AddCommand(livepreviewCmd)
	livepreviewCmd.Flags().StringVar(&streamType, "stream_type", "rtp", "Stream type; should always be rtp")
	livepreviewCmd.Flags().IntVar(&maxResolution, "max_resolution", 0, "Maximum vertical resolution")
	livepreviewCmd.Flags().BoolVar(&active, "active", true, "Make stream active; --active=false to stop")
}

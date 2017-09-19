// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/twhiston/pulsetrainer/api"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		//get config file location
		confPath, err := cmd.PersistentFlags().GetString("config")
		HandleError(err)
		auto, err := cmd.PersistentFlags().GetBool("auto")
		HandleError(err)

		if confPath == "" {
			//TODO - look for conf in the cwd
			log.Fatal("Not implemented yet")
		}

		//Load and parse config file
		cfg := new(api.Config)
		yamlFile, err := ioutil.ReadFile(confPath)
		HandleError(err)
		err = yaml.Unmarshal(yamlFile, cfg)
		HandleError(err)

		//Run the interactive console so users can interact with the process
		if auto != true {
			//Runloop does its own error handling and will exit if errors are found
			api.RunLoop(cfg)
		}

		//TODO - If not in auto then run the sequence automatically
		log.Fatal("should not get here, auto is not implemented yet")
		//L := lua.NewState()
		//defer L.Close()
		//if err := L.DoString(`print("hello")`); err != nil {
		//	panic(err)
		//}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().String("config", "", "A config file to use for the current execution, if not set looks in the current working directory")
	runCmd.PersistentFlags().String("routine", "", "Specify a routine name from the config file to have it preloaded into the interactive console ready for running")
	runCmd.PersistentFlags().Bool("auto", false, "If true then interactive console will be skipped and routine will be run automatically")

}

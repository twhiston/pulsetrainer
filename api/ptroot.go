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

package api

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var engine *DefaultEngine

func init() {
	engine = New()
}

// RootCmd represents the base command when called without any subcommands
var PtRoot = &cobra.Command{
	Use:   "pulsetrainer",
	Short: "",
	Long:  ``,
}

var ptLoad = &cobra.Command{
	Use:   "load",
	Short: "loads a routine",
	Long: `This command will load a routine from the specified configuration file.
	This config file must be in the cwd or passed to the run command using the --config flag.
	The config file cannot be changed while in the interactive console`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			log.Println("no routine specified")
			return
		}
		//cmd2.HandleError(engine.SetActiveRoutine(args[0]))
		engine.SetActiveRoutine(args[0])
		log.Println("loaded config:", args[0])
	},
}

var ptStart = &cobra.Command{
	Use:   "run",
	Short: "run a routine",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		engine.Run()
		//cmd2.HandleError(engine.Run())
	},
}

var ptStop = &cobra.Command{
	Use:   "stop",
	Short: "stops a routine",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		engine.Stop()
	},
}

//var ptReset = &cobra.Command{
//	Use:   "reset",
//	Short: "resets a routine's active step to 0",
//	Long:  ``,
//	Run: func(cmd *cobra.Command, args []string) {
//		if ActiveState.RoutineId == "" {
//			log.Println("No routine set, use the load command first")
//			return
//		}
//
//	},
//}
//
//var ptAdv = &cobra.Command{
//	Use:   "adv",
//	Short: "advances a routine by 1 step",
//	Long:  `The generator can be optionally specified in the commands`,
//	Run: func(cmd *cobra.Command, args []string) {
//		if ActiveState.RoutineId == "" {
//			log.Println("No routine set, use the load command first")
//			return
//		}
//
//	},
//}
//
var ptExit = &cobra.Command{
	Use:   "exit",
	Short: "exits the console",
	Long:  `Running this command will end the pulsetrainer interactive console`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(1)
	},
}

func init() {
	PtRoot.AddCommand(ptStart)
	PtRoot.AddCommand(ptStop)
	PtRoot.AddCommand(ptLoad)
	PtRoot.AddCommand(ptExit)
	//PtRoot.AddCommand(ptConf)
	//PtRoot.AddCommand(ptState)
	//ptState.AddCommand(ptState_id)
	//ptState.AddCommand(ptState_routine)
}

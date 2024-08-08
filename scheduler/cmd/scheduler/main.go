/*
   Copyright 2023 The Kubernetes Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/logs"
	_ "k8s.io/component-base/logs/json/register" // for JSON log format registration
	_ "k8s.io/component-base/metrics/prometheus/clientgo"
	_ "k8s.io/component-base/metrics/prometheus/version" // for version metric registration
	"k8s.io/component-base/version/verflag"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"k8s.io/kubernetes/cmd/kube-scheduler/app/options"

	wasm "sigs.k8s.io/kube-scheduler-wasm-extension/scheduler/plugin"
)

var rootCmd = cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		configFile := cmd.Flag("config").Value.String()
		pluginNames, err := getWasmPluginsFromConfig(configFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get wasm plugins from config: %v\n", err)
			os.Exit(1)
		}

		opt := []app.Option{}
		for _, pluginName := range pluginNames {
			opt = append(opt, app.WithPlugin(pluginName, wasm.PluginFactory(pluginName)))
		}
		command := app.NewSchedulerCommand(opt...)
		command.Execute()
	},
	SilenceUsage: true,
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	opts := options.NewOptions()
	nfs := opts.Flags
	verflag.AddFlags(nfs.FlagSet("global"))
	globalflag.AddGlobalFlags(nfs.FlagSet("global"), rootCmd.Name(), logs.SkipLoggingConfigurationFlags())
	fs := rootCmd.Flags()
	for _, f := range nfs.FlagSets {
		fs.AddFlagSet(f)
	}
}

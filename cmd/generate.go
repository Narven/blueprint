/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"io/ioutil"
	"log"

	"github.com/Narven/blueprint/internal/app"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a new file/folder structure",
	Long:  `Generates a new file/folder structure based on a template`,
	Run: func(cmd *cobra.Command, args []string) {
		templateFlag, err := cmd.Flags().GetString("template")
		if err != nil {
			log.Fatal(err)
		}

		content, err := ioutil.ReadFile(templateFlag)
		if err != nil {
			log.Fatal(err)
		}

		app.Parse(content)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("template", "t", "", "Path to the config (yml) you want to use")
}

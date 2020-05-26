/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "create-shell-script",
	Short: "create-shell-script",
	Long:  `Create a shell script`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("[START]")
		// toggle, _ := cmd.Flags().GetBool("toggle")
		// fmt.Println(toggle)
		dir, _ := cmd.Flags().GetString("dir")
		// fmt.Println(dir)
		fileName, _ := cmd.Flags().GetString("file")
		// fmt.Println(fileName)
		filePath := path.Join(dir, fileName)
		// fmt.Println(filePath)

		// Create the script file
		if _, err := os.Stat(filePath); err == nil {
			fmt.Println("File exists:", filePath)
			os.Exit(0)
		} else {
			// fmt.Println("File does not exists:", filePath)
			// fmt.Println("Creating file:", filePath)
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			// fmt.Println("Writing")
			// content := "hi2"
			// fmt.Println("Content:")
			// fmt.Println(content)
			// _, err = file.WriteString(content)
			// fmt.Println("Finished writing")
			file.Chmod(0755)
			fmt.Println("File created at:", filePath)
		}
		// fmt.Println("[END]")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.create-shell-script.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().String("dir", "", "The script direcotry")
	rootCmd.Flags().String("file", "", "The script file name")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".create-shell-script" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".create-shell-script")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

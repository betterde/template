/*
Copyright © 2026 George <george@betterde.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/betterde/template/mcp/config"
	"github.com/betterde/template/mcp/internal/build"
	"github.com/betterde/template/mcp/internal/journal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	verbose bool
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     build.Name,
	Short:   build.Desc,
	Version: fmt.Sprintf("Version: %s\nBuild at: %s\nCommit hash: %s", build.Version, build.Build, build.Commit),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .mcp-server-mysql.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")

	// Print app version
	rootCmd.SetVersionTemplate("{{ .Version }}\n")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	journal.InitLogger()

	// Parse config from file and env variables
	config.Parse(cfgFile)

	level := viper.GetString("logging.level")
	if verbose {
		level = "DEBUG"
	}

	err := journal.SetLevel(level)
	if err != nil {
		journal.Logger.Sugar().Error("Unable to set logger level", err)
		os.Exit(1)
	}

	journal.Logger.Sugar().Debugf("Configuration file currently in use: %s", viper.ConfigFileUsed())
}

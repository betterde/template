/*
Copyright © 2024 George <george@betterde.com>

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
	"github.com/betterde/template/fiber/internal/journal"
	"github.com/betterde/template/fiber/pkg/api"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
	"reflect"
	"sort"
	"strings"
)

const MethodColWidth = 8

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "List all registered routes",
	Run: func(cmd *cobra.Command, args []string) {
		routes := api.ServerInstance.Engine.GetRoutes(true)
		sort.Slice(routes, func(i, j int) bool {
			iV := reflect.ValueOf(routes[i])
			jV := reflect.ValueOf(routes[j])
			iPosField := iV.FieldByName("pos")
			jPosField := jV.FieldByName("pos")
			return iPosField.Uint() < jPosField.Uint()
		})

		termWidth, _, err := term.GetSize(int(os.Stderr.Fd()))
		if err != nil {
			journal.Logger.Fatal(err)
			os.Exit(1)
		}

		rows := make([][]string, 0, len(routes))

		for _, route := range routes {
			if route.Method == "HEAD" {
				continue
			} else if route.Method == "GET" {
				route.Method = "GET|HEAD"
			}

			rows = append(rows, []string{route.Method, format(route.Path, route.Name, termWidth-MethodColWidth, ".")})
		}

		t := table.New().Border(lipgloss.HiddenBorder()).StyleFunc(func(row, col int) lipgloss.Style {
			if col == 0 {
				return lipgloss.NewStyle().Width(8)
			}
			return lipgloss.NewStyle().Padding(0, 1)
		}).Rows(rows...).Width(termWidth)

		fmt.Println(t.Render())
	},
}

func init() {
	rootCmd.AddCommand(routeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func format(path, name string, terminalWidth int, paddingChar string) string {
	// Calculate lengths
	pathLen := len(path)
	nameLen := len(name)

	// Calculate available space for padding
	// Account for spaces, methods, and fixed spacing
	fixedSpacing := 7 // 2 spaces before methods, 3 spaces after methods, 2 spaces before name
	maxPadding := terminalWidth - pathLen - nameLen - fixedSpacing

	if maxPadding < 0 {
		maxPadding = 0 // Prevent negative padding
	}

	// Build the formatted string
	var builder strings.Builder
	builder.WriteString("  ") // Space after methods
	builder.WriteString(path)
	builder.WriteString(" ") // Space after path

	// Add padding
	builder.WriteString(strings.Repeat(paddingChar, maxPadding))

	// Add name if it exists
	if name != "" {
		builder.WriteString(" ") // Space before name
		builder.WriteString(name)
	} else {
		builder.WriteString(paddingChar)
	}

	return builder.String()
}

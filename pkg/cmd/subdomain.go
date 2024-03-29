/*
Copyright © 2020 Bjarn Bronsveld <bjarn@bronsveld.me>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subdomainCmd represents the subdomain command
var subdomainCmd = &cobra.Command{
	Use:   "subdomain",
	Short: "Add or remove a subdomain for the current project",
	Long: `Add or remove a subdomain for the current project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subdomain called")
	},
}

func init() {
	rootCmd.AddCommand(subdomainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subdomainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subdomainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

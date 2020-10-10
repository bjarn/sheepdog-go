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
	"github.com/bjarn/sheepdog/pkg/install"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Run the initial setup of Sheepdog",
	Long: `Run the initial setup of Sheepdog.`,
	Run: func(cmd *cobra.Command, args []string) {
		install.Run()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

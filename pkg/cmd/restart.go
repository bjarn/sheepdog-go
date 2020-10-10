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
	"github.com/bjarn/sheepdog/pkg/service"

	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart [service]",
	Short: "Restart a service",
	Long:  `Restart a service`,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			service.RestartAll()
			return
		}

		switch args[0] {
		case service.Nginx.Name, service.MySql57.Name, service.MySql80.Name,
			service.Redis.Name, service.Mailhog.Name, service.DnsMasq.Name:
			for _, s := range service.Services {
				if s.Name == args[0] {
					service.RestartSingle(s)
				}
			}
			break
		default:
			fmt.Printf("Service is invalid.\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

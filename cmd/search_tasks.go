package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: `Search text from tasks`,
	Long: `Search text from tasks. example:
	todocli search shopping
	todocli search shopping,car
	todocli search "some text"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var query string

		if len(args) == 0 {
			fmt.Println("no search text given")
			return
		}
		parts := strings.Split(args[0], ",")
		query = "search=" + strings.Join(parts, ",")

		tasks, err := getAllTasks(query)
		if err != nil {
			fmt.Println(fmt.Sprintf("getting tasks failed: %v", err))
			return
		}
		printTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

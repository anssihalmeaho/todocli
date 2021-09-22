package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getWithIDCmd represents the get command
var getWithIDCmd = &cobra.Command{
	Use:   "id <id>",
	Short: `Get task with ID`,
	Long: `Get task with ID. example:
	todocli get id 12
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("task ID not given")
			return
		}
		taskID := args[0]
		tasks, err := getTaskWithID(taskID)
		if err != nil {
			fmt.Println(fmt.Sprintf("getting tasks failed: %v", err))
			return
		}
		printTasks(tasks)
	},
}

func init() {
	getCmd.AddCommand(getWithIDCmd)
}

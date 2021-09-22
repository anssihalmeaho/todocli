package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updCmd represents the upd command
var updCmd = &cobra.Command{
	Use:   "upd <task-ID>",
	Short: `Update task`,
	Long: `Update task:
	todocli upd 15 --desc="Wash the car"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var task Task

		if len(args) == 0 {
			fmt.Println("No task id given")
			return
		}
		taskID := args[0]

		fillTaskStrField(cmd, &task.Name, "name")
		fillTaskStrField(cmd, &task.Descr, "desc")
		fillTaskStrField(cmd, &task.State, "state")
		fillTaskStrListField(cmd, &task.Tags, "tags")

		if err := updTask(task, taskID); err == nil {
			fmt.Println("updated")
		} else {
			fmt.Println(fmt.Sprintf("Error: %v", err))
		}
	},
}

func init() {
	updCmd.Flags().StringP("name", "n", "", "Task name")
	updCmd.Flags().StringP("state", "s", "", "Task state")
	updCmd.Flags().StringP("tags", "t", "", "Tags belonging to task")
	updCmd.Flags().StringP("desc", "d", "", "Task description")
	rootCmd.AddCommand(updCmd)
}

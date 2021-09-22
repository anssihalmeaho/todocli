package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del <task-ID>",
	Short: `Delete task`,
	Long: `Delete task:
	todocli del 15
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No task id given")
			return
		}
		taskID := args[0]
		if err := delTask(taskID); err == nil {
			fmt.Println("deleted")
		} else {
			fmt.Println(fmt.Sprintf("Error: %v", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}

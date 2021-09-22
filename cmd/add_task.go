package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func fillTaskStrField(cmd *cobra.Command, str *string, optName string) {
	option, err := cmd.Flags().GetString(optName)
	if err != nil {
		fmt.Println(fmt.Sprintf("getting flag (%s) failed: %v", optName, err))
		return
	}
	if option != "" {
		*str = option
	}
}

func fillTaskStrListField(cmd *cobra.Command, strList *[]string, optName string) {
	option, err := cmd.Flags().GetString(optName)
	if err != nil {
		fmt.Println(fmt.Sprintf("getting flag (%s) failed: %v", optName, err))
		return
	}
	if option != "" {
		*strList = strings.Split(option, ",")
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add ",
	Short: `Add task`,
	Long: `Add task:
	todocli add --name="Clean house" -tags=house,cleaning
	todocli add --name=carwashing --desc="Wash the car"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var task Task

		fillTaskStrField(cmd, &task.Name, "name")
		fillTaskStrField(cmd, &task.Descr, "desc")
		fillTaskStrField(cmd, &task.State, "state")
		fillTaskStrListField(cmd, &task.Tags, "tags")

		if err := addTask(task); err == nil {
			fmt.Println("added")
		} else {
			fmt.Println(fmt.Sprintf("Error: %v", err))
		}
	},
}

func init() {
	addCmd.Flags().StringP("name", "n", "", "Task name")
	addCmd.Flags().StringP("state", "s", "", "Task state")
	addCmd.Flags().StringP("tags", "t", "", "Tags belonging to task")
	addCmd.Flags().StringP("desc", "d", "", "Task description")
	rootCmd.AddCommand(addCmd)
}

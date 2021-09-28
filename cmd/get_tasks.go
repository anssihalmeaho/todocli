package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func formatTask(task Task) string {
	s := fmt.Sprintf("\n Name: %s\n ID: %d\n Description: %s\n State: %s\n Tags: %s\n",
		task.Name,
		task.TaskID,
		task.Descr,
		task.State,
		strings.Join(task.Tags, ","),
	)
	return s
}

func printTasks(tasks []Task) {
	taskStr := "\nTasks:\n"
	for _, task := range tasks {
		taskStr += formatTask(task)
	}
	fmt.Println(taskStr)
}

func addToQuery(prevQ, addition string) string {
	if addition == "" {
		return prevQ
	}
	if prevQ == "" {
		return addition
	}
	return prevQ + "&" + addition
}

func makeQuery(cmd *cobra.Command, qname, query string) (string, bool) {
	option, err := cmd.Flags().GetString(qname)
	if err != nil {
		fmt.Println(fmt.Sprintf("getting flag (%s) failed: %v", qname, err))
		return "", false
	}
	if option == "" {
		return query, true
	}
	parts := strings.Split(option, ",")
	return addToQuery(query, qname+"="+strings.Join(parts, ",")), true
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: `Get tasks`,
	Long: `Get tasks. example:
	todocli get 
	todocli get --name=cooking --tags=home,kitchen
	todocli get --name=cleaning --state=new,done
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var query string
		var queryOK bool

		if query, queryOK = makeQuery(cmd, "name", query); !queryOK {
			return
		}
		if query, queryOK = makeQuery(cmd, "state", query); !queryOK {
			return
		}
		if query, queryOK = makeQuery(cmd, "tags", query); !queryOK {
			return
		}

		tasks, err := getAllTasks(query)
		if err != nil {
			fmt.Println(fmt.Sprintf("getting tasks failed: %v", err))
			return
		}
		printTasks(tasks)
	},
}

func init() {
	getCmd.Flags().StringP("name", "n", "", "Task name")
	getCmd.Flags().StringP("state", "s", "", "Task state")
	getCmd.Flags().StringP("tags", "t", "", "Tags belonging to task")
	rootCmd.AddCommand(getCmd)
}

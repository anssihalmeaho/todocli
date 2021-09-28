package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func printTags(tags []string) {
	taskStr := "\nTags:\n"
	for _, tag := range tags {
		taskStr += fmt.Sprintf("\n %s", tag)
	}
	fmt.Println(taskStr)
}

// getTagsCmd represents the get tags command
var getTagsCmd = &cobra.Command{
	Use:   "tags",
	Short: `Get all tags`,
	Long: `Get all tags. example:
	todocli get tags
	`,
	Run: func(cmd *cobra.Command, args []string) {
		tags, err := getTags()
		if err != nil {
			fmt.Println(fmt.Sprintf("getting tags failed: %v", err))
			return
		}
		printTags(tags)
	},
}

func init() {
	getCmd.AddCommand(getTagsCmd)
}

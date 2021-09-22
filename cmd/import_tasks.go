package cmd

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import <file-name>",
	Short: `Import tasks`,
	Long: `Import tasks. example:
	todocli import myexport.json
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("file name not given")
			return
		}

		filename := args[0]

		var isZIPFile bool
		fparts := strings.Split(filename, ".")
		if fparts[len(fparts)-1] == "zip" {
			isZIPFile = true
		}

		var content []byte
		var err error

		if isZIPFile {
			reader, err := zip.OpenReader(filename)
			if err != nil {
				fmt.Println(fmt.Sprintf("reading tasks from ZIP (%s) failed: %v", filename, err))
				return
			}
			defer reader.Close()
			if len(reader.File) == 0 {
				fmt.Println(fmt.Sprintf("Empty ZIP (%s)", filename))
				return
			}
			rc, err := reader.File[0].Open()
			if err != nil {
				fmt.Println(fmt.Sprintf("File open in ZIP (%s) failed: %v", filename, err))
				return
			}

			buf := bytes.NewBuffer([]byte{})

			_, err = io.Copy(buf, rc)
			if err != nil {
				fmt.Println(fmt.Sprintf("Copy in ZIP (%s) failed: %v", filename, err))
				return
			}
			rc.Close()
			content = buf.Bytes()
		} else {
			content, err = os.ReadFile(filename)
			if err != nil {
				fmt.Println(fmt.Sprintf("reading tasks from file (%s) failed: %v", filename, err))
				return
			}
		}

		//fmt.Println(string(content))
		err = importTasks(content)
		if err != nil {
			fmt.Println(fmt.Sprintf("importing tasks from file (%s) failed: %v", filename, err))
			return
		}
		fmt.Println(fmt.Sprintf("Imported from file: %s", filename))
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}

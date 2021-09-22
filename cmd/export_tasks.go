package cmd

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: `Export tasks`,
	Long: `Export tasks. example:
	todocli export 
	todocli export --file=myexport.json
	`,
	Run: func(cmd *cobra.Command, args []string) {
		content, err := importAllTasks("")
		if err != nil {
			fmt.Println(fmt.Sprintf("getting tasks failed: %v", err))
			return
		}

		doZIPFile, err := cmd.Flags().GetBool("zip")
		if err != nil {
			fmt.Println(fmt.Sprintf("getting flag (%s) failed: %v", "zip", err))
			return
		}

		var filename string
		targetfile, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(fmt.Sprintf("getting flag (%s) failed: %v", "file", err))
			return
		}
		if targetfile == "" {
			sec := time.Now().Unix()
			filename = fmt.Sprintf("task_export_%v.json", sec)
		} else {
			filename = targetfile
		}

		if doZIPFile {
			fnparts := strings.Split(filename, ".")
			if len(fnparts) == 0 {
				fmt.Println(fmt.Sprintf("Invalid file name (%s)", filename))
				return
			}
			fname := fnparts[0]
			//fmt.Println(fmt.Sprintf("fname = <%s>", fname))

			archive, err := os.Create(fname + ".zip")
			if err != nil {
				panic(err)
			}
			zipWriter := zip.NewWriter(archive)

			w1, err := zipWriter.Create(filename)
			if err != nil {
				fmt.Println(fmt.Sprintf("Creating file in ZIP failed: %v", err))
				return
			}
			buf := bytes.NewBuffer(content)
			if _, err := io.Copy(w1, buf); err != nil {
				fmt.Println(fmt.Sprintf("Writing file to ZIP failed: %v", err))
				return
			}
			zipWriter.Close()

			filename = fname + ".zip"
		} else {
			//fmt.Println(fmt.Sprintf("<%s>", filename))
			err = os.WriteFile(filename, content, 0666)
			if err != nil {
				fmt.Println(fmt.Sprintf("writing tasks to file failed: %v", err))
				return
			}
		}

		fmt.Println(fmt.Sprintf("Exported to file: %s", filename))
	},
}

func init() {
	exportCmd.Flags().StringP("file", "f", "", "File name")
	exportCmd.Flags().Bool("zip", false, "Make ZIP file")
	rootCmd.AddCommand(exportCmd)
}

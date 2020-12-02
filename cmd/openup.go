
package cmd

import (
	"fmt"
	"os"
	"bufio"

	"github.com/spf13/cobra"
	"github.com/pkg/browser"
)

// openupCmd represents the openup command
var openupCmd = &cobra.Command{
	Use:   "openup",
	Short: "opens up urls",
	Long:"A command to open my usual tabs and start my working day",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Rise and Shine, Start your day")

		// Declare variable 
		pathToFile := args[0]

		// Shows the path to the file and declares an error if any

		file, err := os.Open(pathToFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} 
		
		// Scan the file and opens up in browser
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			browser.OpenURL(line)
		}

	},
}

func init() {
	rootCmd.AddCommand(openupCmd)


}

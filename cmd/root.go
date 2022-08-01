package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var GetTodoListFromFile func(string) (map[string]bool, error)

func init() {
}

// rootCmd represents the base command when called without any subcommands

// get readme.md text
func getReadme() string {
	readme, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	return string(readme)
}

var rootCmd = &cobra.Command{
	Use:   "dayml",
	Short: "A simple cli tool for parsing TODO information from notes written in YAML.",
	// get Long from README.md
	Long: getReadme(),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	filePath := rootCmd.Flag("file").Value.String()
	log.Print(filePath)
	todoList, _ := GetTodoListFromFile(filePath)

	fmt.Print("The following tasks are incomplete: \n")
	for k, v := range todoList {
		if !v {
			fmt.Print(k)
		}
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yml.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("file", "f", "", "Help message for toggle")
}

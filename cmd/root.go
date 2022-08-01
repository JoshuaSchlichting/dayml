package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var GetTodoListFromFile func(string) (map[string]bool, error)

func init() {
}

// rootCmd represents the base command when called without any subcommands

var rootCmd = &cobra.Command{
	Use:   "dayml",
	Short: "dayml - A simple cli tool for parsing TODO information from notes written in YAML.",

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
	if filePath == "" {
		filePath = ".dayml.yml"

		fmt.Printf("No file path provided. Checking for %s in current working directory....\n", filePath)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("No file found at %s\n", filePath)
		}

		fmt.Printf("Please specify a file path via \"--file\" or create a .dayml.yml file in the current working directory.\n")
	}

	todoList, _ := GetTodoListFromFile(filePath)
	if todoList == nil {
		log.Fatal("Error parsing file")
	}
	if len(todoList) == 0 {
		log.Fatal("No TODO items found")
	}
	fmt.Print("The following tasks are incomplete: \n")
	for k, v := range todoList {
		if !v {
			fmt.Printf("%s\n", k)
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
	rootCmd.Flags().StringP("file", "f", "", "Filepath to dayml compliant YAML file (.yml, .yaml)")
}

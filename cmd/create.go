/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todo-go/db"

	"github.com/spf13/cobra"
)

var (
	title       string
	description string
	completed   bool
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")

		dbHandler, err := db.Connect()
		if err != nil {
			panic("Could not connect to the database")
		}

		err = dbHandler.AutoMigrate(&db.Todo{})
		if err != nil {
			panic("Could not migrate the database")
		}

		dbHandler.Create(&db.Todo{Title: title, Description: description, Completed: completed})

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringVar(&title, "title", "", "Title of the todo")
	createCmd.Flags().StringVar(&description, "description", "", "Description of the todo")
	createCmd.Flags().BoolVar(&completed, "completed", false, "Status of the todo")

	createCmd.MarkFlagRequired("title")
}

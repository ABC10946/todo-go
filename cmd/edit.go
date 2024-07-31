/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todo-go/db"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")

		dbHandler, err := db.Connect()
		if err != nil {
			panic("Could not connect to the database")
		}

		existTodo := db.Todo{}
		dbHandler.First(&existTodo, id)
		if title != "" {
			existTodo.Title = title
		}
		if description != "" {
			existTodo.Description = description
		}

		existTodo.Completed = completed
		result := dbHandler.Save(&existTodo)
		if result.Error != nil {
			panic("Could not update the todo")
		}

		fmt.Println("Todo updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	editCmd.Flags().StringVar(&title, "title", "", "Title of the todo")
	editCmd.Flags().StringVar(&description, "description", "", "Description of the todo")
	editCmd.Flags().BoolVar(&completed, "completed", false, "Status of the todo")

	editCmd.Flags().IntVar(&id, "id", 0, "ID of the todo")
	editCmd.MarkFlagRequired("id")

}

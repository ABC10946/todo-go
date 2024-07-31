/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"todo-go/cmd"
	"todo-go/db"
)

func main() {
	_, err := db.Connect()
	if err != nil {
		panic("Could not connect to the database")
	}

	cmd.Execute()
}

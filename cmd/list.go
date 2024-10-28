/*
Copyright © 2024 Saman Dehestani <saman8dehestani@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/drippypale/todo-list/config"
	"github.com/drippypale/todo-list/csvHandler"
	"github.com/drippypale/todo-list/model"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all todo tasks",
	Long:  `List of all todo tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		csvPath := viper.GetString("csvPath")

		todoList := model.TodoFromRecords(csvHandler.ReadRecords(csvPath))

		w := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
		defer w.Flush()
		fmt.Fprintln(w, "#\tTask Name\tDue Time\tPriority")

		for _, todo := range todoList {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", todo.Id, todo.Task, todo.DueTime.Format("2006-01-02 15:04"), todo.Priority)
		}

	},
}

func init() {
	config.InitConfig()
	rootCmd.AddCommand(listCmd)
}

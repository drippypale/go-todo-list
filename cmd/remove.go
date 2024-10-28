/*
Copyright © 2024 Saman Dehestani <saman8dehestani@gmail.com>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/drippypale/todo-list/config"
	"github.com/drippypale/todo-list/csvHandler"
	"github.com/drippypale/todo-list/model"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove todo items by list of IDs",
	Long: `Remove todo items by list of IDs

  e.g. todo-list remove 2 3 25 24
  `,
	Run: func(cmd *cobra.Command, args []string) {
		csvPath := viper.GetString("csvPath")

		removeIdList := make([]int, len(args))
		for i, arg := range args {
			id, _ := strconv.Atoi(arg)
			removeIdList[i] = id
		}
		todoList := model.TodoFromRecords(csvHandler.ReadRecords(csvPath))
		todoListMap := make(map[int]model.Todo, len(todoList)-len(removeIdList))

		for _, todo := range todoList {
			todoListMap[todo.Id] = todo
		}

		removeList := make([]model.Todo, len(removeIdList))
		for i, id := range removeIdList {
			removeList[i] = todoListMap[id]
			delete(todoListMap, id)
		}
		todoList = make([]model.Todo, 0)
		for _, todo := range todoListMap {
			todoList = append(todoList, todo)
		}

		fw, err := os.OpenFile(csvPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		defer fw.Close()

		if err != nil {
			log.Fatalf("Can not open the csv file to write: %v\n%s", csvPath, err)
		}

		csvW := csv.NewWriter(fw)

		csvW.WriteAll(model.TodoToRecords(todoList))

		tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
		defer tw.Flush()

		fmt.Printf("Successfully deleted %v tasks:\n\n", len(removeIdList))
		fmt.Fprintln(tw, "#\tTask Name\tDue Time\tPriority")
		for _, todo := range removeList {
			fmt.Fprintf(
				tw,
				"%v\t%v\t%v\t%v\n",
				todo.Id,
				todo.Task,
				todo.DueTime.Format(
					viper.GetString("timeFormat"),
				),
				todo.Priority)
		}
	},
	Args: cobra.MinimumNArgs(1),
}

func init() {
	config.InitConfig()
	rootCmd.AddCommand(removeCmd)
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/drippypale/todo-list/config"
	"github.com/drippypale/todo-list/csvHandler"
	"github.com/drippypale/todo-list/model"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:       "add",
	Short:     "Add a new todo task to your schedule",
	Long:      `Add a new todo task to your schedule`,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"task name"},
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		priority, _ := cmd.Flags().GetInt("priority")
		tomorrow, _ := cmd.Flags().GetBool("tomorrow")
		datetime, _ := cmd.Flags().GetString("datetime")

		if datetime != "" && tomorrow {
			log.Fatalln("You can not specify both datetime and tomorrow args")
		}

		var dueTime time.Time

		if datetime != "" {
			dueTime, _ = time.Parse("2006-01-02 15:04", datetime)
		} else if tomorrow {
			// Get yesterday date
			y, m, d := time.Now().AddDate(0, 0, 1).Date()
			loc, _ := time.LoadLocation("Local")

			dueTime = time.Date(y, m, d, 9, 0, 0, 0, loc)
		}
		todoList := model.TodoFromRecords(csvHandler.ReadRecords(viper.GetString("csvPath")))

		todo := model.Todo{Task: task, DueTime: dueTime, Priority: model.Priority(priority)}

		addTodo(&todoList, &todo)

		w := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
		defer w.Flush()
		fmt.Printf("\nTodo #%v successfully added:\n\n", todo.Id)
		fmt.Fprintln(w, "#\tTask Name\tDue Time\tPriority")
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", todo.Id, todo.Task, todo.DueTime.Format("2006-01-02 15:04"), todo.Priority)
	},
}

func init() {
	config.InitConfig()

	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntP("priority", "P", 0, "Task priority ranging from 0-2 (2 for the lowest)")
	addCmd.Flags().StringP("datetime", "D", "", "Todo due time in the format yyyy-mm-dd HH:MM")
	addCmd.Flags().BoolP("tomorrow", "T", false, "If set, the task will set to tomorrow morning")
}

func getNewId(todoList []model.Todo) int {
	if len(todoList) == 0 {
		return 1
	}
	newId := todoList[len(todoList)-1].Id + 1
	return newId
}

func addTodo(todoList *[]model.Todo, todo *model.Todo) {
	csvPath := viper.GetString("csvPath")

	fw, err := os.OpenFile(csvPath, os.O_APPEND|os.O_WRONLY, 0666)
	defer fw.Close()

	if err != nil {
		panic("Can not write to the CSV file.")
	}

	csvW := csv.NewWriter(fw)
	defer csvW.Flush()

	todo.Id = getNewId(*todoList)
	*todoList = append(*todoList, *todo)

	csvW.Write(todo.ToRecord())
}

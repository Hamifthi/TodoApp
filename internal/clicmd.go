package internal

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
	"todoapp/db"
)

var dbService = db.DataBaseService{
	DB: nil,
}

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI for managing your TODOs.",
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Use add command to add to the todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		var database, _ = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		defer database.Close()
		dbService.DB = database
		// TODO here add, add command logic which is adding task to the database.
		var task = strings.Join(args[:], " ")
		_, err := dbService.AddTask(task)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Use list command to fetch tasks from todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		var database, _ = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		defer database.Close()
		dbService.DB = database
		// TODO here add list command logic which is fetching all tasks from database.
		taskList, err := dbService.ListTasks()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You have the following tasks:")
		for i := 0; i < len(taskList); i++ {
			fmt.Printf("%d. %s\n", i, taskList[i].Value)
		}
	},
}

var RmCmd = &cobra.Command{
	Use:   "remove",
	Short: "Use rm command to remove task from todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		var database, _ = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		defer database.Close()
		dbService.DB = database
		// TODO here add rm command logic which is removing task from database.
		var key = strings.Join(args[:], " ")
		tasks, err := dbService.ListTasks()
		if err != nil {
			log.Fatal(err)
		}
		primaryKey, _ := strconv.Atoi(key)
		if primaryKey < 0 || primaryKey > len(tasks) {
			fmt.Println("Invalid task number:", primaryKey)
		} else {
			task := tasks[primaryKey]
			value := dbService.GetTask(task.Key)
			err = dbService.RemoveTask(task.Key)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("You have deleted the \"%s\" task.\n", value)
		}
	},
}

var DoCmd = &cobra.Command{
	Use:   "do",
	Short: "Use do command to do task from todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		var database, _ = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		defer database.Close()
		dbService.DB = database
		// TODO here add do command logic which is doing task.
		var key = strings.Join(args[:], " ")
		primaryKey, _ := strconv.Atoi(key)
		tasks, err := dbService.ListTasks()
		if err != nil {
			log.Fatal(err)
		}
		if primaryKey < 0 || primaryKey > len(tasks) {
			fmt.Println("Invalid task number:", primaryKey)
		} else {
			task := tasks[primaryKey]
			key, err := dbService.DoTask(task.Key)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("You have completed the \"%s\" task.\n", key)
		}
	},
}

var CompletedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Use completed command to get completed tasks of last 24hours.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO here add compeleted command logic which fetching all last 24hours completed tasks.
		var database, _ = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		defer database.Close()
		dbService.DB = database
		completedTaskList, err := dbService.CompletedTasks()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You have completed the following tasks:")
		for i := 0; i < len(completedTaskList); i++ {
			fmt.Printf("%d. %s\n", i, completedTaskList[i].Value)
		}
	},
}

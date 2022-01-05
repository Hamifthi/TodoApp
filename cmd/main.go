package main

import (
	"todoapp/internal"
)

func main() {
	internal.RootCmd.AddCommand(internal.AddCmd,
		internal.DoCmd,
		internal.ListCmd,
		internal.RmCmd,
		internal.CompletedCmd)
	err := internal.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

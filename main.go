package main

import (
	"dayml/cmd"
	"dayml/core"
)

func main() {
	cmd.GetTodoListFromFile = core.GetTodoListFromFile
	cmd.Execute()
}

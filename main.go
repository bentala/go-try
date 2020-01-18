package main

import (
	"fmt"
	"os"
)

const (
	CommandExit string = "exit" // Exit Command
)

func main() {
	go os.Exit(start())
}

func start() int {
	var pos Position
	fmt.Println("[Starting]") // Info Starting Program
	_, _ = fmt.Scanf("%d %d %s", &pos.X, &pos.Y, &pos.Direction)
	fmt.Print("")
	var command string
	_, _ = fmt.Scanf("%s", &command)
	fmt.Print("")
	status := SelectCommand(command, &pos)
	for {
		if status == "exit" {
			fmt.Println("[Close]") // Info Closing Program
			return 1
		}
		fmt.Println(pos.X, pos.Y, pos.Direction)
		_, _ = fmt.Scanf("%s", &command)
		fmt.Print("")
		status = SelectCommand(command, &pos)
	}
}

func SelectCommand(command string, pos *Position) string {
	servicePosition, err := New(*pos)
	if err != nil {
		return "exit"
	}
	switch command {
	case CommandExit:
		return "exit"
	default:
		for iterationCommand := 0; iterationCommand < len(command); iterationCommand++ {
			switch string(command[iterationCommand]) {
			case "M":
				err = servicePosition.Forward(pos)
			case "L":
				err = servicePosition.Left(pos)
			case "R":
				err = servicePosition.Right(pos)
			default:
				err = fmt.Errorf("Wrong Format")
			}
			if err != nil {
				fmt.Println(err)
				return "exit"
			}
		}
	}
	return "continue"
}

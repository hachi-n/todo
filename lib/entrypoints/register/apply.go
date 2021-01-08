package register

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/hachi-n/todo/lib/models"
	"github.com/hachi-n/todo/lib/util"
	"os"
)

func Apply(markdownPath string) error {
	markdownMessage, err := loadInputMessage(markdownPath)
	if err != nil {
		return err
	}

	currentTodoList := models.LoadTodoList()
	newTodoList := models.NewTodoList(markdownMessage)

	if currentTodoList != nil {
		newTodoList.Merge(currentTodoList)
	}

	if err := newTodoList.Save(); err != nil {
		return err
	}

	fmt.Println(newTodoList)


	return nil
}

func loadInputMessage(markdownPath string) (message string, err error) {
	switch markdownPath {
	case "":
		return loadStdinMessage()
	default:
		return util.LoadFile(markdownPath)
	}
}

func loadStdinMessage() (message string, err error) {
	// TODO
	// While Reading bug.
	fmt.Println("Please enter your todolist. <paste> and <enter>")
	scanner := bufio.NewScanner(os.Stdin)
	buf := new(bytes.Buffer)

	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if line == "" {
			break
		}

		fmt.Fprintf(buf, "%s\n", line)
	}
	message = buf.String()
	return
}

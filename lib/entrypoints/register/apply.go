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

	return newTodoList.Save()
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
	scanner := bufio.NewScanner(os.Stdin)
	buf := new(bytes.Buffer)
	for scanner.Scan() {
		fmt.Fprintf(buf, "%s\n", scanner.Text())
	}
	message = buf.String()
	return
}

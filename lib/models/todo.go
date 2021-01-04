package models

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/hachi-n/todo/lib/util"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

type TodoList []*Todo

const (
	databaseFileName = "todo.md"
	databaseDirname  = "todo"
)

var (
	databaseFilePath = filepath.Join(os.Getenv("HOME"), databaseDirname, databaseFileName)
)

var reg = regexp.MustCompile(`##(.*)`)

func NewTodoList(markdownMessage string) TodoList {
	normalize := func(s string) string {
		return s
	}

	buf := bytes.NewBufferString(markdownMessage)
	scanner := bufio.NewScanner(buf)

	m := make(map[string][]string)
	var currentTask string
	for scanner.Scan() {
		t := scanner.Text()

		if currentTask != "" {
			m[currentTask] = append(m[currentTask], normalize(t))
		}

		// TODO
		//  FIX ME
		//  Bug Case
		//  "### task" -> "# task"
		if reg.MatchString(t) {
			tasks := reg.FindStringSubmatch(scanner.Text())
			if len(tasks) >= 2 {
				currentTask = strings.TrimSpace(tasks[1])
			}
		}
	}

	var todos []*Todo
	for k, v := range m {
		todos = append(todos, NewTodo(k, v))
	}

	return todos
}

func LoadTodoList() TodoList {
	_, err := os.Stat(databaseFilePath)
	if  err != nil{
		return nil
	}

	markdownMessage, err := util.LoadFile(databaseFilePath)
	if err != nil {
		fmt.Println("Crash Your todo database")
		os.Exit(1)
	}
	return NewTodoList(markdownMessage)
}

func (t TodoList) Merge(todoList TodoList) {
	t = append(t, todoList...)
}

func (t *TodoList) Save() error {
	// Convert TodoList To Markdown
	markdownMessage := t.convertMarkdown()
	return util.CreateFile(databaseFilePath, markdownMessage)
}

func (t *TodoList) String() string {
	return t.convertMarkdown()
}

func (t TodoList) convertMarkdown() string {
	var once sync.Once

	buf := new(bytes.Buffer)
	const headerMessage = "# Todo"
	for _, v := range t {
		once.Do(func() {
			fmt.Fprintf(buf, "%s\n", headerMessage)
		})
		fmt.Fprintf(buf, "## %s\n", v.title)
		for _, d := range v.description {
			fmt.Fprintf(buf, "- %s\n", d)
		}
	}

	return buf.String()
}

type Todo struct {
	title       string
	description []string
}

func NewTodo(title string, description []string) *Todo {
	return &Todo{
		title:       title,
		description: description,
	}
}

package models

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/hachi-n/todo/lib/util"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

type TodoList struct {
	todos []*Todo
}

type TodoListView struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func (v *TodoListView) Init() tea.Cmd {

	return nil
}

func (v *TodoListView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return v, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if v.cursor > 0 {
				v.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if v.cursor < len(v.choices)-1 {
				v.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := v.selected[v.cursor]
			if ok {
				delete(v.selected, v.cursor)
			} else {
				v.selected[v.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return v, nil
}

func (v TodoListView) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range v.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if v.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := v.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

const (
	databaseFileName = "todo.md"
	databaseDirname  = ".config/todo"
)

var (
	databaseDirPath  = filepath.Join(os.Getenv("HOME"), databaseDirname)
	databaseFilePath = filepath.Join(databaseDirPath, databaseFileName)
)

var reg = regexp.MustCompile(`##(.*)`)

func NewTodoList(markdownMessage string) *TodoList {
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

	todoList := new(TodoList)
	for k, v := range m {
		todoList.todos = append(todoList.todos, NewTodo(k, v))
	}

	return todoList
}

func LoadTodoList() *TodoList {
	_, err := os.Stat(databaseFilePath)
	if err != nil {
		return nil
	}

	markdownMessage, err := util.LoadFile(databaseFilePath)
	if err != nil {
		fmt.Println("Crash Your todo database")
		os.Exit(1)
	}
	return NewTodoList(markdownMessage)
}

func (t *TodoList) Merge(todoList *TodoList) {
	t.todos = append(t.todos, todoList.todos...)
}

func (t *TodoList) Save() error {
	// Convert TodoList To Markdown
	markdownMessage := t.convertMarkdown()
	if err := util.MkDirAll(databaseDirPath); err != nil {
		return err
	}
	return util.CreateFile(databaseFilePath, markdownMessage)
}

func (t *TodoList) String() string {
	return t.convertMarkdown()
}

func (t *TodoList) ViewModel() *TodoListView {

	viewModel := new(TodoListView)

	// collect task.
	for _, todo := range t.todos {
		viewModel.choices = append(viewModel.choices, todo.title)
	}

	// initializeMap
	viewModel.selected = make(map[int]struct{})

	return viewModel
}

func (t *TodoList) convertMarkdown() string {
	var once sync.Once

	buf := new(bytes.Buffer)
	const headerMessage = "# Todo"
	for _, v := range t.todos {
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

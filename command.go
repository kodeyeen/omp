package omp

import "fmt"

type Command struct {
	Sender   *Player
	Name     string
	Args     []string
	RawValue string
}

type CommandHandler func(cmd *Command)

type commandManager struct {
	commands map[string]CommandHandler
}

func newCommandManager() *commandManager {
	return &commandManager{
		commands: make(map[string]CommandHandler),
	}
}

func (m *commandManager) Add(name string, handler CommandHandler) {
	if m.Has(name) {
		panic(fmt.Sprintf("Command %s is already registered", name))
	}

	m.commands[name] = handler
}

func (m *commandManager) run(name string, cmd *Command) {
	handler := m.commands[name]
	handler(cmd)
}

func (m *commandManager) Has(name string) bool {
	_, ok := m.commands[name]
	return ok
}

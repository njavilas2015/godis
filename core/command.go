package core

import (
	"errors"
	"fmt"
	"strings"

	"github.com/njavilas2015/godis/storage"
)

var (
	kvStore     = storage.NewStorage()
	listStorage = storage.NewListStorage()
	setStore    = storage.NewSetStorage()
	hashStore   = storage.NewHashStorage()
	Dispatcher  = NewCommandDispatcher()
)

type CommandDispatcher struct {
	commands map[string]func(args []string) (string, error)
}

func NewCommandDispatcher() *CommandDispatcher {
	return &CommandDispatcher{
		commands: make(map[string]func(args []string) (string, error)),
	}
}

func (cd *CommandDispatcher) RegisterCommand(name string, handler func(args []string) (string, error)) {
	cd.commands[strings.ToLower(name)] = handler
}

func (cd *CommandDispatcher) Execute(command string, args []string) (string, error) {

	handler, exists := cd.commands[strings.ToLower(command)]

	if !exists {
		return "", errors.New("command not found")
	}

	return handler(args)
}

func RegisterCommand(operation string, args []string) {

	Dispatcher.RegisterCommand("SET", func(args []string) (string, error) {

		if len(args) != 2 {
			return "ERROR: Correct use: SET <key> <value>", nil
		}

		key, value := args[0], args[1]

		kvStore.Set(key, value)

		return "OK", nil
	})

	Dispatcher.RegisterCommand("GET", func(args []string) (string, error) {

		if len(args) != 1 {
			return "ERROR: Correct use: GET <key>", nil
		}

		key := args[0]

		value, exists := kvStore.Get(key)

		if !exists {
			return "(nil)", nil
		}

		return value, nil
	})

	Dispatcher.RegisterCommand("DEL", func(args []string) (string, error) {

		if len(args) != 1 {
			return "ERROR: Correct use: DEL <key>", nil
		}

		key := args[0]

		kvStore.Delete(key)

		return "OK", nil
	})

	Dispatcher.RegisterCommand("LPush", func(args []string) (string, error) {

		if len(args) < 2 {
			return "", errors.New("usage: LPush key value")
		}

		key, value := args[0], args[1]

		listStorage.LeftPush(key, value)

		return "OK", nil
	})

	Dispatcher.RegisterCommand("LIndex", func(args []string) (string, error) {

		if len(args) < 2 {
			return "", errors.New("usage: LIndex key index")
		}

		key := args[0]

		index := 0

		fmt.Sscanf(args[1], "%d", &index)

		result, err := listStorage.ListIndex(key, index)

		if err != nil {
			return "", err
		}

		return result, nil
	})

	Dispatcher.RegisterCommand("LRange", func(args []string) (string, error) {

		if len(args) < 3 {
			return "", errors.New("usage: LRange key start stop")
		}

		key := args[0]

		start, stop := 0, 0

		fmt.Sscanf(args[1], "%d", &start)

		fmt.Sscanf(args[2], "%d", &stop)

		result, err := listStorage.ListRange(key, start, stop)

		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%v", result), nil
	})
}

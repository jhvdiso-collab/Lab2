package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type stringTable interface {
	Put(string, string)
	Get(string) (string, bool)
	Delete(string) bool
}

// runHashConsole keeps the table alive until EXIT or end of input.
func runHashConsole(table stringTable, input io.Reader, output io.Writer) error {
	if _, err := fmt.Fprintln(output, "Команды: SET ключ значение | GET ключ | DEL ключ | EXIT\nКлюч и значение — без пробелов. Данные хранятся до выхода."); err != nil {
		return err
	}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		var response string
		switch command := strings.ToUpper(fields[0]); {
		case command == "EXIT" && len(fields) == 1:
			return nil
		case command == "SET" && len(fields) == 3:
			table.Put(fields[1], fields[2])
			response = "OK"
		case command == "GET" && len(fields) == 2:
			value, found := table.Get(fields[1])
			response = "Не найдено"
			if found {
				response = "Значение: " + value
			}
		case command == "DEL" && len(fields) == 2:
			response = "Не найдено"
			if table.Delete(fields[1]) {
				response = "Удалено"
			}
		default:
			response = "Ошибка: используйте SET ключ значение, GET ключ, DEL ключ или EXIT"
		}
		if _, err := fmt.Fprintln(output, response); err != nil {
			return err
		}
	}
	return scanner.Err()
}

package main

import (
	"encoding/json"
	"fmt"
	"lab2_go/internal/ds"
	"os"
	"sort"
	"strings"
)

func setQuery(path, query string) error {
	q := strings.Fields(query)
	if path == "" || len(q) != 2 {
		return fmt.Errorf("нужны --file путь --query 'SETADD/SETDEL/SET_AT элемент'")
	}
	if q[0] != "SETADD" && q[0] != "SETDEL" && q[0] != "SET_AT" {
		return fmt.Errorf("неизвестная операция")
	}
	var s ds.Set
	b, e := os.ReadFile(path)
	if e != nil && !os.IsNotExist(e) {
		return e
	}
	if e == nil {
		var values []string
		if e = json.Unmarshal(b, &values); e != nil {
			return e
		}
		for _, v := range values {
			s.Add(v)
		}
	}
	switch q[0] {
	case "SET_AT":
		fmt.Println(s.Contains(q[1]))
		return nil
	case "SETADD":
		s.Add(q[1])
	case "SETDEL":
		s.Delete(q[1])
	}
	values := s.Values()
	sort.Strings(values)
	b, e = json.MarshalIndent(values, "", "  ")
	if e != nil {
		return e
	}
	return os.WriteFile(path, append(b, '\n'), 0644)
}

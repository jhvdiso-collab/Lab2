package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"lab2_go/internal/ds"
	"lab2_go/internal/tasks"
	"os"
	"sort"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}
}
func run() error {
	task := flag.String("task", "", "asteroids, partition, subarrays, tree, bloom, lfu")
	file := flag.String("file", "", "JSON-файл множества")
	query := flag.String("query", "", "SETADD/SETDEL/SET_AT элемент")
	flag.Parse()
	if *file != "" || *query != "" {
		return setQuery(*file, *query)
	}
	dec := json.NewDecoder(os.Stdin)
	out := json.NewEncoder(os.Stdout)
	switch *task {
	case "asteroids":
		var a []tasks.Asteroid
		if e := dec.Decode(&a); e != nil {
			return e
		}
		r, e := tasks.Asteroids(a)
		if e != nil {
			return e
		}
		return out.Encode(r)
	case "partition":
		var a []int
		if e := dec.Decode(&a); e != nil {
			return e
		}
		l, r, d, e := tasks.Partition(a)
		if e != nil {
			return e
		}
		return out.Encode(struct {
			Left, Right []int
			Difference  int
		}{l, r, d})
	case "subarrays":
		var in struct {
			Array  []int
			Target int
		}
		if e := dec.Decode(&in); e != nil {
			return e
		}
		return out.Encode(tasks.Subarrays(in.Array, in.Target))
	case "tree":
		var t ds.Tree
		depths := []int{}
		for {
			var v int
			if _, e := fmt.Fscan(os.Stdin, &v); e != nil {
				return fmt.Errorf("ожидается последовательность с завершающим 0: %w", e)
			}
			if v == 0 {
				break
			}
			if d, ok := t.Insert(v); ok {
				depths = append(depths, d)
			}
		}
		return out.Encode(depths)
	case "lfu":
		var in struct {
			Capacity int
			Queries  [][]string
		}
		if e := dec.Decode(&in); e != nil {
			return e
		}
		if in.Capacity < 0 {
			return fmt.Errorf("отрицательная ёмкость")
		}
		c := ds.NewLFU(in.Capacity)
		result := []int{}
		for _, q := range in.Queries {
			if len(q) == 2 && q[0] == "GET" {
				v, ok := c.Get(q[1])
				if !ok {
					v = -1
				}
				result = append(result, v)
			} else if len(q) == 3 && q[0] == "SET" {
				var v int
				if _, e := fmt.Sscan(q[2], &v); e != nil {
					return e
				}
				c.Set(q[1], v)
			} else {
				return fmt.Errorf("неверный запрос %v", q)
			}
		}
		return out.Encode(result)
	case "bloom":
		var queries [][]string
		if e := dec.Decode(&queries); e != nil {
			return e
		}
		var t ds.FilteredTable
		for _, q := range queries {
			if len(q) == 3 && q[0] == "SET" {
				t.Put(q[1], q[2])
			} else if len(q) == 2 && q[0] == "GET" {
				v, ok := t.Get(q[1])
				if e := out.Encode(struct {
					Value string
					Found bool
				}{v, ok}); e != nil {
					return e
				}
			} else if len(q) == 2 && q[0] == "DEL" {
				t.Delete(q[1])
			} else {
				return fmt.Errorf("неверный запрос %v", q)
			}
		}
		return nil
	case "":
		flag.Usage()
		return nil
	default:
		return fmt.Errorf("неизвестное задание %q", *task)
	}
}
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

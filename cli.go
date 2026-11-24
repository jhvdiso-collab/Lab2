package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"lab2_go/internal/ds"
	"os"
)

func run() error {
	task := flag.String("task", "", "asteroids, partition, subarrays, tree, bloom, lfu, cuckoo, chaining")
	file := flag.String("file", "", "JSON-файл множества")
	query := flag.String("query", "", "SETADD/SETDEL/SET_AT элемент")
	flag.Parse()
	if *file != "" || *query != "" {
		return setQuery(*file, *query)
	}
	dec := json.NewDecoder(os.Stdin)
	out := json.NewEncoder(os.Stdout)
	switch *task {
	case "cuckoo":
		return runHashConsole(&ds.Cuckoo[string]{}, os.Stdin, os.Stdout)
	case "chaining":
		return runHashConsole(&ds.Table[string]{}, os.Stdin, os.Stdout)
	case "asteroids":
		return runAsteroids(dec, out)
	case "partition":
		return runPartition(dec, out)
	case "subarrays":
		return runSubarrays(dec, out)
	case "tree":
		return runTree(os.Stdin, out)
	case "lfu":
		return runLFU(dec, out)
	case "bloom":
		return runBloom(dec, out)
	case "":
		flag.Usage()
		return nil
	default:
		return fmt.Errorf("неизвестное задание %q", *task)
	}
}

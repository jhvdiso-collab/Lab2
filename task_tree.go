package main

import (
	"encoding/json"
	"fmt"
	"io"
	"lab2_go/internal/ds"
)

func runTree(input io.Reader, out *json.Encoder) error {
	var t ds.Tree
	depths := []int{}
	for {
		var v int
		if _, e := fmt.Fscan(input, &v); e != nil {
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
}

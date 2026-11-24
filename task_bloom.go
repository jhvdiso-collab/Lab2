package main

import (
	"encoding/json"
	"fmt"
	"lab2_go/internal/ds"
)

func runBloom(dec *json.Decoder, out *json.Encoder) error {
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
}

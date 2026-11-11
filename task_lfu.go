package main

import (
	"encoding/json"
	"fmt"
	"lab2_go/internal/ds"
	"strconv"
)

func runLFU(dec *json.Decoder, out *json.Encoder) error {
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
			v, e := strconv.Atoi(q[2])
			if e != nil {
				return e
			}
			c.Set(q[1], v)
		} else {
			return fmt.Errorf("неверный запрос %v", q)
		}
	}
	return out.Encode(result)
}

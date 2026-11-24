package main

import (
	"encoding/json"
)

func runSubarrays(dec *json.Decoder, out *json.Encoder) error {
	var in struct {
		Array  []int
		Target int
	}
	if e := dec.Decode(&in); e != nil {
		return e
	}
	return out.Encode(Subarrays(in.Array, in.Target))
}

func Subarrays(a []int, target int) [][]int {
	out := [][]int{}
	for i := range a {
		sum := 0
		for j := i; j < len(a); j++ {
			sum += a[j]
			if sum == target {
				out = append(out, append([]int(nil), a[i:j+1]...))
			}
		}
	}
	return out
}

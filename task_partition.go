package main

import (
	"encoding/json"
	"fmt"
	"lab2_go/internal/ds"
	"sort"
	"strconv"
)

func runPartition(dec *json.Decoder, out *json.Encoder) error {
	var a []int
	if e := dec.Decode(&a); e != nil {
		return e
	}
	l, r, d, e := Partition(a)
	if e != nil {
		return e
	}
	return out.Encode(struct {
		Left, Right []int
		Difference  int
	}{l, r, d})
}

// Partition uses subset-sum DP; predecessor links reconstruct the answer.
func Partition(values []int) ([]int, []int, int, error) {
	var set ds.Set
	for _, v := range values {
		if v <= 0 {
			return nil, nil, 0, fmt.Errorf("нужны натуральные числа")
		}
		set.Add(strconv.Itoa(v))
	}
	a := []int{}
	total := 0
	for _, k := range set.Values() {
		v, _ := strconv.Atoi(k)
		if v > int(^uint(0)>>1)-total {
			return nil, nil, 0, fmt.Errorf("сумма слишком велика")
		}
		a = append(a, v)
		total += v
	}
	sort.Ints(a)
	half := total / 2
	if half > 10000000 {
		return nil, nil, 0, fmt.Errorf("сумма слишком велика для DP (лимит половины суммы 10000000)")
	}
	prev := make([]int, half+1)
	for i := range prev {
		prev[i] = -1
	}
	prev[0] = -2
	for i, v := range a {
		for s := half; s >= v; s-- {
			if prev[s] == -1 && prev[s-v] != -1 {
				prev[s] = i
			}
		}
	}
	best := half
	for prev[best] == -1 {
		best--
	}
	chosen := make([]bool, len(a))
	for s := best; s > 0; {
		i := prev[s]
		chosen[i] = true
		s -= a[i]
	}
	left, right := []int{}, []int{}
	for i, v := range a {
		if chosen[i] {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return left, right, total - 2*best, nil
}

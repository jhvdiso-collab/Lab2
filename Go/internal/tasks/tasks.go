package tasks

import (
	"fmt"
	"lab2_go/internal/ds"
	"sort"
	"strconv"
)

type Asteroid struct {
	Size      int
	Direction string
}

func Asteroids(input []Asteroid) ([]Asteroid, error) {
	var s ds.Stack[Asteroid]
	for _, a := range input {
		if a.Size <= 0 || (a.Direction != "L" && a.Direction != "R") {
			return nil, fmt.Errorf("ожидается положительный размер и L/R")
		}
		for a.Size > 0 {
			b, ok := s.Pop()
			if !ok {
				s.Push(a)
				break
			}
			if b.Direction != "R" || a.Direction != "L" {
				s.Push(b)
				s.Push(a)
				break
			}
			if b.Size > a.Size {
				b.Size -= a.Size
				s.Push(b)
				break
			}
			a.Size -= b.Size
		}
	}
	r := []Asteroid{}
	for {
		a, ok := s.Pop()
		if !ok {
			break
		}
		r = append(r, a)
	}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r, nil
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

package main

import (
	"encoding/json"
	"fmt"
	"lab2_go/internal/ds"
)

func runAsteroids(dec *json.Decoder, out *json.Encoder) error {
	var a []Asteroid
	if e := dec.Decode(&a); e != nil {
		return e
	}
	r, e := Asteroids(a)
	if e != nil {
		return e
	}
	return out.Encode(r)
}

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

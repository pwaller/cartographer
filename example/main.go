package main

import "log"

func main() {
	m := map[int]struct{}{}
	for k := range m {
		log.Printf("k: %v", k)
	}
	for k := range f() {
		log.Printf("k: %v", k)
	}
	t := T{}
	for k := range t.m {
		log.Printf("k: %v", k)
	}
	for k := range (T{}).m {
		log.Printf("k: %v", k)
	}
	for k := range (&T{}).m {
		log.Printf("k: %v", k)
	}
}

func f() map[int]struct{} {
	return nil
}

// T ...
type T struct {
	m map[int]struct{}
}

// M ...
type M map[int]struct{}

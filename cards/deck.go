package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, crd := range d {
		fmt.Println(i, crd)
	}
}

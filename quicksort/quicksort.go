package main

import "golang.org/x/exp/constraints"

func quicksort[T constraints.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	menores := []T{}
	maiores := []T{}

	for _, v := range arr[1:] {
		if v <= pivot {
			menores = append(menores, v)
		} else {
			maiores = append(maiores, v)
		}
	}

	menores = quicksort(menores)
	maiores = quicksort(maiores)

	return append(append(menores, pivot), maiores...)
}

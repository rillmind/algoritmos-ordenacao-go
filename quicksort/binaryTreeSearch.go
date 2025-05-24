package main

import "golang.org/x/exp/constraints"

type No[T constraints.Ordered] struct {
	esq   *No[T]
	dir   *No[T]
	valor T
}

func inserir[T constraints.Ordered](no *No[T], valor T) *No[T] {
	if no == nil {
		return &No[T]{valor: valor}
	}
	if valor < no.valor {
		no.esq = inserir(no.esq, valor)
	} else if valor > no.valor {
		no.dir = inserir(no.dir, valor)
	}
	return no
}

func buscar[T constraints.Ordered](no *No[T], valor T) bool {
	if no == nil {
		return false
	}
	if valor == no.valor {
		return true
	} else if valor < no.valor {
		return buscar(no.esq, valor)
	} else {
		return buscar(no.dir, valor)
	}
}

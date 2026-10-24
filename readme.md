# Ordenação

projeto com alguns algoritmos de ordenação em go

## O que faz

Implementa o **quicksort** em Go (`quicksort/quicksort.go`), genérico com `constraints.Ordered`, ordenando
tanto números quanto strings. Também traz uma busca em árvore binária de busca
(`quicksort/binaryTreeSearch.go`). O `quicksort/main.go` demonstra os dois.

## Por que existe

Exercício da disciplina de Estrutura de Dados: implementar algoritmos de ordenação e busca do zero
em Go.

## Instalação

Com o git e o go instalados basta rodar as seguintes linhas no terminal:

```sh
git clone https://github.com/rillmind/revisaoED
cd revisaoED
go run .
```


## Correção

Os comandos de Instalação acima estão errados: eles clonam e entram em `revisaoED`, que é outro
repositório. O certo é:

```sh
git clone https://github.com/rillmind/algoritmos-ordenacao-go
cd algoritmos-ordenacao-go
go run ./quicksort
```

(O código fica no pacote `quicksort/`, então `go run .` na raiz falha com "no Go files".)

## Como testar

Não há testes automatizados. `go vet ./...` confere o código. `go run ./quicksort` mostra:

```
Antes: [10 5 2 3 7 6 1 8 9 4]
Depois: [1 2 3 4 5 6 7 8 9 10]
Antes: [uva banana maçã laranja abacaxi]
Depois: [abacaxi banana laranja maçã uva]
Buscar 7: true
Buscar 11: false
Buscar "maçã": true
Buscar "melancia": false
```

Confira que o "Depois" sai ordenado e que as buscas respondem `true`/`false` corretamente.

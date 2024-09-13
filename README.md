# Trabalho 6 - GraphGen

## Feito pelos alunos:
 - Daniel Kenichi Tiago Tateishi RA: 790837
 - Rodrigo Pavão Coffani Nunes RA: 800345 

## Sobre o trabalho
Um compilador implementado em GO de uma linguagem que descreve relacionamentos de pessoas e grupos, e então gera código em python para desenhar um grafo visualizando as relações de uma pessoa ou grupo utilizando a lib networkx

## Resuisitos

[Antlr](https://github.com/antlr/antlr4/blob/master/doc/go-target.md) \
[Go](https://go.dev/doc/install) \
[Python3](https://www.python.org/downloads/) 

## Getting Started

Primeiro é necessário utilizar o antlr para gerar o código utilizado pelo compilador em GO a partir da gramática definida em Antlr/Graphge.g4

```bash
antlr4 -Dlanguage=Go -o . Antlr/GraphGen.g4
```

Após isso, compile o programa do compilador com

```bash
go build -o graphGen GraphGen/*.go
```

No diretório test_files há exemplos de input para o compilador. Para executar basta utilizar o comando 

```bash
./graphgen <input_file> <output_file>
```
[Demo do trabalho](https://youtu.be/im3inUHMzQU)

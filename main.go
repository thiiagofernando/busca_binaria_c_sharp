package main

import (
	"fmt"
	"sort"
	"strings"
)

func geradorNomes() []string {
	names := "Alice,Ana,André,Beatriz,Bruno,Camila,Carlos,Carolina,Catarina,Daniel,Davi,Eduardo,Elisa,Fabiana,Felipe,Fernanda,Gabriel,Giovanna,Guilherme,Gustavo,Helena,Igor,Isabela,Joana,João,José,Julia,Juliana,Leonardo,Letícia,Lucas,Luiza,Marcela,Marcelo,Maria,Mariana,Mateus,Miguel,Natália,Nicolas,Patrícia,Pedro,Rafael,Renata,Rodrigo,Samuel,Sara,Sofia,Tânia,Thiago,Vanessa,Vitor,Amanda,Ângela,Arthur,Bruna,Caio,Carolina,Cecília,Clara,Cristian,Daniela,Diego,Eduarda,Enzo,Esther,Fabrício,Fernanda,Flávia,Francisco,Gabriela,Giovani,Helena,Henrique,Ingrid,Isabela,Isadora,Ícaro,Jéssica,Júlio,Larissa,Laura,Lívia,Luan,Marcela,Mariana,Marília,Matheus,Melissa,Miguel,Natália,Nicole,Otávio,Paola,Patrícia,Paula,Paulo,Pedro,Rafaela,Raquel,Renan,Ricardo,Roberto,Rogério,Sabrina,Samuel,Sandra,Sara,Silvana,Sônia,Taís,Tamires,Túlio,Valentina,Valter,Vanessa,Vinícius,Vitor,Vivian,Yuri"
	splited := strings.Split(names, ",")
	allNames := []string{}
	for _, v := range splited {
		allNames = append(allNames, strings.TrimLeft(v, " "))
	}
	sort.Strings(allNames)
	return allNames
}

func main() {
	listaDeNomes := geradorNomes()
	nomeProcurado := "Thiago"

	inicio := 0
	final := len(listaDeNomes) - 1

	for inicio <= final {
		meio := (final + inicio) / 2
		nomeAtual := listaDeNomes[meio]

		if nomeAtual == nomeProcurado {
			fmt.Println("NOME ENCONTRADO NA LISTA", nomeAtual)
			break
		}

		if nomeAtual > nomeProcurado {
			fmt.Println("nome atual menor que o procurado")
			final = meio - 1
		}

		if nomeAtual < nomeProcurado {
			fmt.Println("nome atual maior que o procurado")
			inicio = meio + 1
		}
	}

}

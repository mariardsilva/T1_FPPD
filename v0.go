//V0 - jantar dos filósofos sequencial

package main

import (
	"fmt"
	"time"
)

func Filosofo(id, iteration int) {
	tempoPensando := 2 // segundos pensando
	tempoComendo := 3   // segundos comendo

	// Simula o pensar
	for s := 1; s <= (id+1)*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" F%d Pensando (Ciclo %d)\n", id+1, iteration)
	time.Sleep(time.Duration(tempoPensando) * time.Second)

	// Simula o comer
	for s := 1; s <= (id+1)*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" F%d Comendo (Ciclo %d)\n", id+1, iteration)
	time.Sleep(time.Duration(tempoComendo) * time.Second)
}

func main() {
	filosofos := 5
	rodadas := 20 // Cada filósofo deve executar 20 ciclos conforme o enunciado

	fmt.Println("\n      [F1]  [F2]  [F3]  [F4]  [F5]\n")
	start := time.Now()

	// Execução estritamente sequencial (um filósofo executa por vez)
	for i := 1; i <= rodadas; i++ {
		for j := 0; j < filosofos; j++ {
			Filosofo(j, i)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("\n[V0] Jantar Sequencial levou %s\n\n", elapsed)
}
// V3 - Jantar dos Filósofos concorrente aprimorado (ordenação dos garfos)
package main

import (
	"fmt"
	"sync"
	"time"
)

func Filosofo(id int, garfoMenor, garfoMaior *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	rounds := 20

	for i := 1; i <= rounds; i++ {
		// Pensa
		fmt.Printf("Filósofo %d está pensando (Ciclo %d)\n", id+1, i)
		time.Sleep(1 * time.Second)

		// Sempre pega o garfo de menor índice primeiro: evita a espera circular (sem deadlock)
		garfoMenor.Lock()
		garfoMaior.Lock()

		// Come
		fmt.Printf("==> Filósofo %d pegou os garfos e está COMENDO (Ciclo %d)\n", id+1, i)
		time.Sleep(1 * time.Second)

		garfoMaior.Unlock()
		garfoMenor.Unlock()
	}
}

func main() {
	filosofos := 5
	var garfos [5]sync.Mutex
	var wg sync.WaitGroup

	fmt.Println("\nIniciando V3...")
	start := time.Now()

	for i := 0; i < filosofos; i++ {
		wg.Add(1)
		// Garfos do filósofo i: 'i' e '(i+1) % 5'. O de menor índice é sempre pego primeiro
		menor, maior := i, (i+1)%filosofos
		if menor > maior {
			menor, maior = maior, menor
		}
		go Filosofo(i, &garfos[menor], &garfos[maior], &wg)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("\n[V3] Tempo total com ordenação dos garfos: %s\n", elapsed)
}
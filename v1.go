// V1 - jantar dos filósofos concorrente com deadlock
package main

import (
	"fmt"
	"sync"
	"time"
)

func Filosofo(id int, garfoEsq, garfoDir *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	rodadas := 20

	for i := 1; i <= rodadas; i++ {
		// Pensa
		time.Sleep(1 * time.Second)

		// Condição de deadlock, todos pegam o garfo da esquerda primeiro ao mesmo tempo
		garfoEsq.Lock()
		fmt.Printf("Filósofo %d pegou o garfo esquerdo.\n", id+1)
		
		// Pequena pausa para garantir que todos peguem o esquerdo e fiquem bloqueados no direito
		time.Sleep(50 * time.Millisecond) 

		garfoDir.Lock() // Se todos chegarem aqui juntos, o programa trava (Deadlock)
		fmt.Printf("Filósofo %d pegou o garfo direito e está comendo.\n", id+1)

		time.Sleep(1 * time.Second)

		// Libera os recursos
		garfoDir.Unlock()
		garfoEsq.Unlock()
	}
}

func main() {
	filosofos := 5
	var garfos [5]sync.Mutex
	var wg sync.WaitGroup

	fmt.Println("\nIniciando V1 (Sujeita a Deadlock)...")

	for i := 0; i < filosofos; i++ {
		wg.Add(1)
		// Garfo esquerdo é o de índice 'i', garfo direito é o de índice '(i+1) % 5'
		go Filosofo(i, &garfos[i], &garfos[(i+1)%filosofos], &wg)
	}

	wg.Wait()
	fmt.Println("Fim da execução V1 (Se congelar antes disso, o deadlock ocorreu com sucesso).")
}
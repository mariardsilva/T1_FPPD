//V2 - Jantar dos Filósofos com sincronização excessiva
package main

import (
    "fmt"
    "sync"
    "time"
)

func Filosofo(id int, Mutex *sync.Mutex, wg *sync.WaitGroup) {
    defer wg.Done()
    rounds := 20

    for i := 1; i <= rounds; i++ {
        // Pensa
        fmt.Printf("Filósofo %d está pensando (Ciclo %d)\n", id+1, i)
        time.Sleep(1 * time.Second)

        // SINCRONIZAÇÃO EXCESSIVA: Um mutex global (garçom) tranca o acesso aos recursos.
        // Apenas um filósofo por vez consegue passar daqui, serializando a execução.
        Mutex.Lock()

        // Come (recursos adquiridos sob proteção da trava global)
        fmt.Printf("==> Filósofo %d pegou os garfos e está COMENDO (Ciclo %d)\n", id+1, i)
        time.Sleep(1 * time.Second)

        Mutex.Unlock()
    }
}

func main() {
    filosofos := 5
    var Mutex sync.Mutex // O "Garçom" ou trava global de exclusão mútua
    var wg sync.WaitGroup

    fmt.Println("\nIniciando V2...")
    start := time.Now()

    for i := 0; i < filosofos; i++ {
        wg.Add(1)
        go Filosofo(i, &Mutex, &wg)
    }

    wg.Wait()
    elapsed := time.Since(start)
    fmt.Printf("\n[V2] Tempo total com sincronização excessiva: %s\n", elapsed)
}
package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []string{"a", "b", "c", "d", "e"}
	sliceLength := len(slice)
	var wg sync.WaitGroup
	//kaç kere olacağını belirtir
	wg.Add(sliceLength)
	fmt.Println("Running for loop...")
	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			//bir azaltır
			defer wg.Done()
			val := slice[i]
			fmt.Printf("i: %v, val: %v\n", i, val)
		}(i)
	}
	//bitene kadar bekler
	wg.Wait()
	fmt.Println("Finished for loop")
}

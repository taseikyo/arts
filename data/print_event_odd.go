/*
* @Date:   2025-03-22 15:12:29
* @Author: Lewis Tian (taseikyo@gmail.com)
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				fmt.Println("thread1: ", i)
			}
			ch <- 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			if i%2 == 1 {
				fmt.Println("thread2: ", i)
			}
			<-ch
		}
	}()
	wg.Wait()
}

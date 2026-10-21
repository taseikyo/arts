/*
* @Date:   2025-03-22 15:21:43
* @Author: Lewis Tian (taseikyo@gmail.com)
*/
package main

import (
	"fmt"
)

func main() {
	print(5)
}

func print(num int) {
	var chanSlice []chan int
	for i := 0; i < num; i++ {
		chanSlice = append(chanSlice, make(chan int))
	}
	quitChan := make(chan struct{})

	res := 1
	next := 0
	for i := 0; i < num; i++ {
		go func(i int) {
			for {
				<-chanSlice[i]
				if res > 100 {
					quitChan <- struct{}{}
					return
				}
				fmt.Println("thread", i, res)
				res++

				if next == num-1 {
					next = 0
				} else {
					next++
				}
				chanSlice[next] <- 1
			}

		}(i)
	}
	chanSlice[0] <- 1
	select {
	case <-quitChan:
		fmt.Println("exit")
	}

}

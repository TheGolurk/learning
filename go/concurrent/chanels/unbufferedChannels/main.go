package main

import (
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		/* if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		} */

		/*for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}*/
		/*for msg := range ch {
			fmt.Println(msg)
		}*/

		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		/*close(ch)
		ch <- 3 */

		/*for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)*/

		wg.Done()
	}(ch, wg)

	wg.Wait()
}

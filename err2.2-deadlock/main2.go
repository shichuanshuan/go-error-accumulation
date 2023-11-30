package main

import (
	"fmt"
	"time"
)

func talk(msg string, sleep int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			// input1 和 input2 都会执行，但每次 select 只会选择其中一个 case 执行；
			// 所以 input1 和 input2 的结果， 必然有一个被丢弃了，也就不会写入 ch 中；
			// 因此，一共输出 5 次，另外 5 次结果丢掉了
			select {
			case ch <- <-input1:
			case ch <- <-input2:
			}
		}
	}()
	return ch
}

func main() {
	ch := fanIn(talk("A", 10), talk("B", 1000))

	// 循环 10 次，只获得 5 次结果，所以输出 5 次后，报死锁
	for i := 0; i < 10; i++ {
		fmt.Printf("%q\n", <-ch)
	}
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	intChan := make(chan int)

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT)

	go func() {
		i := 1
		for {
			fmt.Println(<-intChan)
			i++
		}
	}()
	go square(intChan, signChan)
}

func square(iCh chan int, sCh chan os.Signal) {
	i := 1
	for {
		select {
		case iCh <- i * i:
			i++
		case <-sCh:
			fmt.Println("выхожу из программы")
			return
		}
	}
}

/*
	intChan := make(chan int)
	go square(intChan)

	for i := range intChan {
		fmt.Println(i)
	}

}

func square(iCh chan int) {
	i := 1
	for {
		iCh <- i * i
		time.Sleep(1 * time.Second)
		i++
	}
}
*/

/*
	intChan := make(chan int) //создали канал для чисел

	signChan := make(chan os.Signal, 1) //создали канал для сигналов

	signal.Notify(signChan, syscall.SIGINT)

	go square(intChan, signChan)

}

func square(iCh chan int, sCh chan os.Signal) {
	i := 1
	for {
		select {
		case <-sCh:
			fmt.Println("выхожу из программы")
			return
		default:
			iCh <- i * i
			fmt.Println(<-iCh)
		}
		i++
	}
*/

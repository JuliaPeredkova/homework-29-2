package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	chanInt := make(chan int)
	chanSignal := make(chan os.Signal, 1)
	signal.Notify(chanSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for i := range chanInt {
			fmt.Println(i)
		}
	}()

	square(chanInt, chanSignal)

}

func square(iCh chan int, sCh chan os.Signal) {
	i := 1
	for {
		select {
		case <-sCh:
			fmt.Println("Выхожу из программы")
			return
		case iCh <- i * i:
			time.Sleep(1 * time.Second)
			i++
		}

	}
}

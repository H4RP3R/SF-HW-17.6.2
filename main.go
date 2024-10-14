// Напишите код, в котором имеются два канала сообщений из целых чисел так,
// чтобы приём сообщений всегда приводил к блокировке. Приёмом сообщений из
// обоих каналов будет заниматься главная горутина. Сделайте так, чтобы во
// время такого «бесконечного ожидания» сообщений выполнялась фоновая работа
// в виде вывода текущего времени в консоль.

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	secTicker := time.NewTicker(1 * time.Second)
	hourTicker := time.NewTicker(1 * time.Hour)

	go sendToChannel(c1, hourTicker)
	go sendToChannel(c2, hourTicker)

	for {
		select {
		case <-c1:
			fmt.Println("Message from channel 1")
		case <-c2:
			fmt.Println("Message from channel 2")
		default:
			now := <-secTicker.C
			fmt.Println(now.Format("15:04:05"))
		}
	}

}

func sendToChannel(c chan<- int, ticker *time.Ticker) {
	for {
		<-ticker.C
		c <- 1
	}
}

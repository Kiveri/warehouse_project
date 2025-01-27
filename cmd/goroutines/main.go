package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	hello := "Привет❤️❤️"
	for h, i := range hello {
		fmt.Println(h, i)
	}

	//cache := NewCache()
	//
	//wg := &sync.WaitGroup{}
	//wg.Add(100)
	//for i := 0; i < 100; i++ {
	//	go func() {
	//		defer wg.Done()
	//		cache.Set("counter", int64(i))
	//	}()
	//}
	//wg.Wait()

	//ch := make(chan int)
	//
	//go sum(1, 2, ch)
	//
	//time.Sleep(time.Second * 5)
	//println(<-ch)
	//println(<-ch)

	// читаем из закрытого канала дефолтное значение типа
	// закрываем канал вместе где пишем в него
	// запись в закрытый канал приведет к панике
	// когда передавать данные между горутинами
	// когда нужно сигналить другой горутине

	ch := make(chan struct{}, 100)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go makeWork(ch, wg)
	go makeWorkAgain(ch, wg)
	wg.Wait()
}

func makeWork(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		println("работаем")
		time.Sleep(time.Millisecond * 500)
	}

	ch <- struct{}{}
}

func makeWorkAgain(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	<-ch

	println("УРАААА РАБОТАЕМ")
}

//func sum(a, b int, ch chan int) {
//	println("работает")
//	defer close(ch)
//	ch <- a + b
//	ch <- a * b
//	println("вышли")
//}

// разницу между горутиной и сис потоком
// виды мьютексов
// вэйт группы
// каналы

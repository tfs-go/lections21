//nolint: unused
package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"

	"github.com/tfs-go/lections21/lection02/code/exchange"
)

type Candle struct {
	name       string
	openPrice  float64
	closePrice float64
}

func buildAndPrint() {
	c1 := Candle{name: "BABA", openPrice: 1.0, closePrice: 1.5}
	fmt.Println(c1)

	c2 := Candle{name: "SPCE", openPrice: 0.55}
	fmt.Println(c2)

	c3 := Candle{}
	fmt.Println(c3)

	c4 := Candle{"AMZN", 1.0, 2.0}
	fmt.Println(c4)
}

func set() {
	c1 := Candle{name: "BABA", openPrice: 1.0, closePrice: 1.5}
	fmt.Println(c1)

	// копирование значения в переменную
	a := c1.closePrice
	fmt.Println(a)

	// запись значений в поля структуры
	c1.name = "AMZN"
	c1.openPrice += 3.0
	fmt.Println(c1)
}

func compare() {
	c1 := Candle{name: "BABA", openPrice: 1.0, closePrice: 1.5}
	c2 := c1

	// можно сравнивать, если все поля можно сравнивать
	fmt.Println(c1 == c2)

	c2.name = "AMZN"
	fmt.Println(c1 == c2)

	//nolint: gocritic,staticcheck
	fmt.Println(Candle{} == Candle{})

	/*
		// слайсы нельзя сравнить
		type notComparable struct {
			s []string
		}
		a, b := notComparable{}, notComparable{}
		fmt.Println(a == b)
	*/
}

func sizeOf() {
	type a struct{}

	_ = struct{}{}

	fmt.Println(unsafe.Sizeof(struct{}{}))
	fmt.Println(unsafe.Sizeof([1000]struct{}{}))
	fmt.Println(unsafe.Sizeof([1000]*int{}))
}

func methods() {
	// конструктор не нужен
	_ = sync.Mutex{}

	_ = new(sync.Mutex)

	// не элегантно; непонятно, какие поля обязательные для коректной работы
	_ = exchange.Exchange{
		Name: "MOEX",
	}

	// удобно и красиво
	e, _ := exchange.New("MOEX", time.Now(), time.Now().Add(time.Hour))

	fmt.Println(e.Duration())

	// просто еще один пример метода
	_ = time.Now().Add(time.Second)
}

type stock struct {
	stockType string
	exchange.Exchange
}

func embedding() {
	e, _ := exchange.New("MOEX", time.Now(), time.Now().Add(time.Hour*8))
	s := stock{
		stockType: "exchange",
		Exchange:  e,
	}

	fmt.Println(s.stockType, s.Name, s.Duration())
}

func pointerValue() {
	e, _ := exchange.New("MOEX", time.Now(), time.Now().Add(time.Hour*8))
	fmt.Println(e.Duration())

	e1, e2 := e, e
	e1.UpdateEndTime(time.Now().Add(time.Hour * 2))
	fmt.Println(e1.Duration())

	e2.UpdateEndTimeValue(time.Now().Add(time.Hour * 2))
	fmt.Println(e2.Duration())

	_ = func(
		s []string,
		c chan string,
		m map[string]bool,
		f func(),
		i interface{},
	) {
		// все это можно передавать by value
	}
}

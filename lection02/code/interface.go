package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

type CandleFetcher interface {
	Fetch() (Candle, error)
}

type MOEXFetcher struct{}

func (m MOEXFetcher) Fetch() (Candle, error) {
	fmt.Println("received candle from MOEX")
	return Candle{}, nil
}

type SPBFetcher struct{}

func (s SPBFetcher) Fetch() (Candle, error) {
	fmt.Println("received candle from SPB")
	return Candle{}, nil
}

type MockFetcher struct{}

func (s MockFetcher) Fetch() (Candle, error) {
	fmt.Println("mock method Fetch called")
	return Candle{}, nil
}

func fetch() {
	for _, f := range []CandleFetcher{MOEXFetcher{}, SPBFetcher{}, MockFetcher{}} {
		_, _ = f.Fetch()
	}
}

type service struct {
	l *log.Logger
	f CandleFetcher
}

func NewService(l *log.Logger, f CandleFetcher) service {
	return service{
		l: l,
		f: f,
	}
}

func (s service) Run() {
	for {
		c, err := s.f.Fetch()
		if err != nil {
			s.l.Println("error happened", err)
			break
		}
		fmt.Println(c)
	}
}

func useInterface() {
	/*
		logger := log.New(os.Stdout, "", 0)

		moex := NewService(logger, MOEXFetcher{})
		moex.Run()

		spb := NewService(logger, SPBFetcher{})
		spb.Run()

		// удобно для тестов
		testService := NewService(logger, MockFetcher{})
		testService.Run()
	*/
}

func iNaming() {
	var (
		_ io.Reader
		_ io.Writer
		_ fmt.Stringer
	)
}

func compose() {
	var (
		_ io.ReadWriter
		_ io.WriteCloser
	)
}

func iError() {
	var b *bytes.Buffer

	f := func(i io.Writer) {
		if i != nil {
			_, _ = i.Write([]byte("Hi, Mark!"))
		}
	}

	f(b)
}

type Person struct {
	name string
}

func (p *Person) Name() string {
	return p.name
}

type Namer interface {
	Name() string
}

func associatedValues() {
	f := func(n Namer) {
		fmt.Println(n.Name())
	}
	f(nil)
}

func emptyInterface() {
	var b interface{}

	b = "Hello"
	_ = interface{}("Hello")
	fmt.Printf("%T, %+v\n", b, b)

	b = []int{1, 2, 3}
	_ = interface{}([]int{1, 2, 3})
	fmt.Printf("%T, %+v\n", b, b)

	b = func() { fmt.Println("Hi, Mark!") }
	_ = interface{}(func() { fmt.Println("Hi, Mark!") })
	fmt.Printf("%T, %+v\n", b, b)
}

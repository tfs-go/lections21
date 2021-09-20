//nolint: unused
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/tfs-go/lections21/lection02/code/exchange"
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

type Service struct {
	l *log.Logger
	f CandleFetcher
}

func NewService(l *log.Logger, f CandleFetcher) Service {
	return Service{
		l: l,
		f: f,
	}
}

func (s Service) Run() {
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
	logger := log.New(os.Stdout, "", 0)

	moex := NewService(logger, MOEXFetcher{})
	moex.Run()

	spb := NewService(logger, SPBFetcher{})
	spb.Run()

	// удобно для тестов
	testService := NewService(logger, MockFetcher{})
	testService.Run()
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

func implementCheck() {
	var (
		_ io.Reader = new(bytes.Buffer)
		// _ io.Reader = "abc" // ошибка компиляции
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

	// slice of anything!
	_ = []interface{}{1, "a", map[string]string{}, struct{}{}, func() {}, make(chan string)}

	// map of anything!
	_ = map[string]interface{}{"int": 1, "string": "Hi, Mark!", "func": func() {}}

	// struct with anything!
	_ = struct {
		A interface{}
		B interface{}
	}{
		A: [1000]int{},
		B: 15.456,
	}
}

func typeAssertion() {
	var b interface{}

	b = "Hi, Mark!" // b still has type interface{}

	// assert to simple type
	if s, ok := b.(string); ok {
		fmt.Printf("b is string: %s\n", s)
	} else {
		fmt.Printf("b was not a string: %s\n", s)
	}

	// assert to complex type
	b = exchange.Exchange{Name: "MOEX"} // b still has type interface{}
	// fmt.Println(b.Name) // no access to object fields
	if e, ok := b.(exchange.Exchange); ok {
		fmt.Println(e.Name)
	}

	// assert to interface
	b = new(bytes.Buffer) // b still has type interface{}
	if r, ok := b.(io.Reader); ok {
		_, _ = r.Read([]byte{})
	}

	// multiple assertions
	switch a := b.(type) {
	case string: // default types
		fmt.Println("b is string", a)
	case exchange.Exchange: // structs
		fmt.Println("b is Exchange", a)
	case io.Reader: // interfaces
		fmt.Println("b is Reader", a)
	case error: // ???
		fmt.Println("b is error", a)
	}
}

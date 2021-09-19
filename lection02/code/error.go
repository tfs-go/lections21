//nolint: unused
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func errorDefinition() {
	var _ error

	var _ error = strError("")

	var et funcError = func() error {
		return errors.New("eat this")
	}
	var _ error = et
}

type strError string

func (s strError) Error() string {
	return string(s)
}

type funcError func() error

func (f funcError) Error() string {
	return f().Error()
}

// errPractices

func errPractices() {
	// return nil(zero), err
	i, err := alwaysError()
	fmt.Println(i, err)

	// return not-zero, nil
	i, err = alwaysInt()
	fmt.Println(i, err)

	// can return nil, nil for slice
	s, err := alwaysSlice()
	fmt.Println(s, err)
	if err == nil {
		_ = len(s)
		_ = cap(s)
		for range s {
			fmt.Println("i'm iterating over empty slice")
		}
	}
}

func alwaysError() (int, error) {
	return 0, errors.New("luck is not on your side")
}

func alwaysInt() (int, error) {
	return 1, nil
}

func alwaysSlice() ([]int, error) {
	return nil, nil
}

// errLogContext

func errLogContext() {
	// guess what went wrong
	_, err := anythingCanGoWrong()
	if err != nil {
		fmt.Println("err happened: ", err)
	}

	// guess what went wrong
	_, err = anythingCanGoWrongButBetter()
	if err != nil {
		fmt.Println("err happened: ", err)
	}
}

func anythingCanGoWrong() (int, error) {
	var sum int

	first, err := randomErr()
	if err != nil {
		return 0, err
	}
	sum += first

	second, err := randomErr()
	if err != nil {
		return 0, err
	}
	sum += second

	third, err := randomErr()
	if err != nil {
		return 0, err
	}
	sum += third

	fourth, err := randomErr()
	if err != nil {
		return 0, err
	}
	sum += fourth

	return sum, nil
}

func anythingCanGoWrongButBetter() (int, error) {
	var sum int

	first, err := randomErr()
	if err != nil {
		return 0, fmt.Errorf("first process: %w", err)
	}
	sum += first

	second, err := randomErr()
	if err != nil {
		return 0, fmt.Errorf("second process: %w", err)
	}
	sum += second

	third, err := randomErr()
	if err != nil {
		return 0, fmt.Errorf("third process: %w", err)
	}
	sum += third

	fourth, err := randomErr()
	if err != nil {
		return 0, fmt.Errorf("fourth process: %w", err)
	}
	sum += fourth

	return sum, nil
}

func randomErr() (int, error) {
	i := rand.Intn(1000)
	if i > 650 {
		return 0, fmt.Errorf("error №%d", i)
	}
	return i, nil
}

// errAsIs

type CustomErr struct {
	Msg  string
	Code int
}

func (c CustomErr) Error() string {
	return c.Msg
}

var ErrNotFound = CustomErr{Msg: "not found", Code: 404}

func errAsIs() {
	f := func() error {
		return fmt.Errorf("opening file error: %w", ErrNotFound)
	}

	err := f()

	_, ok := err.(CustomErr)
	fmt.Println(ok)

	ok = errors.Is(err, ErrNotFound)
	fmt.Println(ok)

	ok = errors.Is(err, &ErrNotFound)
	fmt.Println(ok)

	var ce CustomErr
	ok = errors.As(err, &ce)
	fmt.Printf("%t, %v, %v", ok, ce.Msg, ce.Code)
}

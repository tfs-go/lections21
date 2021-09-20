//nolint: unused
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

func filesOpen() {
	_, err := os.Open("input.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file is not exists")
	}

	// manually allocate space
	f, err := os.Open("lection02/code/input2.txt")
	if err != nil {
		panic(err)
	}
	stat, _ := f.Stat()
	buf := make([]byte, stat.Size())
	_, _ = f.Read(buf)
	_ = f.Close()
	fmt.Println(string(buf))

	// ioutil
	f2, _ := os.Open("lection02/code/input2.txt")
	data, _ := ioutil.ReadAll(f2)
	defer f2.Close()
	fmt.Println(data)
}

func filesCreate() {
	f, _ := os.Create("lection02/code/joke.txt") // create or truncate
	defer f.Close()

	_, _ = f.WriteString("- Как заинтриговать глупца?\n")

	_, _ = f.Write([]byte("- Как?\n"))

	_, _ = fmt.Fprint(f, "- Потом расскажу.")
}

func fileJSON() {
	type User struct {
		Name   string
		Active bool
		age    uint
	}

	data, _ := json.Marshal(User{Name: "Shamil", Active: false, age: 30})
	fmt.Println(string(data))

	var u User
	_ = json.Unmarshal(data, &u)
	fmt.Println(u)

	type UserJ struct {
		Name string `json:"name"`

		Password string `json:"-"`

		Registered bool  `json:"registered,omitempty"`
		Active     *bool `json:"active,omitempty"`

		age uint
	}

	f := false
	data, _ = json.Marshal(UserJ{
		Name:     "Shamil",
		Password: "admin",

		Registered: false,
		Active:     &f,
		age:        30,
	})
	fmt.Println(string(data))
}

func fileYAML() {
	type User struct {
		Name   string
		Active bool
		age    int
	}

	data, _ := yaml.Marshal(User{Name: "Shamil", Active: false, age: 30})
	fmt.Println(string(data))

	type DB struct {
		DBAddress string `yaml:"host"`
		Username  string `yaml:"login"`
		Password  string `yaml:"pwd"`
	}

	type Config struct {
		Postgres DB `yaml:"postgres"`
	}

	data, _ = yaml.Marshal(Config{
		Postgres: DB{
			DBAddress: "localhost:5432",
			Username:  "admin",
			Password:  "dnkroz",
		},
	})
	fmt.Println(string(data))

	var a Config
	_ = yaml.Unmarshal(data, &a)
	fmt.Println(a)
}

type Date struct {
	time.Time
}

var dFormat = "2006-01-02"

func (d Date) MarshalJSON() ([]byte, error) {
	s := d.Format(dFormat)
	return []byte(fmt.Sprintf(`"%s"`, s)), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	s := bytes.Trim(data, "\"")
	t, err := time.Parse(dFormat, string(s))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

//nolint: govet,staticcheck
func customMarshal() {
	type DateTime struct {
		d Date      `json:"date"`
		t time.Time `json:"time"`
	}

	now := time.Now()
	data, _ := json.Marshal(DateTime{
		d: Date{Time: now},
		t: now,
	})

	fmt.Println(string(data))

	var dt DateTime
	_ = json.Unmarshal(data, &dt)

	fmt.Printf("%+v\n", dt)
}

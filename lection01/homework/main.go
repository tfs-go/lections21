package main

import (
	"fmt"

	"github.com/tfs-go/lections21/lection01/homework/figure"
)

func main() {
	fmt.Println(figure.Construct(figure.SandglassFiller, figure.CharModifier("!", 34), figure.SizeModifier(15)))
}
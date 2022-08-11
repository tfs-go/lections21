package main

import "fmt"

type space [][]string
type spaceАlgorithm func(*space)
type propertiesOfSpace func(*space)

const (
	Red    string = "\033[31m"
	Green  string = "\033[32m"
	Yellow string = "\033[33m"
	Blue   string = "\033[34m"
	Purple string = "\033[35m"
	Gray   string = "\033[37m"
	White  string = "\033[39m"
)

func parseInt(numberString string) int {
	var number int
	fmt.Sscan(numberString, &number)
	return number
}
func spaceStandard(sand *space) {
	makeMemory(sand, 15)
	prOfSpace := setStars("X")
	prOfSpace(sand)
	prOfSpace = setColor(White)
	prOfSpace(sand)
}
func makeMemory(field *space, size int) {
	*field = make([][]string, size)
	for i := range *field {
		(*field)[i] = make([]string, size)
	}
}
func setSize(sizeString string) propertiesOfSpace {
	return func(field *space) {
		safety := (*field)[0][0]
		size := parseInt(sizeString)
		makeMemory(field, size)
		fsafety := setStars(safety)
		fsafety(field)
	}
}
func setColor(color string) propertiesOfSpace {
	return func(field *space) {
		fmt.Print(color)
	}
}
func setStars(stars string) propertiesOfSpace {
	return func(field *space) {
		for y, axisY := range *field {
			for x := range axisY {
				(*field)[y][x] = stars
			}
		}
	}
}
func spaceDevelopment(law spaceАlgorithm, options ...propertiesOfSpace) {
	var field space
	spaceStandard(&field)
	for _, prOfSpace := range options {
		prOfSpace(&field)
	}
	law(&field)
	showSpace(&field)
}
func hourglass() spaceАlgorithm {
	return func(sand *space) {
		for y, axisY := range *sand {
			if y == 0 || y == len(*sand)-1 {
				continue
			}
			for x := range axisY {
				if x == y || x+y == len(*sand)-1 {
					continue
				}
				(*sand)[y][x] = " "
			}
		}
	}
}
func showSpace(field *space) {
	for _, axisY := range *field {
		for _, particle := range axisY {
			fmt.Printf("%s", particle)
		}
		fmt.Println()
	}
}
func main() {
	spaceDevelopment(hourglass())
	spaceDevelopment(hourglass(), setStars("0"))
	spaceDevelopment(hourglass(), setSize("20"), setStars("1"), setColor(White))
	spaceDevelopment(hourglass(), setColor(Blue), setColor(Purple), setColor(Gray), setSize("10"))
	spaceDevelopment(hourglass(), setSize("9"), setStars("5"), setColor(Purple), setSize("8"))
}

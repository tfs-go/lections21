package figure

import (
	"fmt"
	"strings"
)

type row []string

type Figure []row

func (f Figure) String() string { 
	var res strings.Builder
	for i := range f {
		var subStr strings.Builder
  		for j := range f[i] {
			subStr.WriteString(f[i][j])
  		}
		subStr .WriteString("\n")

		res.WriteString(subStr.String())
 	}

	return res.String()
}

type Modifier func(intArgs map[string]int, strArgs map[string]string)

type Filler func(fig Figure, intFinalArgs map[string]int, strFinalArgs map[string]string) Figure 

const (
	defaultSize = 15
	defaultColor = 34
	defaultChar = "!"
)

const (
	sizeModifierArg = "size"
	colorMogifierArg = "color"
	charModifierArg = "char"
)

func CharModifier(char string, color int) Modifier {
	return func(intArgs map[string]int, strArgs map[string]string) {
		if intArgs == nil || strArgs == nil {
			return 
		}

		intArgs[colorMogifierArg] = color
		strArgs[charModifierArg] = fmt.Sprintf("\033[%dm%s\033[0m", color, char)
	}
}

func SizeModifier(size int) Modifier {
	return func(intArgs map[string]int, strArgs map[string]string) {
		if intArgs == nil || strArgs == nil {
			return 
		}

		intArgs[sizeModifierArg] = size
	}
}

func ColorModifier(color int) Modifier {
	return func(intArgs map[string]int, strArgs map[string]string) {
		if intArgs == nil || strArgs == nil {
			return 
		}

		intArgs[colorMogifierArg] = color
	}
}

func makeBlankFigure(size int) Figure {
	blankFigure := make(Figure, size)
	for i := range blankFigure {
		blankFigure[i] = make([]string, size)
		for j := range blankFigure[i] {
			blankFigure[i][j] = " "
		}
	}

	return blankFigure
}

func SandglassFiller(fig Figure, intFinalArgs map[string]int, strFinalArgs map[string]string) Figure {
	if intFinalArgs == nil || strFinalArgs == nil {
		return fig
	}

	size := intFinalArgs[sizeModifierArg]
	char := strFinalArgs[charModifierArg]

	for i := range fig { 
		fig[0][i] = char
		fig[size-1][i] = char
		fig[i][i] = char
		fig[i][size-1-i] = char
	}
 
	return fig
}

func setDefaultValues() (map[string]int, map[string]string) {
	intArgs := map[string]int {
		sizeModifierArg: defaultSize,
		colorMogifierArg: defaultColor,
	}
	strArgs := map[string]string {
  		charModifierArg: defaultChar,
 	}

	return  intArgs, strArgs
}

func Construct(figFiller Filler, mods ...Modifier) Figure {
	intArgs, strArgs := setDefaultValues()

 	for _, mod := range mods {
  		mod(intArgs, strArgs)
 	}

	fig := makeBlankFigure(intArgs[sizeModifierArg])
	figFiller(fig, intArgs, strArgs)

	return fig
}
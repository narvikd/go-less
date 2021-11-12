package more

import (
	"strings"
)

func More(input string) {
	lines := fileToLines(input)
	l := makeLens(lines)
	l.listen()
}

func makeLens(lines []string) UI {
	l := UI{}
	l.setBuffer(lines)
	l.top = 0
	return l
}

func fileToLines(contents string) []string {
	return strings.Split(contents, "\n")
}

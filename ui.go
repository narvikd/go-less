package more

import (
	"github.com/nsf/termbox-go"
	"os"
)

type UI struct {
	buffer []string
	top    int
	bottom int
}

const defaultColor = termbox.ColorDefault

func (ui *UI) setBuffer(lines []string) {
	newBuf := make([]string, 0, len(lines))
	for _, v := range lines {
		newBuf = append(newBuf, v)
	}
	ui.buffer = newBuf
}

func (ui *UI) len() int {
	return len(ui.buffer)
}

func (ui *UI) size() int {
	return ui.top - ui.bottom
}

func (ui *UI) printToBuffer() {
	_ = termbox.Clear(defaultColor, defaultColor)

	screen := make([]string, 0, ui.len()+1)

	screen = append(screen, ui.buffer[ui.top:ui.bottom]...)
	screen = append(screen, ": ")

	for lineIndex, line := range screen {
		x := 0

		for _, chr := range line {
			termbox.SetCell(x, lineIndex, chr, defaultColor, defaultColor)
			x++
		}
	}
}

func (ui *UI) down() {
	if ui.bottom < ui.len() {
		ui.bottom = ui.bottom + 1
		ui.top = ui.top + 1
	}

	ui.printToBuffer()
	_ = termbox.Flush()
}

func (ui *UI) up() {
	if ui.top > 0 {
		ui.bottom = ui.bottom - 1
		ui.top = ui.top - 1
	}

	ui.printToBuffer()
	_ = termbox.Flush()
}

func (ui *UI) listen() {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	_, h := termbox.Size()
	ui.bottom = h - 1

	for {
		ui.printToBuffer()
		_ = termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Ch {
			case 'd':
				ui.down()
			case 'w':
				ui.up()
			case 'q':
				termbox.Close()
				os.Exit(0)
			}
		}
	}
}

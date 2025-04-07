package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Построение шахматной доски")
	w.Resize(fyne.NewSize(640, 480))

	label := widget.NewLabel("шахматная доска")

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Введите размер шахматной доски...")

	answer := widget.NewLabel("")

	btn := widget.NewButton("Создать", func() {
		n, err := strconv.ParseInt(entry.Text, 0, 64)

		if err != nil || n <= 1 {
			answer.SetText("Ошибка ввода: введите целое число больше 1")

			return

		} else {
			board := ""
			size := int(n)
			for y := 0; y < size; y++ {
				for x := 0; x < size; x++ {
					if (x+y)%2 == 0 {
						board += "◼️"
					} else {
						board += "◻️"
					}
				}

				board += "\n"
			}
			answer.SetText(board)
		}

	})

	answer.TextStyle.Monospace = true

	w.SetContent(container.NewVBox(
		label,
		entry,
		btn,
		answer,
	))

	w.ShowAndRun()
}

package main

import (
	"bytes"
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/justtaldevelops/go-hcaptcha"
	"github.com/justtaldevelops/go-hcaptcha/solver"
	"github.com/justtaldevelops/go-hcaptcha/solver/manual"
	"time"
)

func main() {
	a := app.New()
	go func() {
		for suc := false; !suc; time.Sleep(time.Second) {
			c, err := hcaptcha.NewChallenge("https://minecraftpocket-servers.com/server/41256/vote/", "e6b7bb01-42ff-4114-9245-3d2b7842ed92")
			if err != nil {
				panic(err)
			}
			err = c.Solve(&manual.ManualSolver{SolveFunc: func(category, object string, tasks []solver.Task) (answers []solver.Task) {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				w := a.NewWindow(fmt.Sprintf("%v - hCaptcha", object))
				var grid []fyne.CanvasObject
				ans := make([]solver.Task, len(tasks))
				for i := range tasks {
					iSafe := i
					image := canvas.NewImageFromReader(bytes.NewReader(tasks[iSafe].Image), tasks[iSafe].Key+".jpg")
					image.FillMode = canvas.ImageFillContain
					image.SetMinSize(fyne.Size{Width: 100, Height: 100})
					var selected bool
					openButton := widget.NewButton("", nil)
					openButton.OnTapped = func() {
						selected = !selected
						if selected {
							ans[iSafe] = tasks[iSafe]
							openButton.Text = "selected"
						} else {
							ans[iSafe].Image = nil
							openButton.Text = ""
						}
					}
					box := container.NewPadded(image, openButton)
					grid = append(grid, box)
				}
				w.SetContent(container.NewVBox(
					container.NewGridWithColumns(3, grid...),
					widget.NewButton("Confirm", func() {
						cancel()
					}),
				))
				w.Show()
				<-ctx.Done()
				w.Close()
				// filter selected tasks
				for _, a := range ans {
					if a.Image != nil {
						answers = append(answers, a)
					}
				}
				return answers
			}})
			if err != nil {
				c.Logger().Debug(err)
			} else {
				c.Logger().Info(c.Token())
				suc = true
				a.Quit()
			}
		}
	}()
	w := a.NewWindow("main") // to keep the app
	go func() {
		for range time.Tick(300 * time.Millisecond) {
			w.Hide()
		}
	}()
	w.ShowAndRun()
}

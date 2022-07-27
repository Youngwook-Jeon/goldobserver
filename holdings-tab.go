package main

import (
	"fmt"
	"goldobserver/repository"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) holdingsTab() *fyne.Container {
	app.Holdings = app.getHoldingSlice()
	app.HoldingsTable = app.getHoldingsTable()
	holdingsContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, app.HoldingsTable),
	)

	return holdingsContainer
}

func (app *Config) getHoldingsTable() *widget.Table {
	t := widget.NewTable(
		func() (int, int) {
			return len(app.Holdings), len(app.Holdings[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			if tci.Col == (len(app.Holdings[0])-1) && tci.Row != 0 {
				// last cell - put in a button
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "", func(deleted bool) {
						if deleted {
							id, _ := strconv.Atoi(app.Holdings[tci.Row][0].(string))
							err := app.DB.DeleteHolding(int64(id))
							if err != nil {
								app.ErrorLog.Println(err)
							}
						}
						// refresh the holdings table
						app.refreshHoldingsTable()
					}, app.MainWindow)
				})

				w.Importance = widget.HighImportance
				co.(*fyne.Container).Objects = []fyne.CanvasObject{
					w,
				}
			} else {
				// we are just putting in textual info
				co.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(app.Holdings[tci.Row][tci.Col].(string)),
				}
			}
		},
	)

	colWidths := []float32{50, 200, 200, 200, 110}
	for i := 0; i < len(colWidths); i++ {
		t.SetColumnWidth(i, colWidths[i])
	}

	return t
}

func (app *Config) getHoldingSlice() [][]interface{} {
	var slice [][]interface{}

	holdings, err := app.currentHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	slice = append(slice, []interface{}{"ID", "Amount", "Price", "Date", "Delete?"})

	for _, x := range holdings {
		var currentRow []interface{}

		currentRow = append(currentRow, strconv.FormatInt(x.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%d toz", x.Amount))
		currentRow = append(currentRow, fmt.Sprintf("$%.2f", float32(x.PurchasePrice/100)))
		currentRow = append(currentRow, x.PurchaseDate.Format("2006-01-02"))
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}

	return slice
}

func (app *Config) currentHoldings() ([]repository.Holdings, error) {
	holdings, err := app.DB.AllHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	return holdings, nil
}

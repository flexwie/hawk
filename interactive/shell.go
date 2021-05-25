package interactive

import (
	"log"

	"github.com/epiclabs-io/winman"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"hawk.wie.gg/db"
	"hawk.wie.gg/models"
	//"hawk.wie.gg/db"
)

var (
	pages *tview.Pages
	app   *tview.Application
)

func Start() {
	app = tview.NewApplication()

	Nav()
	app.SetRoot(pages, true).EnableMouse(true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func Nav() {
	categories := tview.NewList().ShowSecondaryText(false)
	categories.SetBorder(true).SetTitle("Categories")

	table := tview.NewTable().SetBorders(true)
	table.SetBorder(true).SetTitle("Entries")
	table.SetCell(0, 0, &tview.TableCell{Text: "Name", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	table.SetCell(0, 1, &tview.TableCell{Text: "Start", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	table.SetCell(0, 2, &tview.TableCell{Text: "End", Align: tview.AlignCenter, Color: tcell.ColorYellow})

	catButton := tview.NewButton("New Category").SetSelectedFunc(func() {
		app.Stop()
	})
	catButton.SetBorder(true).SetRect(0, 0, 22, 3)

	flex := tview.NewFlex().AddItem(categories, 0, 1, true).AddItem(table, 0, 1, false)

	cats, err := db.GetAllCategories()
	if err != nil {
		cobra.CheckErr(err)
	}

	for _, c := range cats {
		categories.AddItem(c.Name, "", 0, func() {
			//OnSelect(table, &c, categories)
		})
	}

	categories.SetChangedFunc(func(i int, tableName string, t string, s rune) {
		category, err := db.GetCategoryByName(tableName)
		if err != nil {
			cobra.CheckErr(err)
		}

		OnSelect(table, &category, categories)
	})

	categories.SetCurrentItem(0)
	categories.SetDoneFunc(func() {
		app.Stop()
	})
	categories.SetInputCapture(func(key *tcell.EventKey) *tcell.EventKey {
		if key.Key() == tcell.KeyEnter {
			app.SetFocus(table)
			table.SetSelectable(true, true)
		}

		return key
	})

	pages = tview.NewPages().AddPage("Main", flex, true, true)

}

func OnSelect(entries *tview.Table, category *models.Category, categories *tview.List) {
	entries.Clear()

	entries.SetCell(0, 0, &tview.TableCell{Text: "Name", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	entries.SetCell(0, 1, &tview.TableCell{Text: "Start", Align: tview.AlignCenter, Color: tcell.ColorYellow})
	entries.SetCell(0, 2, &tview.TableCell{Text: "End", Align: tview.AlignCenter, Color: tcell.ColorYellow})

	for i, e := range category.Entries {
		entries.SetCell(i+1, 0, &tview.TableCell{Text: e.Name, Color: tcell.ColorWhite})
		entries.SetCell(i+1, 1, &tview.TableCell{Text: e.Start, Color: tcell.ColorWhite})
		entries.SetCell(i+1, 2, &tview.TableCell{Text: e.End, Color: tcell.ColorWhite})
	}

	entries.Select(1, 0).SetFixed(1, 0).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			entries.Clear()
			app.SetFocus(categories)
		}

		if key == tcell.KeyEnter {
			app.SetFocus(entries)
			entries.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row, column int) {
		entries.GetCell(row, column).SetTextColor(tcell.ColorRed)
		Edit(entries.GetCell(0, column).Text, entries.GetCell(row, column).Text, category.Id)
		entries.SetSelectable(false, false)
	})

	//app.SetFocus(entries)
}

func Edit(label string, value string, id string) {
	var newValue string

	wm := winman.NewWindowManager()
	content := tview.NewForm().AddInputField(label, value, 20, nil, func(text string) {

		newValue = text
		log.Println(newValue)
	}).AddButton("Save", func() {
		log.Println(newValue)
		db.UpdateEntry(label, newValue, id)
		pages.SwitchToPage("Main")
	}).AddButton("Quit", func() {
		pages.SwitchToPage("Main")
	})

	window := wm.NewWindow().Show().SetRoot(content).SetDraggable(true).SetTitle("Window")
	window.SetRect(5, 5, 30, 10)

	pages.AddAndSwitchToPage("entry", wm, true)
}

// This package handles initializing TUI widgets and applying listeners to them.
package tui

import (
	"fmt"

	"github.com/dcap0/EZ-weeb-G/pkg/logic"
	"github.com/dcap0/EZ-weeb-G/pkg/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var yotsubatoColor = tcell.NewRGBColor(206, 230, 110)
var yotsubatoCompliment = tcell.NewRGBColor(134, 110, 230)

// InitUI takes a slice of Series structs,
// builds UI around the data. and formats it.
func InitUI(seriesData []models.Series) {
	//Initialize Widgets
	app := tview.NewApplication()
	showList := showListInit(seriesData)
	descriptionText := descriptionTextInit()
	downloadList := downloadListInit()
	controlsView := controlsViewInit()

	//Set startup content
	descriptionText.Clear().SetText(seriesData[showList.GetCurrentItem()].Description)
	controlsView.SetText("(Enter) Get Torrent Links\t\t(1) Titles\t(2) Description\t(3) Torrent Links\t(q) Quit")

	//Set up Layout
	sideFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(descriptionText, 0, 2, false).
		AddItem(downloadList, 0, 3, false)

	contentFlex := tview.NewFlex().
		AddItem(showList, 0, 2, true).
		AddItem(sideFlex, 0, 4, false)

	topFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(contentFlex, 0, 100, true).
		AddItem(controlsView, 0, 1, false)

	//listeners
	showList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		//update description based on selected title.
		descriptionText.Clear().SetText(seriesData[index].Description)
	})

	showList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// if enter is pressed on a title, it will clear the list
		// and populate it with available torrent filenames
		// stes focus on the download list
		if event.Rune() == 13 { //enter key
			populateDownloadList(
				downloadList,
				logic.GetSeriesDownloadLink(seriesData[showList.GetCurrentItem()].Title),
			)
			app.SetFocus(downloadList)
		}

		return event
	})

	downloadList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 13 { //enter key
			//magnet link
			selectedTorrent, _ := downloadList.GetItemText(downloadList.GetCurrentItem())
			downloadMap := logic.GetSeriesDownloadLink(seriesData[showList.GetCurrentItem()].Title)
			err := logic.OpenMagnet(downloadMap[selectedTorrent])
			fmt.Printf("err.Error(): %v\n", err.Error())
			if err != nil {
				modal := messageModalInit("An error occurred while opening the magnet link:\n" + err.Error()).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonLabel == "OK" {
							app.SetRoot(topFlex, false).SetFocus(showList)
						}
					})

				app.SetRoot(modal, false).SetFocus(modal)
			} else {

				modal := messageModalInit("Opening Magnet Link:\n" + selectedTorrent + ":" + err.Error()).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonLabel == "OK" {
							app.SetRoot(topFlex, false).SetFocus(showList)
						}
					})

				app.SetRoot(modal, false).SetFocus(modal)
			}
		}
		return event
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch input := event.Rune(); input {
		case 113: //q
			app.Stop()
		case 49: //1
			app.SetFocus(showList)
			controlsView.Clear().SetText("(Enter) Get Torrent Links\t\t(1) Titles\t(2) Description\t(3) Torrent Links\t(q) Quit")
		case 50: //2
			app.SetFocus(descriptionText)
			controlsView.Clear().SetText("(1) Titles\t(2) Description\t(3) Torrent Links\t(q) Quit")
		case 51: //3
			app.SetFocus(downloadList)
			controlsView.Clear().SetText("(Enter) Download Selected Torrent\t\t(1) Titles\t(2) Description\t(3) Torrent Links\t(q) Quit")
		}
		return event
	})

	if err := app.SetRoot(topFlex, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}

// descriptionTextInit returns a textview widget that has the application styles applied.
func descriptionTextInit() *tview.TextView {
	descriptionText := tview.
		NewTextView().
		SetTextColor(yotsubatoColor)

	descriptionText.
		SetBackgroundColor(yotsubatoCompliment).
		SetBorderPadding(2, 2, 2, 2).
		SetBorder(true)

	return descriptionText
}

// showListInit returns a list widget that has the application styles applied.
// it populates the list with relevent series titles.
func showListInit(seriesData []models.Series) *tview.List {
	showList := tview.
		NewList().
		ShowSecondaryText(false).
		SetMainTextColor(yotsubatoColor)

	showList.
		SetBackgroundColor(yotsubatoCompliment).
		SetBorderPadding(2, 2, 2, 2).
		SetBorder(true)

	for _, show := range seriesData {
		showList.AddItem(show.Title, "", rune(0), nil).
			SetShortcutColor(yotsubatoColor)
	}

	return showList
}

// downloadListInit returns a list widget with application styles applied.
func downloadListInit() *tview.List {
	downloadList := tview.
		NewList().
		ShowSecondaryText(false).
		SetMainTextColor(yotsubatoColor)

	downloadList.SetBackgroundColor(yotsubatoCompliment).
		SetBorderPadding(2, 2, 2, 2).
		SetBorder(true)

	return downloadList
}

// populateDownloadList takes the downloadList widget, clears it,
// and populates it with torrent link names stored in downloadLinks
func populateDownloadList(downloadList *tview.List, downloadLinks map[string]string) {
	downloadList.Clear()

	keys := make([]string, 0, len(downloadLinks))
	for k := range downloadLinks {
		keys = append(keys, k)
	}

	for _, linkTitle := range keys {
		downloadList.AddItem(linkTitle, "", rune(0), nil).
			SetShortcutColor(yotsubatoColor)
	}
}

// controlsViewInit returns a textview widget with application styles applied.
func controlsViewInit() *tview.TextView {
	controlsView := tview.NewTextView()

	controlsView.
		SetTextColor(yotsubatoCompliment).
		SetBackgroundColor(yotsubatoColor)

	return controlsView
}

// messageModalInit returns a modal widget with the provided text.
func messageModalInit(textContent string) *tview.Modal {
	messageModal := tview.NewModal()

	messageModal.SetBackgroundColor(yotsubatoColor).SetTextColor(yotsubatoCompliment)
	messageModal.SetText(textContent)
	messageModal.AddButtons([]string{"OK"})

	return messageModal
}

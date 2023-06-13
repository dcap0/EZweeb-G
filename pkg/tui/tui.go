package tui

import (
	"github.com/dcap0/EZ-weeb-G/pkg/logic"
	"github.com/dcap0/EZ-weeb-G/pkg/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var yotsubatoColor = tcell.NewRGBColor(206, 230, 110)
var yotsubatoCompliment = tcell.NewRGBColor(134, 110, 230)

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
		descriptionText.Clear().SetText(seriesData[index].Description)
	})

	showList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
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
			if err != nil {
				//make this a modal!
				panic(err)
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
		showList.AddItem(show.Title, "", rune(0), nil).SetShortcutColor(yotsubatoColor)
	}

	return showList
}

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

func populateDownloadList(downloadList *tview.List, downloadLinks map[string]string) {
	keys := make([]string, 0, len(downloadLinks))
	for k := range downloadLinks {
		keys = append(keys, k)
	}

	for _, linkTitle := range keys {
		downloadList.AddItem(linkTitle, "", rune(0), nil)
	}
}

func controlsViewInit() *tview.TextView {
	controlsView := tview.NewTextView()

	controlsView.
		SetTextColor(yotsubatoCompliment).
		SetBackgroundColor(yotsubatoColor)

	return controlsView
}

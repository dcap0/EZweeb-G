// This package handles initializing TUI widgets and applying listeners to them.
package tui

import (
	"fmt"
	"sort"

	"github.com/dcap0/EZ-weeb-G/pkg/data"
	"github.com/dcap0/EZ-weeb-G/pkg/logic"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var currentOptions = data.Options{"1080p", "eng"}
var yotsubatoColor = tcell.NewRGBColor(206, 230, 110)
var yotsubatoCompliment = tcell.NewRGBColor(134, 110, 230)

const stringUpArrow string = string(rune(8593))
const stringDownArrow string = string(rune(8595))
const baseDirections = "(1) Titles\t(2) Description\t(3) Torrent Links\t(q) Quit\t(%s %s) Navigate\t\t"
const showListDirections string = baseDirections + "(Enter) Get Torrent Links"
const downloadListDirections string = baseDirections + "(Enter) Download Selected Torrent"

// InitUI takes a slice of Series structs,
// builds UI around the data. and formats it.
func InitUI(seriesData []data.Series) {
	//Initialize Widgets
	app := tview.NewApplication()
	showList := showListInit(seriesData)
	descriptionText := descriptionTextInit()
	downloadList := downloadListInit()
	controlsView := controlsViewInit()

	optionsTree := qualityDropdownInit()
	optionsTree.SetTitle("take 2").SetBorder(true)

	optionsMenu := optionsMenuInit(optionsTree.Box, 40, 10)
	pages := tview.NewPages()

	//Set startup content
	descriptionText.Clear().SetText(seriesData[showList.GetCurrentItem()].Description)
	controlsView.SetText(fmt.Sprintf(showListDirections, stringUpArrow, stringDownArrow))

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

	showList.
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
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
		}).
		SetFocusFunc(func() {
			controlsView.SetText(fmt.Sprintf(showListDirections, stringUpArrow, stringDownArrow))
		})

	downloadList.
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 13 { //enter key
				//magnet link
				selectedTorrent, _ := downloadList.GetItemText(downloadList.GetCurrentItem())
				downloadMap := logic.GetSeriesDownloadLink(seriesData[showList.GetCurrentItem()].Title)
				err := logic.OpenMagnet(downloadMap[selectedTorrent])
				if err != nil {
					modal := messageModalInit("An error occurred while opening the magnet link:\n" + err.Error()).
						SetDoneFunc(func(buttonIndex int, buttonLabel string) {
							if buttonLabel == "OK" {
								app.SetRoot(topFlex, false).SetFocus(showList)
							}
						})

					app.SetRoot(modal, false).SetFocus(modal)
				} else {
					modal := messageModalInit("Opening Magnet Link:\n" + selectedTorrent).
						SetDoneFunc(func(buttonIndex int, buttonLabel string) {
							if buttonLabel == "OK" {
								app.SetRoot(topFlex, false).SetFocus(showList)
							}
						})

					app.SetRoot(modal, false).SetFocus(modal)
				}
			}
			return event
		}).
		SetFocusFunc(func() {
			controlsView.SetText(fmt.Sprintf(downloadListDirections, stringUpArrow, stringDownArrow))
		})

	descriptionText.SetFocusFunc(func() {
		controlsView.Clear().SetText(fmt.Sprintf(baseDirections, stringUpArrow, stringDownArrow))
	})

	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		front, _ := pages.GetFrontPage()

		switch input := event.Rune(); input {
		case 52: //4
			if front == "main" {
				pages.SendToFront("options")
			}
			app.SetFocus(optionsMenu)
		case 113: //q
			app.Stop()
		case 49: //1
			if front == "options" {
				pages.SendToBack("options")
			}
			app.SetFocus(showList)
		case 50: //2
			if front == "options" {
				pages.SendToBack("options")
			}
			app.SetFocus(descriptionText)
		case 51: //3
			if front == "options" {
				pages.SendToBack("options")
			}
			app.SetFocus(downloadList)
		}
		return event
	})

	pages.AddPage("options", optionsMenu, true, true)
	pages.AddPage("main", topFlex, true, true)

	if err := app.SetRoot(pages, true).EnableMouse(false).Run(); err != nil {
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
func showListInit(seriesData []data.Series) *tview.List {
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

	sort.Strings(keys)

	if len(keys) == 0 {
		keys = append(keys, "No torrents found on Nyaa")
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

func optionsMenuInit(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

func qualityDropdownInit() *tview.DropDown {
	qualityDropDown := tview.NewDropDown()
	availableQualities := []string{"480p", "720p", "1080p"}
	qualityDropDown.
		SetOptions(
			availableQualities,
			nil,
		).
		// SetCurrentOption(availableQualities).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 13 {
				_, quality := qualityDropDown.GetCurrentOption()
				currentOptions.Quality = quality
			}

			return event
		})

	return qualityDropDown
}

// func qualityOptionsNode() {

// }

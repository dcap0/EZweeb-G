// This package handles initializing TUI widgets and applying listeners to them.
package tui

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/dcap0/EZ-weeb-G/pkg/data"
	"github.com/dcap0/EZ-weeb-G/pkg/logic"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var currentOptions = data.Options{Year: "", Season: "", Quality: "", SubLang: "eng"}
var yotsubatoColor = tcell.NewRGBColor(206, 230, 110)
var yotsubatoCompliment = tcell.NewRGBColor(134, 110, 230)

const stringUpArrow string = string(rune(8593))
const stringDownArrow string = string(rune(8595))
const baseDirections = "(1) Titles\t(2) Description\t(3) Torrent Links\t(q) Quit\t(%s %s) Navigate\t(o) Options\t(r) Refresh\t"
const optionsDirections = "(q) Quit\t(Tab) Navigate Down\t(Shift+Tab) Navigate Up\t(Enter) Select Item\t(o) Exit Options"
const showListDirections string = baseDirections + "(Enter) Get Torrent Links"
const downloadListDirections string = baseDirections + "(Enter) Download Selected Torrent"
const optionsReadoutData string = "Selected Year: %s Selected Season: %s Selected Quality %s"

var seriesData []data.Series = logic.GetSeriesHtml(currentOptions.Year, currentOptions.Season)

var app *tview.Application

// InitUI takes a slice of Series structs,s
// builds UI around the data. and formats it.
func InitUI() {
	//Get Series Data

	//Initialize main Widgets
	app = tview.NewApplication()
	showList := showListInit()
	descriptionText := descriptionTextInit()
	downloadList := downloadListInit()
	controlsView := infoViewInit()
	optionsReadout := infoViewInit()

	//Initialize options Widgets
	optionsForm := initOptionsForm()

	pages := tview.NewPages()

	//Set startup content
	descriptionText.Clear().SetText(seriesData[showList.GetCurrentItem()].Description)
	controlsView.Clear().SetText(fmt.Sprintf(showListDirections, stringUpArrow, stringDownArrow))
	optionsReadout.Clear().SetText(fmt.Sprintf(optionsReadoutData, currentOptions.Year, currentOptions.Season, currentOptions.Quality))

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
		AddItem(optionsReadout, 0, 1, false).
		AddItem(pages, 0, 100, true).
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
					logic.GetSeriesDownloadLink(seriesData[showList.GetCurrentItem()].Title, currentOptions.Quality, ""),
				)
				app.SetFocus(downloadList)
			}

			return event
		}).
		SetFocusFunc(func() {
			controlsView.Clear().SetText(fmt.Sprintf(showListDirections, stringUpArrow, stringDownArrow))
		})

	downloadList.
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 13 { //enter key
				//magnet link
				selectedTorrent, _ := downloadList.GetItemText(downloadList.GetCurrentItem())
				downloadMap := logic.GetSeriesDownloadLink(seriesData[showList.GetCurrentItem()].Title, currentOptions.Quality, "")
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
			controlsView.Clear().SetText(fmt.Sprintf(downloadListDirections, stringUpArrow, stringDownArrow))
		})

	descriptionText.SetFocusFunc(func() {
		controlsView.Clear().SetText(fmt.Sprintf(baseDirections, stringUpArrow, stringDownArrow))
	})

	contentFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch input := event.Rune(); input {
		case 49: //1
			app.SetFocus(showList)
		case 50: //2
			app.SetFocus(descriptionText)
		case 51: //3
			app.SetFocus(downloadList)
		case 114:
			descriptionText.Clear()
			downloadList.Clear()
			populateShowList(showList, seriesData)
		}
		return event
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		front, _ := pages.GetFrontPage()

		switch input := event.Rune(); input {
		case 111: //o
			if front == "main" {
				pages.SendToFront("options")
				controlsView.Clear().SetText(optionsDirections)
			} else {
				pages.SendToFront("main")
				app.SetFocus(showList)
				controlsView.Clear().SetText(fmt.Sprintf(showListDirections, stringUpArrow, stringDownArrow))
				populateShowList(showList, seriesData)
				//apply options
			}
		case 113: //q
			app.Stop()
		}
		return event
	})

	pages.AddPage("options", optionsForm, true, true)
	pages.AddPage("main", contentFlex, true, true)

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
func showListInit() *tview.List {
	showList := tview.
		NewList().
		ShowSecondaryText(false).
		SetMainTextColor(yotsubatoColor)

	showList.
		SetBackgroundColor(yotsubatoCompliment).
		SetBorderPadding(2, 2, 2, 2).
		SetBorder(true)

	populateShowList(showList, seriesData)

	return showList
}

func populateShowList(showList *tview.List, seriesData []data.Series) {
	showList.Clear()
	for _, show := range seriesData {
		showList.AddItem(show.Title, "", rune(0), nil).
			SetShortcutColor(yotsubatoColor)
	}
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
func infoViewInit() *tview.TextView {
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

func initOptionsForm() *tview.Form {
	optionsForm := tview.NewForm()
	optionsForm.
		SetFieldBackgroundColor(yotsubatoCompliment).
		SetFieldTextColor(yotsubatoColor).
		SetButtonBackgroundColor(yotsubatoCompliment).
		SetTitleColor(yotsubatoCompliment).
		SetBorder(true).
		SetBorderColor(yotsubatoColor).
		SetTitle("Options")
	optionsForm.
		AddDropDown("Quality", []string{"", "1080p", "720p", "480p"}, 0, nil).
		AddDropDown("Season", []string{"", "winter", "spring", "summer", "fall"}, 0, nil).
		AddInputField(
			"Year",
			currentOptions.Year,
			20,
			func(textToCheck string, lastChar rune) bool {
				_, err := strconv.Atoi(textToCheck)
				return err == nil
			},
			nil,
		).
		AddButton(
			"Set to Current Season", func() {
				currentOptions.Year = ""
				currentOptions.Season = ""
				optionsForm.GetFormItemByLabel("Year").(*tview.InputField).SetText(currentOptions.Year)
				optionsForm.GetFormItemByLabel("Season").(*tview.DropDown).SetCurrentOption(0)
			},
		).
		AddButton("Save", func() {
			currentOptions.Year = optionsForm.GetFormItemByLabel("Year").(*tview.InputField).GetText()
			_, season := optionsForm.GetFormItemByLabel("Season").(*tview.DropDown).GetCurrentOption()
			currentOptions.Season = season
			_, quality := optionsForm.GetFormItemByLabel("Quality").(*tview.DropDown).GetCurrentOption()
			currentOptions.Quality = quality
			seriesData = logic.GetSeriesHtml(currentOptions.Year, currentOptions.Season)
			optionsForm.GetFormItemByLabel("Last Saved:").(*tview.TextView).SetText(time.Now().Format(time.RFC3339))
		}).
		AddTextView("Last Saved:", "", 0, 1, false, false).
		SetFocus(3)

	return optionsForm
}

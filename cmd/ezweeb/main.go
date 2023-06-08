package main

import (
	"fmt"

	"github.com/dcap0/EZ-weeb-G/pkg/logic"
	"github.com/rivo/tview"
)

func main() {

	var app = tview.NewApplication()

	err := app.SetRoot(tview.NewBox(), true).EnableMouse(false).Run()

	if err != nil {
		panic(err)
	}

	seriesData := logic.GetSeriesHtml()

	//choose a series via CLI

	downloadMap := logic.GetSeriesDownloadLink(seriesData[0].Title)

	keys := make([]string, 0, len(downloadMap))
	for k := range downloadMap {
		keys = append(keys, k)
	}

	for _, s := range keys {
		fmt.Printf("show: %v\n", s)
		fmt.Printf("link: %v\n", downloadMap[s])

	}

	//Choose a link via CLI

}

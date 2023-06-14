package main

import (
	"github.com/dcap0/EZ-weeb-G/pkg/logic"
)

// main starts the program.
// initializes the TUI and provides it the series data
func main() {
	// tui.InitUI(logic.GetSeriesHtml())
	s := logic.GetSeriesHtml()

	logic.GetSeriesDownloadLink(s[0].Title)
}

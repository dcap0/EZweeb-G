package main

import (
	"github.com/dcap0/EZ-weeb-G/pkg/logic"
	"github.com/dcap0/EZ-weeb-G/pkg/tui"
)

// main starts the program.
// initializes the TUI and provides it the series data
func main() {
	tui.InitUI(logic.GetSeriesHtml())
}

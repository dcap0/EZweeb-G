package main

import (
	"github.com/dcap0/EZ-weeb-G/pkg/logic"
	"github.com/dcap0/EZ-weeb-G/pkg/tui"
)

func main() {
	tui.InitUI(logic.GetSeriesHtml())
}

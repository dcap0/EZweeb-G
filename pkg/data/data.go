// This package provides structs for options data and show data
// for a series.
package data

// Holds data for a series
type Series struct {
	Title, Description string
}

// Holds user options
type Options struct {
	Year, Season, Quality, SubLang string
	SafetyLevel                    Safety
}

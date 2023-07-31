// This package implements utility functions for
// getting and handling data related to currently airing anime seasons.
//
// The package is used to get and parse HTML across several anime sites,
// and to open associated magnet links.
package logic

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dcap0/EZ-weeb-G/pkg/data"
)

// GetSeriesHtml sends a request to MAL to get the current season page.
// It takes a year and season.
// HTML is parsed to pull series titles and descriptions.
// Returns a Slice of type models.Series struct.
func GetSeriesHtml(year, season string) []data.Series {
	seriesData := make([]data.Series, 0)

	queryString := "https://myanimelist.net/anime/season"

	if year != "" && season != "" {
		queryString += "/" + year + "/" + season
	}

	resp, err := http.Get(queryString)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	titles := make([]string, 0)
	descriptions := make([]string, 0)

	doc.Find(".link-title").Each(func(i int, s *goquery.Selection) {
		titles = append(titles, s.Text())
	})

	doc.Find(".preline").Each(func(i int, s *goquery.Selection) {
		descriptions = append(descriptions, s.Text())
	})

	for i := 0; i < len(titles); i++ {
		seriesData = append(seriesData, data.Series{Title: titles[i], Description: descriptions[i]})
	}

	return seriesData
}

// GetSeriesDownloadLink sends a request to nyaa.si with a title to query, as well as video quality and subtitle language.
// HTML is parsed to pull "successful" torrents as well as their associated magnet links.
// Returns a map of torrent file name to magnet link.
func GetSeriesDownloadLink(title, quality, subtitle string, safety data.Safety) map[string]string {
	retVal := make(map[string]string)
	const queryUri string = "https://nyaa.si/?f=0&c=0_0&q="
	queryTitle := strings.ReplaceAll(title, " ", "+")

	if quality != "" {
		queryTitle += "+" + quality
	}

	if subtitle != "" {
		queryTitle += "+" + subtitle
	}

	resp, err := http.Get(queryUri + queryTitle)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	failCounter := 0

	var htmlSearchParam string

	switch safety {
	case data.TRUSTED:
		htmlSearchParam = ".success"
	case data.DANGER:
		htmlSearchParam = ".danger"
	case data.DEFAULT:
		htmlSearchParam = ".default"
	case data.ALL:
		htmlSearchParam = "tr"
	default:
		htmlSearchParam = ".success"
	}

	//get success files
	doc.Find(htmlSearchParam).Each(func(i int, s *goquery.Selection) {
		var fileName string

		//Element holding filename is always the secondTD (colspan=2)
		fileNameElement := s.Find("td").First().Next().Find("a")

		//For whatever reason, nyaa puts the comment right next to the filename, instead of its own box.
		switch fileNameElement.Length() {
		case 1:
			fileName = fileNameElement.First().Text()
		case 2:
			fileName = fileNameElement.First().Next().Text()
		default:
			failCounter++
			fileName = "Failed to Find Filename: " + fmt.Sprint(failCounter)
		}

		//magnetLinks have a child
		magnetLink, magnetLinkExists := s.Find(".fa-magnet").Parent().Attr("href")

		if magnetLinkExists {
			retVal[strings.TrimSpace(fileName)] = magnetLink
		} else {
			retVal[strings.TrimSpace(fileName)] = ""
		}

	})
	return retVal
}

// OpenMagnet takes a string representation of a magnet link and opens it with the OS's default handler.
// Returns an error if there's an issue opening the file or if platform is not supported.
// Otherwise returns nil
func OpenMagnet(magnetLink string) error {
	var err error

	if magnetLink == "" {
		return errors.New("no magnet link found")
	}

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", magnetLink).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", magnetLink).Start()
	case "darwin":
		err = exec.Command("open", magnetLink).Start()
	default:
		err = errors.New("unsupported platform")
	}

	return err
}

// func IndexOf(arr interface{}, value any) (int, error) {

// 	arrType := arr.([]string)
// 	valueType := reflect.TypeOf(value)

// 	if arrType != valueType {
// 		return 0, fmt.Errorf("type mismatch: arr is %v, value is %v", arrType, valueType)
// 	}

// 	if len(arr) == 0 {
// 		return 0, errors.New("empty array")
// 	}

// 	return 10, nil
// }

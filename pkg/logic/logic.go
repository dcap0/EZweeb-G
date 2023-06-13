package logic

import (
	"errors"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dcap0/EZ-weeb-G/pkg/models"
)

// GetSeriesHtml sends a request to MAL to get the current season page.
// HTML is parsed to pull series titles and descriptions.
// It returns a Slice of type models.Series struct
func GetSeriesHtml() []models.Series {
	seriesData := make([]models.Series, 0)

	resp, err := http.Get("https://myanimelist.net/anime/season")

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
		seriesData = append(seriesData, models.Series{Title: titles[i], Description: descriptions[i]})
	}

	return seriesData
}

func GetSeriesDownloadLink(title string) map[string]string {
	retVal := make(map[string]string)
	const queryUri string = "https://nyaa.si/?f=0&c=0_0&q="
	queryTitle := strings.ReplaceAll(title, " ", "+")
	queryTitle += "+sub"
	resp, err := http.Get(queryUri + queryTitle)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".success").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find(".text-center").Find("a").Siblings().Attr("href")
		retVal[strings.TrimSpace(s.Find("a").Text())] = href
	})

	return retVal
}

func OpenMagnet(magnetLink string) error {
	var err error = nil
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", magnetLink).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", magnetLink).Start()
	case "darwin":
		err = exec.Command("open", magnetLink).Start()
	default:
		err = errors.New("Unsupported Platform")
	}

	return err
}

package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	s "bestia/structs"

	"github.com/gocolly/colly"
)

var REGEX_YEAR *regexp.Regexp = regexp.MustCompile(`^[0-9]+$`)

func Initializate() *colly.Collector {
	var scrapper *colly.Collector = colly.NewCollector()

	scrapper.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	scrapper.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	return scrapper
}

func FindCommonMovies(movies [][]s.Movie, maxCount int) []s.Movie {
	var moviesCount map[s.Movie]int = make(map[s.Movie]int)

	for _, movieSet := range movies {
		for _, movie := range movieSet {
			moviesCount[movie]++
		}
	}

	var commonMovies []s.Movie = make([]s.Movie, 0)
	for movie, count := range moviesCount {
		if count >= maxCount {
			commonMovies = append(commonMovies, movie)
		}
	}

	return commonMovies
}

func GetMaxPages(c *colly.Collector, url string) int {
	var maxPages int = 1
	c.OnHTML("div.pagination", func(e *colly.HTMLElement) {
		maxPages, _ = strconv.Atoi(e.DOM.Find("li").Last().Text())
	})

	c.Visit(url)

	return maxPages
}

func GetTitleAndYear(movieInfo string) (string, string) {
	var resultYear string = ""

	var titleYear []string = strings.Split(movieInfo, "-")
	var lastPosition int = len(titleYear) - 1
	var possibleYear string = titleYear[lastPosition]

	if len(titleYear) > 1 && REGEX_YEAR.MatchString(possibleYear) {
		resultYear = titleYear[lastPosition]
		titleYear = titleYear[:lastPosition]
	}

	return strings.Join(titleYear, " "), resultYear
}

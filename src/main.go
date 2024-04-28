package main

import (
	"fmt"
	"os"
	"strings"

	s "bestia/structs"
	"bestia/utils"

	"github.com/gocolly/colly"
)

const BASE_URL = "https://letterboxd.com/"

func main() {
	var users []string = os.Args[1:]
	var globalResult [][]s.Movie

	// var users []string = []string{"daniel_alba", "malejaa"}
	c := utils.Initializate()

	for _, user := range users {
		var page int = 1
		var results []s.Movie = []s.Movie{}
		var url string = fmt.Sprintf("%s%s/watchlist/", BASE_URL, user)
		var max_pages int = utils.GetMaxPages(c, url)

		c.OnHTML("div.really-lazy-load", func(e *colly.HTMLElement) {
			movie := s.Movie{}
			movieInfo := e.Attr("data-film-slug")

			movie.Title, movie.Year = utils.GetTitleAndYear(movieInfo)
			results = append(results, movie)
		})

		for {
			if page > max_pages {
				break
			}

			var url string = fmt.Sprintf("%s%s/watchlist/page/%d", BASE_URL, user, page)
			c.Visit(url)
			page++
		}

		globalResult = append(globalResult, results)

	}

	var commonMovies []s.Movie = utils.FindCommonMovies(globalResult, len(users))
	fmt.Println("Total movies in common", len(commonMovies))
	for _, movie := range commonMovies {
		fmt.Println(strings.ToTitle(movie.Title))
	}

}

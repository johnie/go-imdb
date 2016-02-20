package main

import (
  "github.com/gin-gonic/gin"
  "github.com/ernesto-jimenez/scraperboard"
)

func Top(c *gin.Context) {
  scraper, _ := scraperboard.NewScraperFromString(topScraper)

  var topresponse TopResponse

  scraper.ExtractFromURL("http://www.imdb.com/chart/top", &topresponse)

  c.JSON(200, topresponse)

}

type TopResponse struct {
  TopResults []topresult
}

type topresult struct {
  ID string
  Title string
  Year string
  Rating string
  Url string
}

var topScraper = `
  <Scraper>
    <Each name="topresults" selector=".lister-list > tr">
      <Property name="id" selector=".titleColumn a">
        <Filter type="first"/>
        <Filter type="attr" argument="href"/>
        <Filter type="regex" argument="(tt[\d]{7})"/>
      </Property>
      <Property name="title" selector=".titleColumn a" />
      <Property name="year" selector=".titleColumn span" />
      <Property name="rating" selector=".ratingColumn strong" />
      <Property name="url" selector=".titleColumn a">
				<Filter type="first"/>
				<Filter type="attr" argument="href"/>
        <Filter type="regex" argument="(\/title\/tt[\d]{7}\/)"/>
			</Property>
    </Each>
  </Scraper>
`

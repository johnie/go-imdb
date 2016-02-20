package main

import (
  "github.com/gin-gonic/gin"
  "github.com/ernesto-jimenez/scraperboard"
)

func Top(c *gin.Context) {
  scraper, _ := scraperboard.NewScraperFromFile("scraper_files/top.xml")

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

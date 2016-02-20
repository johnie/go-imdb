package main

import (
  "github.com/gin-gonic/gin"
  "github.com/ernesto-jimenez/scraperboard"
)

func TopTV(c *gin.Context) {
  scraper, _ := scraperboard.NewScraperFromFile("scraper_files/toptv.xml")

  var toptvresponse TopTvResponse

  scraper.ExtractFromURL("http://www.imdb.com/chart/toptv", &toptvresponse)

  c.JSON(200, toptvresponse)
}

type TopTvResponse struct {
  TopTvResults []toptvresult
}

type toptvresult struct {
  ID string
  Title string
  Year string
  Rating string
  Url string
}

package main

import (
  "github.com/martini-contrib/render"
  "github.com/ernesto-jimenez/scraperboard"
)

func Top(r render.Render) {
  scraper, _ := scraperboard.NewScraperFromFile("scraper_files/top.xml")

  var topresponse TopResponse

  scraper.ExtractFromURL("http://www.imdb.com/chart/top", &topresponse)

  r.JSON(200, topresponse)
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

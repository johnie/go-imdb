package main

import (
  "github.com/martini-contrib/render"
  "github.com/ernesto-jimenez/scraperboard"
)

func TopTV(r render.Render) {
  scraper, _ := scraperboard.NewScraperFromFile("scraper_files/toptv.xml")

  var toptvresponse TopTvResponse

  scraper.ExtractFromURL("http://www.imdb.com/chart/toptv", &toptvresponse)

  r.JSON(200, toptvresponse)
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

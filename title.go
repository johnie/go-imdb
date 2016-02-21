package main

import (
  "fmt"
  "strings"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/ernesto-jimenez/scraperboard"
  "github.com/leebenson/conform"
)

func Title(r render.Render, params martini.Params) {

  titleUrl := fmt.Sprintf("http://www.imdb.com/title/%s", params["id"])

  scraper, _ := scraperboard.NewScraperFromFile("scraper_files/title.xml")

  var titleresponse TitleResponse

  scraper.ExtractFromURL(titleUrl, &titleresponse)

  // IMDB sends back a lot of whitespaces, let's trim it
  conform.Strings(&titleresponse)

  // Get the full size of the poster
  poster := titleresponse.Poster
  poster = strings.Replace(poster, "_V1_UX182_CR0,0,182,268_AL_", "_V1_SX640_SY720_", -1)
  titleresponse.Poster = poster

  r.JSON(200, titleresponse)
}

type TitleResponse struct {
  Title       string `conform:"trim"`
  Plot        string `conform:"trim"`
  Poster      string
  Rating      string `conform:"trim"`
  RatingCount string `conform:"trim"`
  Duration    string `conform:"trim"`
  Genres      []string
  Creator     string `conform:"trim"`
  Cast        []cast
}

type cast struct {
  Url       string
  Name      string
  Character string
}

package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/cors"
)

/**
 * The urls we're gonna use
 */
var urls = []string{"/top", "/toptv", "/title/(?P<id>tt[0-9]{0,7})", "/name/(?P<id>nm[0-9]{0,8})"}

func Index(r render.Render) {
  r.JSON(200, map[string]interface{}{"api_urls": urls})
}

func PageNotFound(r render.Render) {
  r.JSON(404, map[string]interface{}{"code": 404, "message": "Page not found"})
}

func main() {
  // Run Martini with basic middlewares
  app := martini.Classic()

  // HTML, JSON and XML middleware
  app.Use(render.Renderer())

  // CORS for *
  // - GET methods
  // - Origin header
  app.Use(cors.Allow(&cors.Options{
    AllowOrigins:   []string{"*"},
    AllowMethods:   []string{"GET"},
    AllowHeaders:   []string{"Origin"},
    ExposeHeaders:  []string{"Content-Length"},
  }))

  // Group urls in `/api/v1`
  app.Group("/api/v1", func(v1 martini.Router) {
    // Index route
    v1.Get("/", Index)

    // `/top` route
    v1.Get(urls[0], Top)

    // `/toptv` route
    v1.Get(urls[1], TopTV)

    // /title/:id route
    v1.Get(urls[2], Title)

    // /name/:id route
    v1.Get(urls[3], Name)
  })

  // Handle Not Found
  app.NotFound(PageNotFound)

  // Let's go with the year IMDB was created
  app.RunOnAddr(":1990")
}

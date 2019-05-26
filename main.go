package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	nico "nicodive-api/api"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/v1/video/:videoid", func(c echo.Context) error {
		videoID := c.Param("videoid")
		videoInfo, err := nico.GetVideoInfo(videoID)
		if err != nil {
			log.Printf("[ERROR] failed to call GetVideoInfo")
			return c.JSON(http.StatusNotFound, fmt.Sprintf("%s", err))
		}

		return c.JSON(http.StatusOK, videoInfo)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

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

	e.GET("/v1/ranking/:category/:period", func(c echo.Context) error {
		category := c.Param("category")
		period := c.Param("period")
		ranking, err := nico.GetRanking(category, period)
		if err != nil {
			log.Printf("[ERROR] failed to call GetRanking")
			return c.JSON(http.StatusNotFound, fmt.Sprintf("%s", err))
		}
		return c.JSON(http.StatusOK, ranking)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

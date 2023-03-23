package handle

import (
	"fmt"
	"net/http"
	"strings"

	db "github.com/Calgorr/URL_Shortener/database"
	"github.com/Calgorr/URL_Shortener/model"
	"github.com/labstack/echo/v4"
)

func SaveUrl(c echo.Context) error {
	url := c.FormValue("url")
	if url == "" {
		return c.String(http.StatusBadRequest, "url is required")
	}
	url = strings.Replace(url, "www.", "", -1)
	if !strings.Contains(url, "http://") {
		url = "http://" + url
	}
	link := model.NewLink(url)
	err := db.AddLink(link)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.String(http.StatusOK, "Your Shortened link is "+c.Request().Host+"/"+link.Hash)
}

func Redirect(c echo.Context) error {
	var err error
	var link *model.Link
	if c.Param("hash") != "" {
		hash := c.Param("hash")
		link, err = db.GetLink(hash)
		if link.Address != "" {
			db.IncrementUsage(hash)
			err = c.Redirect(http.StatusTemporaryRedirect, link.Address)
		} else {
			err = c.String(http.StatusBadRequest, "Invalid url")
		}
	}
	return err
}

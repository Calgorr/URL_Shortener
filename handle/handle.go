package handle

import (
	"net/http"
	"strings"

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
	//add to database
	return c.String(http.StatusOK, "Your Shortened link is "+c.Request().Host+"/"+link.Hash)
}

func Redirect(c echo.Context) error {
	hash := c.Param("hash")
	if hash != "" {
		//get the link from database
		var link model.Link
		return c.Redirect(http.StatusPermanentRedirect, link.Address)
	}
	return c.String(http.StatusBadRequest, "Hash parameter can not be empty")
}

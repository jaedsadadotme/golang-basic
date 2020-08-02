package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCaseToday(c echo.Context) error {
	var url = "https://covid19.th-stat.com/api/open/today"
	resp, err := http.Get(url)
	if err != nil {
		print(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return c.JSON(200, result)
}

func main() {
	e := echo.New()
	e.GET("/casetoday", GetCaseToday)
	e.Logger.Fatal(e.Start(":1323"))
}

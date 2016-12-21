package blog

import (
    "net/http"
    "github.com/labstack/echo"
)


func init() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    http.Handle("/", e)
}

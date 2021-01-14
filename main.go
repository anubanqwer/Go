package main

import (
	"net/http"
	"strconv"

	"github.com/anubanqwer/Go/fizzbuzz"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/fizzbuzz/:number", fizzbuzzHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func fizzbuzzHandler(c echo.Context) error {
	number := c.Param("number")

	n, err := strconv.Atoi(number)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fizzbuzz.Say(n))
}

/*
go run main.go
go build -o api main.go

go mod init github.com/anubanqwer/Go
*/

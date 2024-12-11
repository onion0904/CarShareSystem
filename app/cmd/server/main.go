package main

import (
    "log"
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()


    log.Println("Starting server on :8080")
    if err := e.Start(":8080"); err != nil {
        log.Fatal(err)
    }
}

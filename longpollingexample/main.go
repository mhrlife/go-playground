package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"sync/atomic"
	"time"
)

func main() {
	e := echo.New()

	didTheMove := new(atomic.Bool)
	channel := make(chan struct{}, 1)

	e.GET("/status", func(c echo.Context) error {
		startListeningTime := time.Now() // just for info

		select {
		case <-channel:
			log.Println("Yay! opponent did the move!")
		case <-time.After(time.Minute):
			log.Println("just a timeout for avoiding goroutine leak")
		case <-c.Request().Context().Done():
			log.Println("user canceled the connection")
		}

		if didTheMove.Load() {
			return c.String(200, "now your turn, you waited for "+time.Since(startListeningTime).String())
		} else {
			return c.String(200, "waiting for your opponent's move")
		}
	})

	e.GET("/move", func(context echo.Context) error {
		didTheMove.Store(true)
		channel <- struct{}{}
		return context.String(200, "done")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

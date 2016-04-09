package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	cs := make(chan string)

	go func() {

		for {
			url := <-cs
			fmt.Println("received: ", url)

			timeout := time.Duration(5 * time.Second)
			client := http.Client{
				Timeout: timeout,
			}

			resp, errs := client.Get(url)

			if errs == nil {
				fmt.Printf("%v", resp)
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)

				if err == nil {
					fmt.Println(os.Stdout, string(body))
				}

			} else {
				fmt.Printf("%v", errs)
			}
		}
	}()

	e := echo.New()

	e.Get("/add", func(c echo.Context) error {
		url := c.QueryParam("url")

		if url == "" {
			return c.String(http.StatusOK, "!OK")
		} else {
			cs <- url
			return c.String(http.StatusOK, "OK")
		}
	})

	e.Post("/add", func(c echo.Context) error {
		return c.String(http.StatusOK, "[POST] Hello, World!")
	})

	e.Run(standard.New("localhost:1323"))
}

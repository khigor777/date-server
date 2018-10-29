package main

import (
	"github.com/khigor777/date-server/rest"
	"github.com/khigor777/date-server/services/timeservice"
	"github.com/labstack/gommon/log"
)

func main() {
	eng := web.NewEngine(&web.Services{
		TimeService: timeservice.NewTimeService(nil),
	})
	log.Fatal(eng.Run(":8080"))

}

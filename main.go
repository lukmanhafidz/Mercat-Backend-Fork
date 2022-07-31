package main

import (
	"Mercat/config"
	"Mercat/utils/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Getconfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateDB(db)

	e := echo.New()
	address := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(address))
}

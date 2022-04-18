package main

import (
	"github.com/tyetis/goapiexample/app"
	"github.com/tyetis/goapiexample/app/utils"
)

func main() {
	utils.LoadConfig()
	app.CreateServer(3000)
}

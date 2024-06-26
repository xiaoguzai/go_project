package main

import (
	"gimchat1/router"
	"gimchat1/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	r.Run(":8081")
}

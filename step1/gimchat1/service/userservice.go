package service

import (
	"fmt"
	"gimchat1/model"
	"gimchat1/utils"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	data := make([]*model.UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	c.JSON(200, gin.H{
		"message": data,
	})
}

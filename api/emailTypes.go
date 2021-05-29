package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/notification/service/mail"
)

func GetMailTypes(c *gin.Context) {
	mailTypes := make([]string, 0)
	mailTypes = append(mailTypes, mail.ContentTypeHTML)
	mailTypes = append(mailTypes, mail.ContentTypeText)

	c.JSON(0, map[string]interface{}{"types": mailTypes})
	return
}

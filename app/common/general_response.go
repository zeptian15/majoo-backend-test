package common

import "github.com/gin-gonic/gin"

func SendSuccess(c *gin.Context, httpStatus int, code int, message string, data interface{}) {
	// Send Success Response
	c.JSON(httpStatus, gin.H{
		"code":    code,
		"success": true,
		"message": message,
		"data":    data,
	})
}

func SendError(c *gin.Context, httpStatus int, code int, message string, err interface{}) {
	// Send Error Response
	c.JSON(httpStatus, gin.H{
		"code":    code,
		"success": false,
		"message": message,
		"error":   err,
	})
}

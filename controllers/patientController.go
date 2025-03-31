package controllers

import "github.com/gin-gonic/gin"

func PatientGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "pong",
	})
}

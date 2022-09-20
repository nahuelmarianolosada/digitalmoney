package controllers

import (
	_ "digitalmoney/api/model"
	accountRepo "digitalmoney/api/repository/account"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func FindAccount(c *gin.Context) {
	userId := GetUserIDFromJWT(c)
	if userId == nil {
		return
	}
	userFound, err := accountRepo.GetByUserID(*userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userFound)
}

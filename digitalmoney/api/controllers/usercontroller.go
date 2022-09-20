package controllers

import (
	"strings"
	"digitalmoney/api/middlewares/auth"
	"digitalmoney/api/model"
	"digitalmoney/api/repository"
	accountRepo "digitalmoney/api/repository/account"
	aliasRepo "digitalmoney/api/repository/alias"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

func RegisterUser(context *gin.Context) {
	var user model.User
	newDefaultAccount := model.Account{AvailableAmount: "0"}
	if err := context.ShouldBindJSON(&user); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	record := repository.DB.Create(&user)
	if record.Error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		return
	}

	newDefaultAccount.UserID = user.ID
	var validCvuAndAlias bool
	for !validCvuAndAlias {
		randomCvu := strconv.Itoa(rand.Intn(999999999999999999))
		
		alias, errAlias := aliasRepo.GetRandomAlias()
		if errAlias != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errAlias.Error})
			return
		}
	
		
		if _, errCheck := accountRepo.GetByAliasCvu(*alias, randomCvu); errCheck != nil && strings.Contains(errCheck.Error(), "account not found") {
			newDefaultAccount.Cvu = &randomCvu
			newDefaultAccount.Alias = alias
			validCvuAndAlias = true
		} else if errCheck != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errCheck.Error()})
			return
		}
	}

	accountCreated, err := accountRepo.Create(newDefaultAccount)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user_id": user.ID, "account_id": accountCreated.ID, "email": user.Email})
}

func GetUserIDFromJWT(c *gin.Context) *int {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
		return nil
	}

	claims, err := auth.GetClaims(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return nil
	}

	userID := claims.Username
	if userID == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user_id is required"})
		return nil
	}

	userIdInt, err := strconv.Atoi(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user_id is not a number"})
		return nil
	}

	return &userIdInt
}

package controllers

import (
	"digitalmoney/api/model"
	accountRepo "digitalmoney/api/repository/account"
	cardRepo "digitalmoney/api/repository/card"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCardsByAccountID(c *gin.Context) {
	accountID := c.Param("account_id")

	accountIdInt, err := strconv.Atoi(accountID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = accountRepo.GetByID(accountIdInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cardsFounded, err := cardRepo.GetByAccountID(accountIdInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cardsFounded)
}

func CreateCard(c *gin.Context) {
	accountID := c.Param("account_id")

	accountIdInt, err := strconv.Atoi(accountID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "account_id must be a number"})
		return
	}

	var cardBody model.CardRequest
	if err := c.ShouldBindJSON(&cardBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card := model.Card{
		AccountID:      (uint(accountIdInt)),
		NumberID:       cardBody.NumberID,
		ExpirationDate: cardBody.ExpirationDate,
		FirstLastname:  cardBody.FirstLastname,
		Cod:            cardBody.Cod,
	}

	carCreated, err := cardRepo.Create(card)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, carCreated)
}

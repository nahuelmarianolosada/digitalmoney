package controllers

import (
	"digitalmoney/api/model"
	accountRepo "digitalmoney/api/repository/account"
	transferenceRepo "digitalmoney/api/repository/transference"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTransferencesByAccountID(c *gin.Context) {
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

	transferencesFounded, err := transferenceRepo.GetByAccountID(accountIdInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transferencesFounded)
}

func CreateTransference(c *gin.Context) {
	accountID := c.Param("account_id")

	accountIdInt, err := strconv.Atoi(accountID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "account_id must be a number"})
		return
	}

	var transferenceBody model.TransferenceRequest
	if err := c.ShouldBindJSON(&transferenceBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transference := model.Transference {
		AccountID: (uint(accountIdInt)),
		Origin: transferenceBody.Origin,
		Destination: transferenceBody.Destination,
		Amount: transferenceBody.Amount,
		Dated: transferenceBody.Dated,
	}

	transferenceCreated, err := transferenceRepo.Create(transference)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, transferenceCreated)
}



package card

import (
	"digitalmoney/api/model"
	"digitalmoney/api/repository"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func GetByID(cardID int) (*model.Card, error) {
	var card model.Card
	result := repository.DB.First(&card, cardID)

	if result.Error != nil {
		fmt.Printf("ERROR %v", result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("account not found")
		}
		return nil, result.Error
	}
	return &card, nil
}

func GetByAccountID(accountID int) ([]model.Card, error) {
	var cards []model.Card
	result := repository.DB.Where("account_id = ?", accountID).Find(&cards)

	if result.Error != nil {
		fmt.Printf("ERROR %v", result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return cards, nil
		}
		return nil, result.Error
	}
	return cards, nil
}

func Create(card model.Card) (*model.Card, error) {
	tx := repository.DB.Create(&card)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &card, nil
}

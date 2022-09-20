package transference

import (
	"gorm.io/gorm"
	"errors"
	"fmt"
	"digitalmoney/api/repository"
	"digitalmoney/api/model"
)

func GetByID(transferenceID int) (*model.Transference, error) {
	var transference model.Transference
	result := repository.DB.First(&transference, transferenceID)

	if result.Error != nil {
		fmt.Printf("ERROR %v", result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("account not found")
		}
		return nil, result.Error
	}
	return &transference, nil
}

func GetByAccountID(accountID int) ([]model.Transference, error) {
	var transferences []model.Transference
	result := repository.DB.Where("account_id = ?", accountID).Find(&transferences)

	if result.Error != nil {
		fmt.Printf("ERROR %v", result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return transferences, nil
		}
		return nil, result.Error
	}
	return transferences, nil
}

func Create(transference model.Transference) (*model.Transference, error) {
	tx := repository.DB.Create(&transference)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &transference, nil
}


package repository

import (
	"go-pos/model"
	"gorm.io/gorm"
)

type AuthenticationRepo interface {
	Passcode(id int) model.PasscodeData
	CheckUser(id uint, passcode string) bool
}

type authenticationRepo struct {
	db *gorm.DB
}

func (a *authenticationRepo) Passcode(id int) model.PasscodeData {
	result := model.PasscodeData{}
	a.db.Raw("SELECT * FROM cashiers where id = ?", id).Scan(&result)
	return result
}

func (a *authenticationRepo) CheckUser(id uint, passcode string) bool {
	var result model.Cashier
	a.db.Raw("SELECT * FROM cashiers where passcode = ? and id = ? and deleted_at is null", passcode, id).Scan(&result)
	if result.Passcode != passcode {
		return false
	}

	if result.ID != id {
		return false
	}
	return true
}

func NewAuthRepo(db *gorm.DB) AuthenticationRepo {
	return &authenticationRepo{db: db}
}

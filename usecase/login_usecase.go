package usecase

import (
	"go-pos/model"
	"go-pos/repository"
)

type LoginUseCase interface {
	GetPasscode(id int) model.PasscodeData
	LoginUser(id uint, passcode string) bool
}

type loginUseCase struct {
	authRepo repository.AuthenticationRepo
}

func (l *loginUseCase) GetPasscode(id int) model.PasscodeData {
	return l.authRepo.Passcode(id)
}

func (l *loginUseCase) LoginUser(id uint, passcode string) bool {
	return l.authRepo.CheckUser(id, passcode)
}

func NewLoginUseCase(authRepo repository.AuthenticationRepo) LoginUseCase {
	return &loginUseCase{authRepo: authRepo}
}

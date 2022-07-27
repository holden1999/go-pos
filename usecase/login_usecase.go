package usecase

import "go-pos/repository"

type LoginUseCase interface {
	LoginUser(id uint, passcode string) bool
}

type loginUseCase struct {
	authRepo repository.AuthenticationRepo
}

func (l loginUseCase) LoginUser(id uint, passcode string) bool {
	return l.authRepo.CheckUser(id, passcode)
}

func NewLoginUseCase(authRepo repository.AuthenticationRepo) LoginUseCase {
	return &loginUseCase{authRepo: authRepo}
}

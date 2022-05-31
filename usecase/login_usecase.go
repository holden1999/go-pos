package usecase

type LoginUseCase interface {
	LoginUser(email string, password string) bool
}

type loginUseCase struct {
	email    string
	password string
}

func (l loginUseCase) LoginUser(email string, password string) bool {
	return l.email == email && l.password == password
}

func NewLoginUseCase() LoginUseCase {
	return &loginUseCase{
		email:    "gunawanholden15@gmail.com",
		password: "testing",
	}
}

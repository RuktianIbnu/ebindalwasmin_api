package general

import (
	"ebindalwasmin_api/helpers/bcrypt"
	"ebindalwasmin_api/helpers/jwt"
	ur "ebindalwasmin_api/repository/user"
	"errors"
	"fmt"
)

// Usecase ...
type Usecase interface {
	Login(email, password string) (string, error)
	// UpdateOneByID(data *model.User) (rowsAffected int64, err error)
	// GetOneByID(id int64) (result *model.User, err error)
	// GetAll(limit, page int, search string) (result []*model.User, err error)
	// DeleteOneByID(id int64) (rowsAffected int64, err error)
}

type usecase struct {
	userRepo ur.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		ur.NewRepository(),
	}
}

func (m *usecase) Login(email, password string) (string, error) {
	id, pwd, err := m.userRepo.GetUserPasswordByEmail(email)
	if err != nil {
		return "", err
	}

	if !bcrypt.Compare(password, pwd) {
		return "", errors.New("incorrect email or password")
	}

	token, err := jwt.CreateToken(id, email)
	if err != nil {
		return "", fmt.Errorf("failed generating jwt token %s", err.Error())
	}

	return token, nil
}

// func (m *usecase) UpdateOneByID(data *model.User) (rowsAffected int64, err error) {
// 	rowsAffected, err = m.userRepo.UpdateOneByID(data)

// 	if rowsAffected <= 0 {
// 		return rowsAffected, errors.New("no rows affected or data not found")
// 	}

// 	return
// }

// func (m *usecase) GetOneByID(id int64) (result *model.User, err error) {
// 	return m.userRepo.GetOneByID(id)
// }

// func (m *usecase) GetAll(limit, page int, search string) (result []*model.User, err error) {
// 	return m.userRepo.GetAll(limit, (page-1)*limit, search)
// }

// func (m *usecase) DeleteOneByID(id int64) (rowsAffected int64, err error) {
// 	return m.userRepo.DeleteOneByID(id)
// }

package user

import (
	ur "ebindalwasmin_api/repository/user"

	"ebindalwasmin_api/model"
)

// Usecase ...
type Usecase interface {
	// Create(data *model.User) (err error)
	// UpdateOneByID(data *model.User) (rowsAffected int64, err error)
	GetOneByID(id int64) (result *model.User, err error)
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

// func (m *usecase) Create(data *model.User) (err error) {
// 	lastID, err := m.userRepo.Create(data)
// 	if err != nil {
// 		return err
// 	}
// 	data.ID = lastID

// 	return
// }

// func (m *usecase) UpdateOneByID(data *model.User) (rowsAffected int64, err error) {
// 	rowsAffected, err = m.userRepo.UpdateOneByID(data)

// 	if rowsAffected <= 0 {
// 		return rowsAffected, errors.New("no rows affected or data not found")
// 	}

// 	return
// }

func (m *usecase) GetOneByID(id int64) (result *model.User, err error) {
	return m.userRepo.GetOneByID(id)
}

// func (m *usecase) GetAll(limit, page int, search string) (result []*model.User, err error) {
// 	return m.userRepo.GetAll(limit, (page-1)*limit, search)
// }

// func (m *usecase) DeleteOneByID(id int64) (rowsAffected int64, err error) {
// 	return m.userRepo.DeleteOneByID(id)
// }

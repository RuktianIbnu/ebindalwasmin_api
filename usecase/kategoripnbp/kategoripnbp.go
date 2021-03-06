package kategoripnbp

import (
	kr "ebindalwasmin_api/repository/kategoripnbp"

	"ebindalwasmin_api/model"
)

// Usecase ...
type Usecase interface {
	// Create(data *model.User) (err error)
	// UpdateOneByID(data *model.User) (rowsAffected int64, err error)
	// GetOneByID(id int64) (result *model.User, err error)
	GetAllByParent() (result []*model.KategoriParentPNBP, err error)
	// DeleteOneByID(id int64) (rowsAffected int64, err error)
}

type usecase struct {
	kategoriPNBPRepo kr.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		kr.NewRepository(),
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

// func (m *usecase) GetOneByID(id int64) (result *model.User, err error) {
// 	return m.userRepo.GetOneByID(id)
// }

func (m *usecase) GetAllByParent() (result []*model.KategoriParentPNBP, err error) {
	parentList, err := m.kategoriPNBPRepo.GetAllByParent(0)
	if err != nil {
		return nil, nil
	}

	for _, v := range parentList {
		childList, err := m.kategoriPNBPRepo.GetAllByParent(v.ID)
		if err != nil {
			return nil, nil
		}
		v.Child = childList

		for _, j := range childList {
			itemList, err := m.kategoriPNBPRepo.GetAllByParent(j.ID)
			if err != nil {
				return nil, nil
			}
			j.Child = itemList
		}
	}

	return parentList, nil
}

// func (m *usecase) DeleteOneByID(id int64) (rowsAffected int64, err error) {
// 	return m.userRepo.DeleteOneByID(id)
// }

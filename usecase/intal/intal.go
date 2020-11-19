package intal

import (
	ki "ebindalwasmin_api/repository/intal"

	"ebindalwasmin_api/model"
)

// Usecase ...
type Usecase interface {
	// Create(data *model.User) (err error)
	// UpdateOneByID(data *model.User) (rowsAffected int64, err error)
	// GetOneByID(id int64) (result *model.User, err error)
	GetAllByDate(date int64, id_satker int64) (result []*model.Intal, err error)
	GetPivotPerwilayah() (result []*model.PasporPivotPerwilayah, err error)
	GetKelaminPer10hari() (result []*model.IntalPermohonanperKelaminPer10hari, err error)
	// DeleteOneByID(id int64) (rowsAffected int64, err error)
}

type usecase struct {
	intalRepo ki.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		ki.NewRepository(),
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

func (m *usecase) GetAllByDate(date int64, id_satker int64) (result []*model.Intal, err error) {
	return m.intalRepo.GetAllByDate(date, id_satker)
}

func (m *usecase) GetPivotPerwilayah() (result []*model.PasporPivotPerwilayah, err error) {
	return m.intalRepo.GetPivotPerwilayah()
}

func (m *usecase) GetKelaminPer10hari() (result []*model.IntalPermohonanperKelaminPer10hari, err error) {
	return m.intalRepo.GetKelaminPer10hari()
}

// func (m *usecase) DeleteOneByID(id int64) (rowsAffected int64, err error) {
// 	return m.userRepo.DeleteOneByID(id)
// }

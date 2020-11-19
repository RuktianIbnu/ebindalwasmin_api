package pnbp

import (
	pv "ebindalwasmin_api/repository/pnbp"

	"ebindalwasmin_api/model"
)

// Usecase ...
type Usecase interface {
	// Create(data *model.User) (err error)
	// UpdateOneByID(data *model.User) (rowsAffected int64, err error)
	// GetOneByID(id int64) (result *model.User, err error)
	GetAllByDate(date int64, id_satker int64) (result []*model.Pnbp, err error)
	GetAllkategoriPNBPPerbulanTahun() (result []*model.PnbpAllKategoriPerbulanTahun, err error)
	GetTotalPnbp() (result []*model.PnbpGetTotalPnbp, err error)
	GetKelaminPer10hari() (result []*model.PNBPPermohonanperKelaminPer10hari, err error)
	// DeleteOneByID(id int64) (rowsAffected int64, err error)
}

type usecase struct {
	pnbpRepo pv.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		pv.NewRepository(),
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

func (m *usecase) GetAllByDate(date int64, id_satker int64) (result []*model.Pnbp, err error) {
	return m.pnbpRepo.GetAllByDate(date, id_satker)
}

func (m *usecase) GetAllkategoriPNBPPerbulanTahun() (result []*model.PnbpAllKategoriPerbulanTahun, err error) {
	return m.pnbpRepo.GetAllkategoriPNBPPerbulanTahun()
}

func (m *usecase) GetTotalPnbp() (result []*model.PnbpGetTotalPnbp, err error) {
	return m.pnbpRepo.GetTotalPnbp()
}

func (m *usecase) GetKelaminPer10hari() (result []*model.PNBPPermohonanperKelaminPer10hari, err error) {
	return m.pnbpRepo.GetKelaminPer10hari()
}

// func (m *usecase) DeleteOneByID(id int64) (rowsAffected int64, err error) {
// 	return m.userRepo.DeleteOneByID(id)
// }

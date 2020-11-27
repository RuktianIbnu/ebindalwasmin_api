package paspor

import (
	rs "ebindalwasmin_api/repository/master"

	"ebindalwasmin_api/model"
)

// Usecase ...
type Usecase interface {
	// Create(data *model.User) (err error)
	// UpdateOneByID(data *model.User) (rowsAffected int64, err error)
	// GetOneByID(id int64) (result *model.User, err error)
	GetAllSatker() (result []*model.Satker, err error)
	GetReportMonthYear(tm_awal int64, tm_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error)
	// DeleteOneByID(id int64) (rowsAffected int64, err error)
}

type usecase struct {
	satkerRepo rs.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rs.NewRepository(),
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

func (m *usecase) GetAllSatker() (result []*model.Satker, err error) {
	return m.satkerRepo.GetAllSatker()
}

func (m *usecase) GetReportMonthYear(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error) {

	if cekbox == true && id_jenis == 0 && id_satker == 99 {
		return m.satkerRepo.GetReportMonthYear(tgl_awal, tgl_akhir, cekbox, id_jenis, id_satker)
	} else if cekbox == true && id_jenis == 0 && id_satker != 99 {
		return m.satkerRepo.GetReportMonthYearWithIdSatker(tgl_awal, tgl_akhir, cekbox, id_jenis, id_satker)
	} else if cekbox == true && id_jenis != 0 && id_satker == 99 {
		return m.satkerRepo.GetReportMonthYearWithIdJenis(tgl_awal, tgl_akhir, cekbox, id_jenis, id_satker)
	} else if cekbox == true && id_jenis != 0 && id_satker != 99 {
		return m.satkerRepo.GetReportMonthYearWithIdSatkerAndIdJenis(tgl_awal, tgl_akhir, cekbox, id_jenis, id_satker)
	} else if cekbox == false && id_jenis == 0 && id_satker == 99 {
		return m.satkerRepo.GetReportDateAll(tgl_awal, tgl_akhir)
	} else if cekbox == false && id_jenis == 0 && id_satker != 99 {
		return m.satkerRepo.GetReportDateWithIdSatker(tgl_awal, tgl_akhir, cekbox, id_jenis, id_satker)
	} else if cekbox == false && id_jenis != 0 && id_satker == 99 {
		return m.satkerRepo.GetReportDateWithIdJenis(tgl_awal, tgl_akhir, cekbox, id_jenis, id_satker)
	} else if cekbox == false && id_jenis != 0 && id_satker != 99 {
		return m.satkerRepo.GetReportDateWithIdSatkerAndIdJenis(tgl_awal, tgl_akhir, cekbox, id_jenis, id_satker)
	}

	return
}

// func (m *usecase) DeleteOneByID(id int64) (rowsAffected int64, err error) {
// 	return m.userRepo.DeleteOneByID(id)
// }

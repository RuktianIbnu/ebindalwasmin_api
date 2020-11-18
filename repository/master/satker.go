package master

import (
	"database/sql"
	"ebindalwasmin_api/helpers"
	"ebindalwasmin_api/model"
	"log"
)

// Repository ...
type Repository interface {
	// Create(data *model.User) (lastID int64, err error)
	// UpdateOneByID(data *model.User) (rowsAffected int64, err error)
	// GetOneByID(id int64) (*model.User, error)
	// GetUserPasswordByEmail(email string) (int64, string, error)
	GetAllSatker() (result []*model.Satker, err error)
	GetReportMonthYear(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error)
	GetReportMonthYearWithIdSatker(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error)
	GetReportMonthYearWithIdJenis(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error)
	GetReportMonthYearWithIdSatkerAndIdJenis(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error)

	GetReportDateAll(tgl_awal int64, tgl_akhir int64) (result []*model.ReportMonthYear, err error)
	GetReportDateWithIdSatker(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error)
	GetReportDateWithIdJenis(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error)
	GetReportDateWithIdSatkerAndIdJenis(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error)

	// DeleteOneByID(id int64) (rowsAffected int64, err error)
}

type repository struct {
	DB *sql.DB
}

// NewRepository ...
func NewRepository() Repository {
	return &repository{
		DB: helpers.Init().SQL,
	}
}

// func (m *repository) Create(data *model.User) (lastID int64, err error) {
// 	query := `INSERT INTO users(
// 		id_client, id_client_parent, nama_domain, status
// 	) VALUES($1, $2, $3, $4) RETURNING id`

// 	res, err := m.DB.Exec(query,
// 		data.ID,
// 		data.IDClient,
// 		data.IDClientParent,
// 		data.NamaDomain,
// 		data.Status,
// 	)
// 	if err != nil {
// 		return -1, err
// 	}

// 	lastID, _ = res.LastInsertId()

// 	return
// }

// func (m *repository) UpdateOneByID(data *model.User) (rowsAffected int64, err error) {
// 	query := `UPDATE sm_domain_client set
// 	id_client = $2, id_client_parent = $3, nama_domain = $4, status = $5
// 	WHERE id = $1`

// 	res, err := m.DB.Exec(query,
// 		data.ID,
// 		data.IDClient,
// 		data.IDClientParent,
// 		data.NamaDomain,
// 		data.Status,
// 	)
// 	if err != nil {
// 		return -1, err
// 	}

// 	rowsAffected, _ = res.RowsAffected()

// 	return
// }

// func (m *repository) GetOneByID(id int64) (*model.User, error) {
// 	query := `SELECT
// 	coalesce(id, 0),
// 	coalesce(type, ''),
// 	coalesce(name, ''),
// 	coalesce(email, ''),
// 	email_verified_at,
// 	coalesce(password, ''),
// 	password_changed_at,
// 	coalesce(active, 0),
// 	coalesce(timezone, ''),
// 	last_login_at,
// 	coalesce(last_login_ip, ''),
// 	coalesce(to_be_logged_out, 0),
// 	coalesce(provider, ''),
// 	coalesce(provider_id, ''),
// 	coalesce(remember_token, ''),
// 	created_at,
// 	updated_at,
// 	deleted_at,
// 	coalesce(id_kantor, 0)
// 	FROM users
// 	WHERE id = ? AND active = 1`

// 	data := model.User{}

// 	if err := m.DB.QueryRow(query, id).Scan(
// 		&data.ID,
// 		&data.Type,
// 		&data.Name,
// 		&data.Email,
// 		&data.EmailVerifiedAt,
// 		&data.Password,
// 		&data.PasswordChangeAt,
// 		&data.Active,
// 		&data.Timezone,
// 		&data.LastLoginAt,
// 		&data.LastLoginIP,
// 		&data.ToBeLoggedOut,
// 		&data.Provider,
// 		&data.ProviderID,
// 		&data.RememberToken,
// 		&data.CreatedAt,
// 		&data.UpdatedAt,
// 		&data.DeletedAt,
// 		&data.IDKantor,
// 	); err != nil {
// 		return nil, err
// 	}

// 	return &data, nil
// }

// func (m *repository) GetUserPasswordByEmail(email string) (int64, string, error) {
// 	query := `SELECT
// 	id,
// 	coalesce(password, '')
// 	FROM users
// 	WHERE email = ? AND active = 1`

// 	var (
// 		id  int64
// 		pwd string
// 	)

// 	if err := m.DB.QueryRow(query, email).Scan(
// 		&id,
// 		&pwd,
// 	); err != nil {
// 		return -1, "", err
// 	}

// 	return id, pwd, nil
// }

func (m *repository) GetAllSatker() (result []*model.Satker, err error) {
	query := `select 
	coalesce(id_kantor, 0), 
	coalesce(nama_kantor, 0)
	from kantor`

	var (
		list = make([]*model.Satker, 0)
	)

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Satker
		)

		if err := rows.Scan(
			&data.IDKantor,
			&data.Nama_Kantor,
		); err != nil {
			return nil, err
		}

		list = append(list, &data)
	}
	log.Println(list)

	return list, nil
}

func (m *repository) GetReportMonthYear(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error) {
	var query string

	if cekbox == true && id_jenis == 0 && id_satker == 0 {
		query = `SELECT concat(MONTHNAME(tanggal),' ',YEAR(tanggal)) AS periode,
		SUM(total) AS total
		from merge_table_pelayanan GROUP BY periode`
	}

	var (
		list = make([]*model.ReportMonthYear, 0)
	)

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.ReportMonthYear
		)

		if err := rows.Scan(
			&data.Periode,
			&data.Total,
		); err != nil {
			return nil, err
		}
		list = append(list, &data)
	}
	return list, nil
}

func (m *repository) GetReportMonthYearWithIdSatker(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error) {
	var query string

	if cekbox == true && id_jenis == 0 && id_satker != 0 {
		query = `SELECT concat(MONTHNAME(tanggal),' ',YEAR(tanggal)) AS periode, SUM(total) AS total 
		from merge_table_pelayanan where id_satker = ? GROUP BY periode`
	}

	var (
		list = make([]*model.ReportMonthYear, 0)
	)

	rows, err := m.DB.Query(query, id_satker)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.ReportMonthYear
		)

		if err := rows.Scan(
			&data.Periode,
			&data.Total,
		); err != nil {
			return nil, err
		}
		list = append(list, &data)
	}
	return list, nil
}

func (m *repository) GetReportMonthYearWithIdJenis(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error) {
	var query string

	if cekbox == true && id_jenis != 0 && id_satker == 0 {
		switch id_jenis {
		case 1: // paspor ...
			query = `SELECT concat(MONTHNAME(tanggal),' ',YEAR(tanggal)) AS periode, SUM(total) AS total 
				from data_paspor GROUP BY periode`
		case 2: // visa ...
			query = `SELECT concat(MONTHNAME(tanggal),' ',YEAR(tanggal)) AS periode, SUM(total) AS total 
				from data_visa GROUP BY periode`
		case 3: // intal ...
			query = `SELECT concat(MONTHNAME(tanggal),' ',YEAR(tanggal)) AS periode, SUM(total) AS total 
				from data_izinkeimigrasian GROUP BY periode`
		case 4: // pnbp lainnya ...
			query = `SELECT concat(MONTHNAME(tanggal),' ',YEAR(tanggal)) AS periode, SUM(total) AS total 
				from data_pnbplainnya GROUP BY periode`
		}
	}

	var (
		list = make([]*model.ReportMonthYear, 0)
	)

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.ReportMonthYear
		)

		if err := rows.Scan(
			&data.Periode,
			&data.Total,
		); err != nil {
			return nil, err
		}
		list = append(list, &data)
	}
	return list, nil
}

func (m *repository) GetReportMonthYearWithIdSatkerAndIdJenis(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error) {
	var query string

	if cekbox == true && id_jenis != 0 && id_satker != 0 {
		switch id_jenis {
		case 1: // paspor ...
			query = `SELECT concat(MONTHNAME(tanggal),' ',YEAR(tanggal)) AS periode, SUM(total) AS total 
				from data_paspor where id_kantor = ? GROUP BY periode`
		case 2: // visa ...
			query = `SELECT concat(MONTHNAME(tanggal),' ',YEAR(tanggal)) AS periode, SUM(total) AS total 
				from data_visa where id_kantor = ? GROUP BY periode`
		case 3: // intal ...
			query = `SELECT concat(MONTHNAME(tanggal),' ',YEAR(tanggal)) AS periode, SUM(total) AS total 
				from data_izinkeimigrasian where id_kantor = ? GROUP BY periode`
		case 4: // pnbp lainnya ...
			query = `SELECT concat(MONTHNAME(tanggal),' ',YEAR(tanggal)) AS periode, SUM(total) AS total 
				from data_pnbplainnya where id_kantor = ? GROUP BY periode`
		}
	}

	var (
		list = make([]*model.ReportMonthYear, 0)
	)

	rows, err := m.DB.Query(query, id_satker)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.ReportMonthYear
		)

		if err := rows.Scan(
			&data.Periode,
			&data.Total,
		); err != nil {
			return nil, err
		}
		list = append(list, &data)
	}
	return list, nil
}

func (m *repository) GetReportDateAll(tgl_awal int64, tgl_akhir int64) (result []*model.ReportMonthYear, err error) {
	var query string

	if tgl_awal != 0 && tgl_akhir != 0 {
		query = `SELECT jenispnbp, tanggal, total 
		from merge_table_pelayanan where tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d')`
	}

	var (
		list = make([]*model.ReportMonthYear, 0)
	)

	rows, err := m.DB.Query(query, tgl_awal, tgl_akhir)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.ReportMonthYear
		)

		if err := rows.Scan(
			&data.JenisPnbp,
			&data.Tanggal,
			&data.Total,
		); err != nil {
			return nil, err
		}
		list = append(list, &data)
	}
	return list, nil
}

func (m *repository) GetReportDateWithIdSatker(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error) {
	var query string

	if cekbox == true && id_jenis == 0 && id_satker != 0 {
		query = `SELECT k.nama_layanan AS jenispnbp, p.tanggal, p.total 
		from data_izinkeimigrasian AS p 
		INNER JOIN kategoripnbps AS k ON k.id = p.id_jenis 
		where tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d') AND id_satker = ?`
	}
	var (
		list = make([]*model.ReportMonthYear, 0)
	)

	rows, err := m.DB.Query(query, tgl_awal, tgl_akhir, id_satker)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.ReportMonthYear
		)

		if err := rows.Scan(
			&data.JenisPnbp,
			&data.Tanggal,
			&data.Total,
		); err != nil {
			return nil, err
		}
		list = append(list, &data)
	}
	return list, nil
}

func (m *repository) GetReportDateWithIdJenis(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error) {
	var query string

	log.Println(cekbox, id_jenis, id_satker, tgl_awal, tgl_akhir)
	if cekbox == false && id_jenis != 0 && id_satker == 0 {
		switch id_jenis {
		case 1: // paspor ...
			query = `SELECT k.nama_layanan AS jenispnbp, p.tanggal, p.total 
			from data_paspor AS p 
			INNER JOIN kategoripnbps AS k ON k.id = p.id_jenis 
			where tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d')
			`
		case 2: // visa ...
			query = `SELECT k.nama_layanan AS jenispnbp, p.tanggal, p.total 
			from data_visa AS p 
			INNER JOIN kategoripnbps AS k ON k.id = p.id_jenis 
			where tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d')`
		case 3: // intal ...
			query = `SELECT k.nama_layanan AS jenispnbp, p.tanggal, p.total 
			from data_izinkeimigrasian AS p 
			INNER JOIN kategoripnbps AS k ON k.id = p.id_jenis
			where tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d')`
		case 4: // pnbp lainnya ...
			query = `SELECT k.nama_layanan AS jenispnbp, p.tanggal, p.total 
			from data_pnbplainnya AS p 
			INNER JOIN kategoripnbps AS k ON k.id = p.id_jenis
			where tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d')`
		}
	}

	var (
		list = make([]*model.ReportMonthYear, 0)
	)

	rows, err := m.DB.Query(query, tgl_awal, tgl_akhir)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.ReportMonthYear
		)

		if err := rows.Scan(
			&data.JenisPnbp,
			&data.Tanggal,
			&data.Total,
		); err != nil {
			return nil, err
		}
		list = append(list, &data)
	}
	return list, nil
}

func (m *repository) GetReportDateWithIdSatkerAndIdJenis(tgl_awal int64, tgl_akhir int64, cekbox bool, id_jenis int64, id_satker int64) (result []*model.ReportMonthYear, err error) {
	var query string

	if cekbox == false && id_jenis != 0 && id_satker != 0 {
		switch id_jenis {
		case 1: // paspor ...
			query = `SELECT k.nama_layanan AS jenispnbp, p.tanggal, p.total 
			from data_paspor AS p 
			INNER JOIN kategoripnbps AS k ON k.id = p.id_jenis
			where tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d') AND id_kantor = ?`
		case 2: // visa ...
			query = `SELECT k.nama_layanan AS jenispnbp, p.tanggal, p.total 
			from data_visa AS p 
			INNER JOIN kategoripnbps AS k ON k.id = p.id_jenis
			where tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d') AND id_kantor = ?`
		case 3: // intal ...
			query = `SELECT k.nama_layanan AS jenispnbp, p.tanggal, p.total 
			from data_izinkeimigrasian AS p 
			INNER JOIN kategoripnbps AS k ON k.id = p.id_jenis 
			where tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d') AND id_kantor = ?`
		case 4: // pnbp lainnya ...
			query = `SELECT k.nama_layanan AS jenispnbp, p.tanggal, p.total 
			from data_pnbplainnnya AS p 
			INNER JOIN kategoripnbps AS k ON k.id = p.id_jenis
			where tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d') AND id_kantor = ?`
		}
	}

	var (
		list = make([]*model.ReportMonthYear, 0)
	)

	rows, err := m.DB.Query(query, tgl_awal, tgl_akhir, id_satker)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.ReportMonthYear
		)

		if err := rows.Scan(
			&data.JenisPnbp,
			&data.Tanggal,
			&data.Total,
		); err != nil {
			return nil, err
		}
		list = append(list, &data)
	}
	return list, nil
}

// func (m *repository) DeleteOneByID(id int64) (rowsAffected int64, err error) {
// 	query := `DELETE FROM sm_domain_client WHERE id = $1`

// 	res, err := m.DB.Exec(query, id)
// 	if err != nil {
// 		return -1, err
// 	}

// 	rowsAffected, _ = res.RowsAffected()

// 	return
// }

package paspor

import (
	"database/sql"
	"ebindalwasmin_api/helpers"
	"ebindalwasmin_api/model"
	"log"
	"time"
)

// Repository ...
type Repository interface {
	// Create(data *model.User) (lastID int64, err error)
	// UpdateOneByID(data *model.User) (rowsAffected int64, err error)
	// GetOneByID(id int64) (*model.User, error)
	// GetUserPasswordByEmail(email string) (int64, string, error)
	GetAllByDate(date int64, id_satker int64) (result []*model.Paspor, err error)
	GetPivotPerwilayah() (result []*model.PasporPivotPerwilayah, err error)
	GetKelaminPer10hari(id_kantor int64) (result []*model.PasporPermohonanperKelaminPer10hari, err error)
	GetPnbpPaspor(id_layanan int64, id_kantor int64) (result []*model.GeneralPnbp, err error)
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

func (m *repository) GetPivotPerwilayah() (result []*model.PasporPivotPerwilayah, err error) {
	query := `select 
	coalesce(periode, ''), 
	coalesce(kab_bangka, 0), 
	coalesce(kab_bangka_barat, 0), 
	coalesce(kab_bangka_selatan, 0), 
	coalesce(kab_bangka_tengah, 0), 
	coalesce(pangkal_pinang, 0),
	coalesce(kab_belitung_timur, 0), 
	coalesce(kab_belitung, 0), 
	coalesce(kab_lainnya_1, 0), 
	coalesce(kab_lainnya_2, 0)
	from paspor_pivot_wilayah`

	var (
		list = make([]*model.PasporPivotPerwilayah, 0)
	)

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.PasporPivotPerwilayah
		)

		if err := rows.Scan(
			&data.Periode,
			&data.KabBangka,
			&data.KabBangkaBarat,
			&data.KabBangkaSelatan,
			&data.KabBangkaTengah,
			&data.PangkalPinang,
			&data.KabBelitungTimur,
			&data.KabBelitung,
			&data.KabLainnya1,
			&data.KabLainnya2,
		); err != nil {
			return nil, err
		}

		list = append(list, &data)
	}
	log.Println(list)

	return list, nil
}

func (m *repository) GetAllByDate(date int64, id_satker int64) (result []*model.Paspor, err error) {
	query := `select 
	coalesce(id, 0), 
	coalesce(id_jenis, 0), 
	coalesce(id_user, 0), 
	coalesce(id_kantor, 0), 
	tanggal, 
	coalesce(laki, 0), 
	coalesce(perempuan, 0), 
	coalesce(total, 0), 
	coalesce(id_wilayah_kerja, 0)
	from data_paspor where tanggal = FROM_UNIXTIME(?, '%Y-%m-%d') AND id_kantor = ?`

	var (
		list = make([]*model.Paspor, 0)
	)

	rows, err := m.DB.Query(query, date, id_satker)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Paspor
		)

		if err := rows.Scan(
			&data.ID,
			&data.IDJenis,
			&data.IDUser,
			&data.IDKantor,
			&data.Tanggal,
			&data.Laki,
			&data.Perempuan,
			&data.Total,
			&data.IDWilayahKerja,
		); err != nil {
			return nil, err
		}

		list = append(list, &data)
	}
	log.Println(list, date)

	return list, nil
}

func (m *repository) GetKelaminPer10hari(id_kantor int64) (result []*model.PasporPermohonanperKelaminPer10hari, err error) {
	var (
		query string
	)
	if id_kantor == 99 {
		query = `select 
		COALESCE(sum(laki), 0) AS laki,
		COALESCE(sum(perempuan), 0) AS wanita,
		id_kantor
		from paspor WHERE tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d')`
		var (
			list = make([]*model.PasporPermohonanperKelaminPer10hari, 0)
		)

		now := time.Now()

		then := now.AddDate(0, 0, -150)

		log.Println(then.Unix(), now.Unix())
		rows, err := m.DB.Query(query, then.Unix(), now.Unix())
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				data model.PasporPermohonanperKelaminPer10hari
			)

			if err := rows.Scan(
				&data.Laki,
				&data.Perempuan,
				&data.IDKantor,
			); err != nil {
				return nil, err
			}

			list = append(list, &data)
		}

		return list, nil
	} else if id_kantor != 99 {
		query = `select 
		COALESCE(sum(laki), 0) AS laki,
		COALESCE(sum(perempuan), 0) AS wanita,
		id_kantor
		from paspor WHERE id_kantor = ? and tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d')
		GROUP BY id_kantor;`

		var (
			list = make([]*model.PasporPermohonanperKelaminPer10hari, 0)
		)

		now := time.Now()

		then := now.AddDate(0, 0, -150)

		log.Println(then.Unix(), now.Unix())
		rows, err := m.DB.Query(query, id_kantor, then.Unix(), now.Unix())
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				data model.PasporPermohonanperKelaminPer10hari
			)

			if err := rows.Scan(
				&data.Laki,
				&data.Perempuan,
				&data.IDKantor,
			); err != nil {
				return nil, err
			}

			list = append(list, &data)
		}

		return list, nil
	}

	return
}

func (m *repository) GetPnbpPaspor(id_layanan int64, id_kantor int64) (result []*model.GeneralPnbp, err error) {
	var (
		query string
		rows  *sql.Rows
	)

	if id_layanan == 1 {
		if id_kantor == 99 {
			// PASPOR
			query = `SELECT k.nama_kantor,CONCAT(CAST(MONTHNAME(p.tanggal) AS CHAR(3)),' ',YEAR(p.tanggal)) AS periode,
			SUM(p.total) AS total
			from data_paspor AS p 
			inner join kantor AS k ON k.id_kantor = p.id_kantor 
			GROUP BY year(p.tanggal),month(p.tanggal) order BY p.tanggal`

			rows, err = m.DB.Query(query)
			if err != nil {
				return nil, err
			}
			defer rows.Close()

			var (
				list = make([]*model.GeneralPnbp, 0)
			)

			for rows.Next() {
				var (
					data model.GeneralPnbp
				)

				if err := rows.Scan(
					&data.NamaKantor,
					&data.Periode,
					&data.Total,
				); err != nil {
					return nil, err
				}

				list = append(list, &data)
			}

			return list, nil
		} else {
			query = `SELECT k.nama_kantor,CONCAT(CAST(MONTHNAME(p.tanggal) AS CHAR(3)),' ',YEAR(p.tanggal)) AS periode,
			SUM(p.total) AS total
			from data_paspor AS p 
			inner join kantor AS k ON k.id_kantor = p.id_kantor 
			WHERE p.id_kantor = ? GROUP BY year(p.tanggal),month(p.tanggal) order BY p.tanggal`

			rows, err = m.DB.Query(query, id_kantor)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			var (
				list = make([]*model.GeneralPnbp, 0)
			)

			for rows.Next() {
				var (
					data model.GeneralPnbp
				)

				if err := rows.Scan(
					&data.NamaKantor,
					&data.Periode,
					&data.Total,
				); err != nil {
					return nil, err
				}

				list = append(list, &data)
			}

			return list, nil
		}
	} else if id_layanan == 2 {
		if id_kantor == 99 {
			// VISA
			query = `SELECT k.nama_kantor,CONCAT(CAST(MONTHNAME(p.tanggal) AS CHAR(3)),' ',YEAR(p.tanggal)) AS periode,
			SUM(p.total) AS total
			from data_visa AS p 
			inner join kantor AS k ON k.id_kantor = p.id_kantor 
			GROUP BY year(p.tanggal),month(p.tanggal) order BY p.tanggal`

			rows, err = m.DB.Query(query)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			var (
				list = make([]*model.GeneralPnbp, 0)
			)

			for rows.Next() {
				var (
					data model.GeneralPnbp
				)

				if err := rows.Scan(
					&data.NamaKantor,
					&data.Periode,
					&data.Total,
				); err != nil {
					return nil, err
				}

				list = append(list, &data)
			}

			return list, nil
		} else {
			query = `SELECT k.nama_kantor,CONCAT(CAST(MONTHNAME(p.tanggal) AS CHAR(3)),' ',YEAR(p.tanggal)) AS periode,
			SUM(p.total) AS total
			from data_visa AS p 
			inner join kantor AS k ON k.id_kantor = p.id_kantor 
			WHERE p.id_kantor = ? GROUP BY year(p.tanggal),month(p.tanggal) order BY p.tanggal`

			rows, err = m.DB.Query(query, id_kantor)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			var (
				list = make([]*model.GeneralPnbp, 0)
			)

			for rows.Next() {
				var (
					data model.GeneralPnbp
				)

				if err := rows.Scan(
					&data.NamaKantor,
					&data.Periode,
					&data.Total,
				); err != nil {
					return nil, err
				}

				list = append(list, &data)
			}

			return list, nil
		}
	} else if id_layanan == 3 {
		if id_kantor == 99 {
			// INTAL
			query = `SELECT k.nama_kantor,CONCAT(CAST(MONTHNAME(p.tanggal) AS CHAR(3)),' ',YEAR(p.tanggal)) AS periode,
			SUM(p.total) AS total
			from data_izinkeimigrasian AS p 
			inner join kantor AS k ON k.id_kantor = p.id_kantor 
			GROUP BY year(p.tanggal),month(p.tanggal) order BY p.tanggal`

			rows, err = m.DB.Query(query)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			var (
				list = make([]*model.GeneralPnbp, 0)
			)

			for rows.Next() {
				var (
					data model.GeneralPnbp
				)

				if err := rows.Scan(
					&data.NamaKantor,
					&data.Periode,
					&data.Total,
				); err != nil {
					return nil, err
				}

				list = append(list, &data)
			}

			return list, nil
		} else {
			query = `SELECT k.nama_kantor,CONCAT(CAST(MONTHNAME(p.tanggal) AS CHAR(3)),' ',YEAR(p.tanggal)) AS periode,
			SUM(p.total) AS total
			from data_izinkeimigrasian AS p 
			inner join kantor AS k ON k.id_kantor = p.id_kantor 
			WHERE p.id_kantor = ? GROUP BY year(p.tanggal),month(p.tanggal) order BY p.tanggal`

			rows, err = m.DB.Query(query, id_kantor)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			var (
				list = make([]*model.GeneralPnbp, 0)
			)

			for rows.Next() {
				var (
					data model.GeneralPnbp
				)

				if err := rows.Scan(
					&data.NamaKantor,
					&data.Periode,
					&data.Total,
				); err != nil {
					return nil, err
				}

				list = append(list, &data)
			}

			return list, nil
		}
	} else if id_layanan == 4 {
		if id_kantor == 99 {
			// PNBP LAINNYA
			query = `SELECT k.nama_kantor,CONCAT(CAST(MONTHNAME(p.tanggal) AS CHAR(3)),' ',YEAR(p.tanggal)) AS periode,
			SUM(p.total) AS total
			from data_pnbplainnya AS p 
			inner join kantor AS k ON k.id_kantor = p.id_kantor 
			GROUP BY year(p.tanggal),month(p.tanggal) order BY p.tanggal`

			rows, err = m.DB.Query(query)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			var (
				list = make([]*model.GeneralPnbp, 0)
			)

			for rows.Next() {
				var (
					data model.GeneralPnbp
				)

				if err := rows.Scan(
					&data.NamaKantor,
					&data.Periode,
					&data.Total,
				); err != nil {
					return nil, err
				}

				list = append(list, &data)
			}

			return list, nil
		} else {
			query = `SELECT k.nama_kantor,CONCAT(CAST(MONTHNAME(p.tanggal) AS CHAR(3)),' ',YEAR(p.tanggal)) AS periode,
			SUM(p.total) AS total
			from data_pnbplainnya AS p 
			inner join kantor AS k ON k.id_kantor = p.id_kantor 
			WHERE p.id_kantor = ? GROUP BY year(p.tanggal),month(p.tanggal) order BY p.tanggal`

			rows, err = m.DB.Query(query, id_kantor)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			var (
				list = make([]*model.GeneralPnbp, 0)
			)

			for rows.Next() {
				var (
					data model.GeneralPnbp
				)

				if err := rows.Scan(
					&data.NamaKantor,
					&data.Periode,
					&data.Total,
				); err != nil {
					return nil, err
				}

				list = append(list, &data)
			}

			return list, nil
		}
	}

	return
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

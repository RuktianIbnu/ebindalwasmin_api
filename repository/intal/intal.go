package intal

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
	GetAllByDate(date int64, id_satker int64) (result []*model.Intal, err error)
	GetPivotPerwilayah() (result []*model.PasporPivotPerwilayah, err error)
	GetKelaminPer10hari(id_kantor int64) (result []*model.IntalPermohonanperKelaminPer10hari, err error)
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

func (m *repository) GetAllByDate(date int64, id_satker int64) (result []*model.Intal, err error) {
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
	from data_izinkeimigrasian where tanggal = FROM_UNIXTIME(?, '%Y-%m-%d') AND id_kantor = ?`

	var (
		list = make([]*model.Intal, 0)
	)

	rows, err := m.DB.Query(query, date, id_satker)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Intal
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
	from intal_pivot_wilayah`

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

func (m *repository) GetKelaminPer10hari(id_kantor int64) (result []*model.IntalPermohonanperKelaminPer10hari, err error) {
	var (
		query string
	)
	if id_kantor == 0 {
		query = `select 
		coalesce(sum(laki) AS laki, 0), 
		coalesce(sum(perempuan) AS wanita, 0),
		id_kantor
		from izintinggal WHERE tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d')`
		var (
			list = make([]*model.IntalPermohonanperKelaminPer10hari, 0)
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
				data model.IntalPermohonanperKelaminPer10hari
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
	} else if id_kantor > 0 {
		query = `select 
		coalesce(sum(laki) AS laki, 0), 
		coalesce(sum(perempuan) AS wanita, 0),
		id_kantor
		from izintinggal WHERE id_kantor = ? and tanggal BETWEEN FROM_UNIXTIME(?, '%Y-%m-%d') AND FROM_UNIXTIME(?, '%Y-%m-%d')
		GROUP BY id_kantor;`

		var (
			list = make([]*model.IntalPermohonanperKelaminPer10hari, 0)
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
				data model.IntalPermohonanperKelaminPer10hari
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

// func (m *repository) DeleteOneByID(id int64) (rowsAffected int64, err error) {
// 	query := `DELETE FROM sm_domain_client WHERE id = $1`

// 	res, err := m.DB.Exec(query, id)
// 	if err != nil {
// 		return -1, err
// 	}

// 	rowsAffected, _ = res.RowsAffected()

// 	return
// }

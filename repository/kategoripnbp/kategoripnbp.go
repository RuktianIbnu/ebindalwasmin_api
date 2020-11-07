package kategoripnbp

import (
	"database/sql"
	"ebindalwasmin_api/helpers"
	"ebindalwasmin_api/model"
)

// Repository ...
type Repository interface {
	// Create(data *model.User) (lastID int64, err error)
	// UpdateOneByID(data *model.User) (rowsAffected int64, err error)
	// GetOneByID(id int64) (*model.User, error)
	// GetUserPasswordByEmail(email string) (int64, string, error)
	GetAllByParent(parent int64) (result []*model.KategoriParent1PNBP, err error)
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

func (m *repository) GetAllByParent(parent int64) (result []*model.KategoriParent1PNBP, err error) {
	query := `select 
	coalesce(id, 0), 
	coalesce(nama_layanan, ''), 
	coalesce(satuan, ''), 
	coalesce(parent, 0), 
	coalesce(tarifpnbp, 0), 
	created_at, 
	updated_at, 
	coalesce(status, 0) 
	from kategoripnbps where parent = ?`

	var (
		list1 = make([]*model.KategoriParent1PNBP, 0)
	)

	rows1, err := m.DB.Query(query, parent)
	if err != nil {
		return nil, err
	}
	defer rows1.Close()

	for rows1.Next() {
		var (
			data1 model.KategoriParent1PNBP
		)

		if err := rows1.Scan(
			&data1.ID,
			&data1.NamaLayanan,
			&data1.Satuan,
			&data1.Parent,
			&data1.TarifPNBP,
			&data1.CreatedAt,
			&data1.UpdatedAt,
			&data1.Status,
		); err != nil {
			return nil, err
		}
		// Batas ...
		// var (
		// 	list2 = make([]*model.KategoriParent2PNBP, 0)
		// )

		// rows2, err := m.DB.Query(query, data1.ID)
		// if err != nil {
		// 	return nil, err
		// }
		// defer rows2.Close()

		// for rows2.Next() {
		// 	var (
		// 		data2 model.KategoriParent2PNBP
		// 	)

		// 	if err := rows2.Scan(
		// 		&data2.ID,
		// 		&data2.NamaLayanan,
		// 		&data2.Satuan,
		// 		&data2.Parent,
		// 		&data2.TarifPNBP,
		// 		&data2.CreatedAt,
		// 		&data2.UpdatedAt,
		// 		&data2.Status,
		// 	); err != nil {
		// 		return nil, err
		// 	}
		// 	list2 = append(list2, &data2)
		// 	log.Println("parent 2 : ", data2)
		// }
		// Batas ...
		list1 = append(list1, &data1)
		//log.Println("parent 1 : ", data1.ID)
	}
	return list1, nil
}

// func (m *repository) getStructParent2(id int64) (result []*model.KategoriParent2PNBP, err error) {
// 	query := `select
// 	coalesce(id, 0),
// 	coalesce(nama_layanan, ''),
// 	coalesce(satuan, ''),
// 	coalesce(parent, 0),
// 	coalesce(tarifpnbp, 0),
// 	created_at,
// 	updated_at,
// 	coalesce(status, 0)
// 	from kategoripnbps where parent = ?`

// 	var (
// 		list2 = make([]*model.KategoriParent2PNBP, 0)
// 	)

// 	rows2, err := m.DB.Query(query, id)
// 	if err != nil {

// 	}
// 	defer rows2.Close()

// 	for rows2.Next() {
// 		var (
// 			data2 model.KategoriParent2PNBP
// 		)

// 		if err := rows2.Scan(
// 			&data2.ID,
// 			&data2.NamaLayanan,
// 			&data2.Satuan,
// 			&data2.Parent,
// 			&data2.TarifPNBP,
// 			&data2.CreatedAt,
// 			&data2.UpdatedAt,
// 			&data2.Status,
// 		); err != nil {
// 		}
// 		m.getStructParent2(data2.ID)
// 		list2 = append(list2, &data2)
// 		log.Println("parent 2 : ", data2.ID)
// 	}
// 	return list2, nil
// }

// func (m *repository) getStructParent3(id int64) (result []*model.KategoriParent3PNBP, err error) {
// 	query := `select
// 	coalesce(id, 0),
// 	coalesce(nama_layanan, ''),
// 	coalesce(satuan, ''),
// 	coalesce(parent, 0),
// 	coalesce(tarifpnbp, 0),
// 	created_at,
// 	updated_at,
// 	coalesce(status, 0)
// 	from kategoripnbps where parent = ?`

// 	var (
// 		list3 = make([]*model.KategoriParent3PNBP, 0)
// 	)

// 	rows3, err := m.DB.Query(query, id)
// 	if err != nil {

// 	}
// 	defer rows3.Close()

// 	for rows3.Next() {
// 		var (
// 			data3 model.KategoriParent3PNBP
// 		)

// 		if err := rows3.Scan(
// 			&data3.ID,
// 			&data3.NamaLayanan,
// 			&data3.Satuan,
// 			&data3.Parent,
// 			&data3.TarifPNBP,
// 			&data3.CreatedAt,
// 			&data3.UpdatedAt,
// 			&data3.Status,
// 		); err != nil {
// 		}
// 		list3 = append(list3, &data3)
// 		//log.Println("parent 3 : ", data3)
// 	}
// 	return list3, nil
// }

// func (m *repository) DeleteOneByID(id int64) (rowsAffected int64, err error) {
// 	query := `DELETE FROM sm_domain_client WHERE id = $1`

// 	res, err := m.DB.Exec(query, id)
// 	if err != nil {
// 		return -1, err
// 	}

// 	rowsAffected, _ = res.RowsAffected()

// 	return
// }

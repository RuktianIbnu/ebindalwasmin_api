package user

import (
	"database/sql"
	"ebindalwasmin_api/helpers"
	"ebindalwasmin_api/model"
)

// Repository ...
type Repository interface {
	Create(data *model.User) (lastID int64, err error)
	UpdateOneByID(data *model.User) (rowsAffected int64, err error)
	GetOneByID(id int64) (result *model.User, err error)
	GetAll(limit, page int, search string) (result []*model.User, err error)
	DeleteOneByID(id int64) (rowsAffected int64, err error)
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

func (m *repository) Create(data *model.User) (lastID int64, err error) {
	query := `INSERT INTO users(
		id_client, id_client_parent, nama_domain, status
	) VALUES($1, $2, $3, $4) RETURNING id`

	res, err := m.DB.Exec(query,
		data.ID,
		data.IDClient,
		data.IDClientParent,
		data.NamaDomain,
		data.Status,
	)
	if err != nil {
		return -1, err
	}

	lastID, _ = res.LastInsertId()

	return
}

func (m *repository) UpdateOneByID(data *model.User) (rowsAffected int64, err error) {
	query := `UPDATE sm_domain_client set 
	id_client = $2, id_client_parent = $3, nama_domain = $4, status = $5 
	WHERE id = $1`

	res, err := m.DB.Exec(query,
		data.ID,
		data.IDClient,
		data.IDClientParent,
		data.NamaDomain,
		data.Status,
	)
	if err != nil {
		return -1, err
	}

	rowsAffected, _ = res.RowsAffected()

	return
}

func (m *repository) GetOneByID(id int64) (result *model.User, err error) {
	query := `SELECT id, id_client, id_client_parent, nama_domain, status 
	FROM sm_domain_client 
	WHERE id = $1`

	var (
		data model.User

		idS            sql.NullInt64
		idClient       sql.NullInt64
		idClientParent sql.NullInt64
		namaDomain     sql.NullString
		status         sql.NullBool
	)

	if err := m.DB.QueryRow(query, id).Scan(
		&idS,
		&idClient,
		&idClientParent,
		&namaDomain,
		&status,
	); err != nil {
		return nil, err
	}

	data.ID = idS.Int64
	data.IDClient = idClient.Int64
	data.IDClientParent = idClientParent.Int64
	data.NamaDomain = namaDomain.String
	data.Status = status.Bool

	return &data, nil
}

func (m *repository) GetAll(limit, page int, search string) (result []*model.User, err error) {
	query := `SELECT id, id_client, id_client_parent, nama_domain, status 
	FROM sm_domain_client
	WHERE nama_domain LIKE '%' || $3 || '%'
	LIMIT $1 OFFSET $2`

	var (
		list = make([]*model.User, 0)
	)

	rows, err := m.DB.Query(query, limit, page, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.User

			idS            sql.NullInt64
			idClient       sql.NullInt64
			idClientParent sql.NullInt64
			namaDomain     sql.NullString
			status         sql.NullBool
		)

		if err := rows.Scan(
			&idS,
			&idClient,
			&idClientParent,
			&namaDomain,
			&status,
		); err != nil {
			return nil, err
		}

		data.ID = idS.Int64
		data.IDClient = idClient.Int64
		data.IDClientParent = idClientParent.Int64
		data.NamaDomain = namaDomain.String
		data.Status = status.Bool

		list = append(list, &data)
	}

	return list, nil
}

func (m *repository) DeleteOneByID(id int64) (rowsAffected int64, err error) {
	query := `DELETE FROM sm_domain_client WHERE id = $1`

	res, err := m.DB.Exec(query, id)
	if err != nil {
		return -1, err
	}

	rowsAffected, _ = res.RowsAffected()

	return
}
